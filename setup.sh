export V1=$(chain-maind keys show validator1 -a --keyring-backend test)
export V2=$(chain-maind keys show validator2 -a --keyring-backend test) 
export S1=$(chain-maind keys show signer1 -a --keyring-backend test) 
export S2=$(chain-maind keys show signer2 -a --keyring-backend test)
export I2=$(chain-maind keys show test2 -a --keyring-backend test)
export VV1=crocncl1t8etx0hzfrpxl0yk8nawupvz6lhyh85nhsjqlt

echo V1= $V1
echo V2= $V2
echo S1= $S1
echo S2= $S2
echo VV1= $VV1
