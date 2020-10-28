#!/usr/bin/python3
import tendermint.rpc.grpc.types_pb2_grpc
import grpc

channel= grpc.insecure_channel('localhost:9091')
stub= tendermint.rpc.grpc.types_pb2_grpc.BroadcastAPIStub(channel)
request = tendermint.rpc.grpc.types_pb2.RequestBroadcastTx()
print(f'request={request}')
response = stub.BroadcastTx(request)
print(f'response={response}')
