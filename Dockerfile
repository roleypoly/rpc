FROM golang:1.13-alpine

WORKDIR /src

RUN apk add --no-cache nodejs protoc protobuf protobuf-dev git npm && \
    go get -u github.com/golang/protobuf/...

COPY hack /src/hack
COPY package.json package-lock.json go.mod go.sum /src/
RUN npm ci && \
    go mod download && \
    go build -o /bin/rpc-gen /src/hack/gen.go

CMD ["/bin/rpc-gen"] 
