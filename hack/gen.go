package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var (
	rootDir         string
	globalJSExports = []string{}
)

func main() {
	var err error
	rootDir, err = os.Getwd()
	if err != nil {
		log.Fatalln("os.Getwd() error", err)
	}

	os.Chdir(rootDir) // sometimes go forgets where it is

	dirs, err := getRelevantDirs()
	if err != nil {
		log.Fatalln("getRelevantDirs() error", err)
	}

	for _, dir := range dirs {
		cleanDir(dir)
		err := generate(dir)
		if err != nil {
			log.Fatalln(err)
		}
	}

	writeGlobalJS()
	runTsc()
}

type langconfig struct {
	Go        bool
	JS        bool
	TS        bool
	noService bool
}

func configFromGenconfig(genconfig []byte) langconfig {
	lc := langconfig{}

	fields := strings.Fields(string(genconfig))
	for _, field := range fields {
		switch field {
		case "js":
			lc.JS = true
		case "ts":
			lc.TS = true
		case "go":
			lc.Go = true
		case "--noService":
			lc.noService = true
		}
	}

	return lc
}

func generate(dir string) error {
	log.Printf("[%s] generating...", dir)

	genconfig, err := ioutil.ReadFile(filepath.Join(dir, ".genconfig"))
	if err != nil {
		return err
	}

	config := configFromGenconfig(genconfig)

	flags := getProtocFlags(config, dir)
	if len(flags) == 0 {
		return errors.New("no valid configuration was present")
	}

	err = runProtoc(dir, flags...)
	if err != nil {
		return err
	}

	if config.JS {
		err = generateJs(dir, config)
	}

	return err
}

func generateJs(dir string, config langconfig) error {
	globalJSExports = append(globalJSExports, dir)
	protos := getLocalProtoNames(dir)

	exports := []string{}
	for _, protoName := range protos {
		exports = append(
			exports,
			fmt.Sprintf("./%s_pb", protoName),
		)

		if !config.noService {
			exports = append(
				exports,
				fmt.Sprintf("./%s_pb_service", protoName),
			)
		}
	}

	tsContents := strings.Builder{}

	for _, export := range exports {
		safeName := strings.Replace(filepath.Base(export), ".proto", "", -1)
		tsContents.WriteString(fmt.Sprintf("export * from './%s';\n", safeName))
	}

	tsContents.WriteString("\n")

	err := ioutil.WriteFile(filepath.Join(dir, "index.ts"), []byte(tsContents.String()), 0666)
	if err != nil {
		return err
	}
	return nil
}

func writeGlobalJS() error {
	tsContents := strings.Builder{}

	tsContents.WriteString("export { Empty } from 'google-protobuf/google/protobuf/empty_pb';\n")

	for _, export := range globalJSExports {
		jsName := toCamel(strings.Replace(export, "/", "_", -1))
		tsContents.WriteString(fmt.Sprintf("import * as %s from './%s'; export { %s };\n", jsName, export, jsName))
	}

	tsContents.WriteString("\n")

	err := ioutil.WriteFile(filepath.Join(rootDir, "index.ts"), []byte(tsContents.String()), 0666)
	if err != nil {
		return err
	}

	return nil
}

func getRelevantDirs() ([]string, error) {
	dirs := []string{}

	matches, err := glob(".", ".genconfig", -1)
	if err != nil {
		return dirs, err
	}

	for _, match := range matches {
		dirs = append(
			dirs,
			filepath.Clean("./"+
				strings.Replace(match, ".genconfig", "", -1),
			),
		)
	}

	return dirs, nil
}

func getProtocFlags(lc langconfig, outputDir string) []string {
	flags := []string{"-I" + rootDir}

	if lc.JS {
		flags = append(flags, "--js_out=import_style=commonjs,binary:.")
	}

	if lc.TS {
		if lc.noService {
			flags = append(flags, "--ts_out=service=false:.")
		} else {
			flags = append(flags, "--ts_out=service=true:.")
		}
	}

	if lc.Go {
		if lc.noService {
			flags = append(flags, "--go_out=paths=source_relative:.")
		} else {
			flags = append(flags, "--go_out=paths=source_relative,plugins=grpc:.")
		}
	}

	return flags
}

func assertGlob(i []string, err error) []string {
	if err != nil {
		log.Fatalln(err)
	}

	return i
}

func cleanDir(dir string) {
	log.Printf(`[%s] cleaning...`, dir)

	jsFiles := assertGlob(glob(dir, ".js", -1))
	goFiles := assertGlob(glob(dir, ".go", -1))
	tsFiles := assertGlob(glob(dir, ".ts", -1))
	mapFiles := assertGlob(glob(dir, ".map", -1))

	files := append(jsFiles, goFiles...)
	files = append(files, tsFiles...)
	files = append(files, mapFiles...)

	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func getLocalProtoNames(dir string) []string {
	return assertGlob(glob(dir, ".proto", 1))
}

func runProtoc(dir string, flags ...string) error {
	log.Printf("[%s] ... protoc\n", dir)

	args := []string{
		fmt.Sprintf("--plugin=protoc-gen-ts=%s/node_modules/.bin/protoc-gen-ts", rootDir),
	}

	files, err := glob(dir, ".proto", 1)
	if err != nil {
		return err
	}

	args = append(args, flags...)
	args = append(args, files...)

	cmd := exec.Command("protoc", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func runTsc() {
	cmd := exec.Command("npm", "run", "tsc")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// -- ripped from https://stackoverflow.com/a/26809999 --
func glob(dir string, ext string, depth int) ([]string, error) {
	depth += strings.Count(dir, "/") // adds basis depth

	files := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if f != nil && f.IsDir() && strings.HasPrefix(path, "node_modules") {
			return filepath.SkipDir
		}

		if strings.HasPrefix(path, ".") {
			return nil
		}

		if f != nil && f.IsDir() {
			return nil
		}

		if depth != -1 && strings.Count(filepath.Dir(path), "/") > depth-1 {
			return nil
		}

		if filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})

	sort.Strings(files)
	return files, err
}

// -- ripped from https://github.com/iancoleman/strcase --
var numberSequence = regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
var numberReplacement = []byte(`$1 $2 $3`)

func addWordBoundariesToNumbers(s string) string {
	b := []byte(s)
	b = numberSequence.ReplaceAll(b, numberReplacement)
	return string(b)
}

func toCamelInitCase(s string, initCase bool) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	capNext := initCase
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' || v == '-' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}

func toCamel(s string) string {
	return toCamelInitCase(s, true)
}
