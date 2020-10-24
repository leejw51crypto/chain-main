#!/bin/bash
protoc --proto_path=./proto --proto_path=./third_party/proto --js_out=./proto_js/ $(find ./proto/cosmos -iname "*.proto") --grpc_out=./proto_js  --plugin=protoc-gen-grpc=$HOME/bin/grpc_node_plugin
protoc --proto_path=./third_party/proto --proto_path=./proto --js_out=./proto_js/ $(find ./third_party/proto -iname "*.proto") --grpc_out=./proto_js --plugin=protoc-gen-grpc=$HOME/bin/grpc_node_plugin

