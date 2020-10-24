#!/bin/bash
protoc --proto_path=./proto --proto_path=./third_party/proto --python_out=./proto_python/ $(find ./proto/cosmos -iname "*.proto") --grpc_python_out=./proto_python  --plugin=protoc-gen-grpc_python=$HOME/bin/grpc_python_plugin
protoc --proto_path=./third_party/proto --proto_path=./proto --python_out=./proto_python/ $(find ./third_party/proto -iname "*.proto") --grpc_python_out=./proto_python  --plugin=protoc-gen-grpc_python=$HOME/bin/grpc_python_plugin

