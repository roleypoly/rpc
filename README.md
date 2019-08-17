# roleypoly rpc

This package includes automatically generated files based on protobuf definition files. Currently, gRPC Go, gRPC-Web JS and TS definitions exist here. Nothing here is a functional part of an app, just data wiring.

## Importing

- **Go:** `go get -u github.com/roleypoly/rpc`
- **JS/TS:** `yarn add @roleypoly/rpc`

In both cases, you may directly import a sub-directory, e.g. `@roleypoly/rpc/discord` or `github.com/roleypoly/rpc/discord`, which resolves to the generated protobuf file.

## Developing

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

If more are needed, please add them to `./hack/gen.sh`.

## Using

You are free to use my definitions as you see fit. They might be useful in your app, or they might not. 