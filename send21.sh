. ./setup.sh
FROM=$S2
TO=$S1
AMOUNT=1basecro
chain-maind tx bank  send $FROM $TO $AMOUNT --chain-id $CHAINID --keyring-backend $KEYRING 
