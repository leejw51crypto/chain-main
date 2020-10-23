import cosmos.staking.v1beta1.query_pb2 

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