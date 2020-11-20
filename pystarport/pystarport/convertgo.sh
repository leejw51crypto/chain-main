#!/bin/bash
OUTPUT=./proto_go
COSMOS=./cosmos-sdk
TENDERMINT=./tendermint
TMP=$(whereis grpc_python_plugin)
PLUGIN="$(cut -d' ' -f2 <<<"$TMP")"
mkdir $OUTPUT
git clone --branch v0.40.0-rc2 https://github.com/cosmos/cosmos-sdk.git
git clone --branch v0.34.0-rc5 https://github.com/tendermint/tendermint.git
cp -Rf $COSMOS/third_party/proto/* $COSMOS/proto/ 
# cosmos
protoc --proto_path=$COSMOS/proto --proto_path=$COSMOS/third_party/proto --go_out=$OUTPUT $(find $COSMOS/proto/cosmos -iname "*.proto") --go-grpc_out=$OUTPUT  
# cosmos third-party
protoc --proto_path=$COSMOS/third_party/proto --proto_path=$COSMOS/proto --go_out=$OUTPUT $(find $COSMOS/third_party/proto -iname "*.proto") --go-grpc_out=$OUTPUT  
# tendermint
protoc --proto_path=$TENDERMINT/proto --proto_path=$TENDERMINT/third_party/proto --go_out=$OUTPUT $(find $TENDERMINT/proto/tendermint -iname "*.proto") --go-grpc_out=$OUTPUT  
