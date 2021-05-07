#!/bin/bash
. ./setup.sh
chain-maind  query bank balances  $T2 --node tcp://mynode:26657
#echo  "============================================================="
