FROM golang:1.13-alpine

WORKDIR /src

RUN apk add --no-cache nodejs protoc protobuf protobuf-dev git npm && \
    go get -u github.com/golang/protobuf/...

COPY package.json package-lock.json /src/
RUN npm ci

CMD ["go", "generate"]