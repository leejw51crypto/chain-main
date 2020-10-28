#!/usr/bin/python3
import cosmos.staking.v1beta1.query_pb2 
import cosmos.staking.v1beta1.query_pb2_grpc
import cosmos.bank.v1beta1.tx_pb2
import cosmos.bank.v1beta1.tx_pb2_grpc
import grpc
from cosmos.base.v1beta1.coin_pb2 import Coin
def test() :
    a = cosmos.staking.v1beta1.query_pb2.QueryValidatorRequest()
    a.validator_addr="OK"
    print("grpc test")
    print(f'request={a.SerializeToString()}')

    b = cosmos.staking.v1beta1.query_pb2.QueryValidatorResponse()
    b.validator.jailed=False
    b.validator.operator_address="a1"
    b.validator.unbonding_height=20
    b.validator.min_self_delegation="b2"
    print(f'b={b}')
    print(f"response={b.SerializeToString()}")
    

def validators():    
    print("go!")
    channel= grpc.insecure_channel('localhost:9090')
    stub= cosmos.staking.v1beta1.query_pb2_grpc.QueryStub(channel)
    response=stub.Validators( cosmos.staking.v1beta1.query_pb2.QueryValidatorsRequest())
    print(f'response={response}')
    

print("go2")
channel= grpc.insecure_channel('localhost:9091')
stub= cosmos.bank.v1beta1.tx_pb2_grpc.MsgStub(channel)
request = cosmos.bank.v1beta1.tx_pb2.MsgSend()
#(from_address="cro1n3yen4f88e4agtx0tz72g6jdcaaptqrvjaqey7",to_address="cro1d2r3lnwtfdkmmlhstjfnsn3utfrgjmgx2am424",amount=Coin("2basecro"))
response = stub.Send(request)
print(f'response={response}')
