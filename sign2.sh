. ./setup.sh
chain-maind tx sign-batch  tx.txt --chain-id $CHAINID --keyring-backend $KEYRING --from $S1 --sequence 3 --account-number 2  --offline
