# roleypoly rpc

This package includes automatically generated files based on protobuf definition files. Currently, gRPC Go, gRPC-Web JS and TS definitions exist here. Nothing here is a functional part of an app, just data wiring.

## Importing

- **Go:** `go get -u github.com/roleypoly/rpc`
- **JS/TS:** `yarn add @roleypoly/rpc`

In both cases, you may directly import a sub-directory, e.g. `@roleypoly/rpc/discord` or `github.com/roleypoly/rpc/discord`, which resolves to the generated protobuf file.

## Developing

Make your changes **only** to .proto files, run `go generate` to update definitions. All `index.js` and `index.d.ts` files are generated for you.

To update definitions, you'll need to run `yarn` for **protoc-gen-ts**, and fetch **protoc**, and **protoc-gen-go** on your own.

## Using

You are free to use my definitions as you see fit. They might be useful in your app, or they might not. 