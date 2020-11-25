#!/bin/bash
OUTPUT=./proto_python
COSMOS=./cosmos-sdk
TENDERMINT=./tendermint
TMP=$(whereis grpc_python_plugin)
PLUGIN="$(cut -d' ' -f2 <<<"$TMP")"
echo $PLUGIN
