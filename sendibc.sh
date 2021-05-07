./s1.sh
chain-maind tx bank send $I2 $S1 2ibc/38844A85644DEB18B4F00A470720AAC2490A6E90A815D61E79A0C20A3B0E0565   --chain-id test2 --keyring-backend test
echo "-------------------------------------------"
sleep 2
./s1.sh
