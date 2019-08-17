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
}

type langconfig struct {
	Go bool
	JS bool
	TS bool
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

	flags := getProtocFlags(config)
	if len(flags) == 0 {
		return errors.New("No valid configuration was present.")
	}

	err = runProtoc(dir, flags...)
	if err != nil {
		return err
	}

	if config.JS {
		err = generateJs(dir)
	}

	return err
}

func generateJs(dir string) error {
	globalJSExports = append(globalJSExports, dir)
	protos := getLocalProtoNames(dir)

	exports := []string{}
	for _, protoName := range protos {
		exports = append(
			exports,
			fmt.Sprintf("./%s_pb", protoName),
			fmt.Sprintf("./%s_pb_service", protoName),
		)
	}

	jsContents := strings.Builder{}
	tsContents := strings.Builder{}

	jsContents.WriteString("module.exports = {\n")
	for _, export := range exports {
		safeName := strings.Replace(filepath.Base(export), ".proto", "", -1)
		jsContents.WriteString(fmt.Sprintf("  ...require('./%s'),\n", safeName))
		tsContents.WriteString(fmt.Sprintf("export * from './%s'\n", safeName))
	}
	jsContents.WriteString("}\n\n")
	tsContents.WriteString("\n")

	err := ioutil.WriteFile(filepath.Join(dir, "index.js"), []byte(jsContents.String()), 0666)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(dir, "index.d.ts"), []byte(tsContents.String()), 0666)
	if err != nil {
		return err
	}
	return nil
}

func writeGlobalJS() error {
	jsContents := strings.Builder{}
	tsContents := strings.Builder{}

	jsContents.WriteString("module.exports = {\n")
	for _, export := range globalJSExports {
		jsName := toCamel(strings.Replace(export, "/", "_", -1))
		tsContents.WriteString(fmt.Sprintf("import * as %s from './%s'\n", jsName, export))
		jsContents.WriteString(fmt.Sprintf("  %s: require('./%s'),\n", jsName, export))
	}
	tsContents.WriteString("\n")
	jsContents.WriteString("}\n\n")

	err := ioutil.WriteFile(filepath.Join(rootDir, "index.js"), []byte(jsContents.String()), 0666)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(rootDir, "index.d.ts"), []byte(tsContents.String()), 0666)
	if err != nil {
		return err
	}

	return nil
}

func getRelevantDirs() ([]string, error) {
	dirs := []string{}

	matches, err := glob(".", ".genconfig")
	if err != nil {
		return dirs, err
	}

	for _, match := range matches {
		dirs = append(
			dirs,
			filepath.Clean(
				strings.Replace(match, ".genconfig", "", -1),
			),
		)
	}

	return dirs, nil
}

func getProtocFlags(lc langconfig) []string {
	flags := []string{}

	if lc.JS {
		flags = append(flags, "--js_out=import_style=commonjs,binary:.")
	}

	if lc.TS {
		flags = append(flags, "--ts_out=service=true:.")
	}

	if lc.Go {
		flags = append(flags, "--go_out=plugins=grpc:.")
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

	jsFiles := assertGlob(glob(dir, ".js"))
	goFiles := assertGlob(glob(dir, ".go"))
	tsFiles := assertGlob(glob(dir, ".ts"))

	files := append(jsFiles, goFiles...)
	files = append(files, tsFiles...)

	for _, file := range files {
		log.Println(dir, file)
		err := os.Remove(file)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func getLocalProtoNames(dir string) []string {
	return assertGlob(glob(dir, ".proto"))
}

func runProtoc(dir string, flags ...string) error {
	log.Printf("[%s] ... protoc\n", dir)

	args := []string{
		// "--proto-path=$GOPATH/src:.",
		fmt.Sprintf("--plugin=protoc-gen-ts=%s/node_modules/.bin/protoc-gen-ts", rootDir),
	}

	err := os.Chdir(dir)
	if err != nil {
		return err
	}

	files, err := glob(".", ".proto")
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

	err = os.Chdir(rootDir)
	return err
}

// -- ripped from https://stackoverflow.com/a/26809999 --
func glob(dir string, ext string) ([]string, error) {

	files := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if strings.HasPrefix(path, "node_modules") || strings.HasPrefix(path, ".") {
			return nil
		}
		if filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})

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
