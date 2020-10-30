from datetime import timedelta
import pytest
from dateutil.parser import isoparse
from pystarport.ports import rpc_port
from .utils import parse_events, wait_for_block_time, wait_for_port, wait_for_new_blocks

import cosmos.staking.v1beta1.query_pb2 
import cosmos.staking.v1beta1.query_pb2_grpc
import cosmos.bank.v1beta1.tx_pb2
import cosmos.bank.v1beta1.tx_pb2_grpc
import grpc

def test_query_validators(cluster):    
    print("go!")
    wait_for_new_blocks(cluster,5)
    baseport = cluster.base_port(0)
    print(f'baseport={baseport}')
    ipport = cluster.ipport_grpc(0)
    print(f'grpc port={ipport}')
    channel= grpc.insecure_channel(ipport)
    stub= cosmos.staking.v1beta1.query_pb2_grpc.QueryStub(channel)
    response=stub.Validators( cosmos.staking.v1beta1.query_pb2.QueryValidatorsRequest())
    print(f'response={response}')
    assert  False
