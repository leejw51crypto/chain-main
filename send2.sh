. ./setup.sh
FROM=$S1
TO=$S2
AMOUNT=1basecro
chain-maind tx bank  send $FROM $TO $AMOUNT --chain-id $CHAINID --keyring-backend $KEYRING  --generate-only
