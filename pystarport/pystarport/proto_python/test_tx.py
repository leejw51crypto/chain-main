#!/usr/bin/python3
import tendermint.rpc.grpc.types_pb2_grpc
import cosmos.tx.v1beta1.tx_pb2
import grpc
import json
from cosmos.tx.v1beta1.tx_pb2 import Tx, TxBody, AuthInfo,TxRaw

def send(tx):
    channel= grpc.insecure_channel('localhost:9091')
    stub= tendermint.rpc.grpc.types_pb2_grpc.BroadcastAPIStub(channel)
    request = tendermint.rpc.grpc.types_pb2.RequestBroadcastTx()
    request.tx = tx
    response = stub.BroadcastTx(request)
    print(f'reponse={response}')


print("OK")
f = open("sign.txt", "rb")
tx=f.read()
a=tx.decode('utf-8')
b=json.loads(a)
c=json.dumps(b, indent=4)
print(c)

txmessages=['0','1']
txbody= TxBody(memo="abc", messages=txmessages)
txsignatures= [b'1', b'2']
txauth= AuthInfo()
new_tx= Tx(body=txbody, auth_info= txauth,signatures= txsignatures) 
print(new_tx)
#print(json.dumps(tx_json, indent=4))
#print(tx) 
send(new_tx.SerializeToString())
#tx =Tx(body=TxBody(memo="ok"))
#send(tx.SerializeToString())
#print(tx)