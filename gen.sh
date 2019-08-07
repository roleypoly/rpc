#!/bin/sh

DIRS=$(/bin/ls -1 -d */)

for d in $DIRS; do
  pushd $d

    protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. *.proto

  popd
done