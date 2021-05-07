#!/bin/bash
export CHAIN=chain-maind
export CLI=chain-maind
rm -rf $HOME/.$CHAIN
rm -rf $HOME/.$CLI
$CHAIN unsafe-reset-all

$CHAIN init testnode --chain-id test -o

$CLI keys add validator1 --keyring-backend test
$CLI keys add validator2 --keyring-backend test
$CLI keys add signer1 --keyring-backend test
$CLI keys add signer2 --keyring-backend test

$CHAIN add-genesis-account $($CLI keys show validator1 -a --keyring-backend test) 10000000000000cro
$CHAIN add-genesis-account $($CLI keys show validator2 -a --keyring-backend test) 10000000000000cro
$CHAIN add-genesis-account $($CLI keys show signer1 -a --keyring-backend test) 10000000000cro
$CHAIN add-genesis-account $($CLI keys show signer2 -a --keyring-backend test) 10000000000cro

$CHAIN gentx validator1  100000000cro --keyring-backend test   --chain-id test

$CHAIN collect-gentxs

$CHAIN validate-genesis
