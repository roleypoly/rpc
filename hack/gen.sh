#!/bin/sh

DIRS=$(/bin/ls -1 -d */)
ROOT=$(pwd)
PROTOC_GEN_TS_PATH=$ROOT/node_modules/.bin/protoc-gen-ts

echo 'module.exports = {}' | tee $ROOT/index.js >/dev/null
echo '' | tee $ROOT/index.d.ts >/dev/null

write_node_files () {
  echo "import * as $1 from './$1'; export {$1}" \
   | tee -a $ROOT/index.d.ts >/dev/null
  echo "module.exports.$1 = require('./$1')" \
   | tee -a $ROOT/index.js >/dev/null

  echo "module.exports = { ...require('./$1_pb'), ...require('./$1_pb_service') }" \
   | tee index.js >/dev/null
  echo "export * from './$1_pb'; export * from './$1_pb_service'" \
   | tee index.d.ts >/dev/null
}

for d in $DIRS; do
  name=$(basename $d)
  pushd $d >/dev/null

    find *.proto 2>/dev/null 1>/dev/null && {
      echo "[$name] generating..."
      echo "[$name] ... protoc"
      protoc \
        --proto_path=$GOPATH/src:. \
        --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
        --js_out="import_style=commonjs,binary:." \
        --ts_out="service=true:." \
        --go_out="plugins=grpc:." \
          *.proto

      echo "[$name] ... ts/js index"
      write_node_files $name
      echo "[$name] done."
    }

  popd >/dev/null
done