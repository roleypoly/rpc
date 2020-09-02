FROM golang:1.15.1-alpine

WORKDIR /src

COPY package.json package-lock.json go.mod go.sum hack/gen.go ./

RUN apk add --no-cache nodejs protoc protobuf protobuf-dev git npm && \
    go get -u github.com/golang/protobuf/... && \
    npm ci && \
    go mod download && \
    go build -o /bin/rpc-gen /src/gen.go

CMD ["/bin/rpc-gen"] 
