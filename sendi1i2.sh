rly q balance test -i 
rly q balance test2 -i 
rly tx transfer test test2 11basecro $(rly keys show test2)
echo "--------------------------------------------------"
sleep 5
./relay.sh
rly q balance test -i
rly q balance test2 -i

