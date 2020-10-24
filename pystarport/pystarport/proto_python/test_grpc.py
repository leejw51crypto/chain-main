import cosmos.staking.v1beta1.query_pb2 
import cosmos.staking.v1beta1.query_pb2_grpc
import grpc

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
    
    
print("go!")
channel= grpc.insecure_channel('localhost:9090')
stub= cosmos.staking.v1beta1.query_pb2_grpc.QueryStub(channel)
response=stub.Validators( cosmos.staking.v1beta1.query_pb2.QueryValidatorsRequest())
print(f'response={response}')