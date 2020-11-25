#!/bin/bash
rm -rf $HOME/.chain-maind
. ./setup3.sh
chain-maind unsafe-reset-all

chain-maind init testnode --chain-id $CHAINID  -o

chain-maind keys add validator1 --keyring-backend test
chain-maind keys add validator2 --keyring-backend test
chain-maind keys add signer1 --keyring-backend test
chain-maind keys add signer2 --keyring-backend test

chain-maind add-genesis-account $(chain-maind keys show validator1 -a --keyring-backend test) 2000000000cro 
chain-maind add-genesis-account $(chain-maind keys show validator2 -a --keyring-backend test) 2000000000cro
chain-maind add-genesis-account $(chain-maind keys show signer1 -a --keyring-backend test) 100000000cro
chain-maind add-genesis-account $(chain-maind keys show signer2 -a --keyring-backend test) 100000000cro

chain-maind gentx validator1 100000000cro  --chain-id $CHAINID --keyring-backend test

chain-maind collect-gentxs

chain-maind validate-genesis
