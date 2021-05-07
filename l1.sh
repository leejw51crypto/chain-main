#!/bin/bash
chain-maind  query bank balances  $L1 --node tcp://mynode:26657

echo "================================================================"
chain-maind query staking delegation $L1  $VV1
