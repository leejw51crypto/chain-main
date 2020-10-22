import cosmos.staking.v1beta1.query_pb2 

a = cosmos.staking.v1beta1.query_pb2.QueryValidatorRequest()
print("grpc test")
print(f'request={a.SerializeToString()}')