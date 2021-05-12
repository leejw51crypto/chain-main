#!/bin/bash
. ./setup.sh
chain-maind  query bank balances  $T1 --node tcp://mynode:26657
#echo  "============================================================="
