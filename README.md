# roleypoly rpc

This package includes automatically generated files based on protobuf definition files. Currently, gRPC Go, gRPC-Web JS and TS definitions exist here. Nothing here is a functional part of an app, just data wiring.

Versioning is done automatically via semantic-release, and no stability guarantees will be made in 1.x.

## Importing

- **Go:** `go get -u github.com/roleypoly/rpc`
- **JS/TS:** `yarn add @roleypoly/rpc`

In both cases, you may directly import a sub-directory, e.g. `@roleypoly/rpc/discord` or `github.com/roleypoly/rpc/discord`, which resolves to the generated protobuf file.

## Developing

### With Docker

Make your changes **only** to .proto files, run `./generate.sh` or `./generate.ps1`. You must have Docker installed and available on PATH.

If you need more fresh NPM packages, you can build the docker image locally with `docker build . -t sometag`, and use the `docker run` command within `./generate.ps1`. This should be automatically done during normal CI/CD processes.

### Without Docker

Make your changes **only** to .proto files, run `go generate` to update definitions. All `index.js` and `index.d.ts` files are generated for you.

To update definitions, you'll need to run `yarn` for **protoc-gen-ts**, and fetch **protoc**, and **protoc-gen-go** on your own.

### Making new packages

All packages for this repo need to have a protobuf file, and `.genconfig` with every generator needed on separate lines.

For instance, a shared gRPC-Web and gRPC-Go package will want to have a `.genconfig` with the following contents:
```
go
js
ts
```

These are automatically configured to these arguments to protoc:
- `ts` => `--ts_out=services=true:.`
- `js` => `--js_out=import_style=commonjs,binary:.`
- `go` => `--go_out=plugins=grpc:.`
- `--noService` => Disables service generation for shared libraries without RPCs

If more are needed, please add them to `./hack/gen.go`.

## Using

You are free to use my definitions as you see fit. They might be useful in your app, or they might not. 