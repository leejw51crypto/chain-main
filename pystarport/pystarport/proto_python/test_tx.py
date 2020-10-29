#!/usr/bin/python3
import tendermint.rpc.grpc.types_pb2_grpc
import cosmos.tx.v1beta1.tx_pb2
import grpc
import json
import base64
from cosmos.tx.v1beta1.tx_pb2 import Tx, TxBody, AuthInfo,TxRaw

def send(tx):
    channel= grpc.insecure_channel('localhost:9091')
    stub= tendermint.rpc.grpc.types_pb2_grpc.BroadcastAPIStub(channel)
    request = tendermint.rpc.grpc.types_pb2.RequestBroadcastTx()
    request.tx = tx
    response = stub.BroadcastTx(request)
    print(f'reponse={response}')

def show() :
    print("OK")
    f = open("sign.txt", "rb")
    tx=f.read()
    a=tx.decode('utf-8')
    b=json.loads(a)
    c=json.dumps(b, indent=4)
    print(c)


f = open("encode.txt", "rb")
tx=f.read()
a=tx.decode('utf-8')
print(a)
b=base64.b64decode(a)
print(b)
send(b)