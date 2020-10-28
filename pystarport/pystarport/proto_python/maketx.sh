#!/bin/bash
chain-maind tx bank send $SIGNER1  $SIGNER2 1basecro  --keyring-backend test --chain-id test --generate-only > tx.txt
chain-maind tx sign ./tx.txt  --keyring-backend test --chain-id test --from $SIGNER1 > sign.txt
chain-maind tx sign ./tx.txt  --keyring-backend test --chain-id test --from $SIGNER1 --sign-mode direct > sign2.txt
cat sign.txt
cp sign.txt ~
chain-maind tx encode sign.txt > encode.txt
cp encode.txt ~  
#chain-maind tx broadcast sign.txt
