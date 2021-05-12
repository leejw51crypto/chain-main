. ./setup0.sh
export V1=$(chain-maind keys show validator1 -a --keyring-backend test)
export V2=$(chain-maind keys show validator2 -a --keyring-backend test) 
export S1=$(chain-maind keys show signer1 -a --keyring-backend test) 
export S2=$(chain-maind keys show signer2 -a --keyring-backend test)
export T1=$(chain-maind keys show test1 -a --keyring-backend test) 
export T2=$(chain-maind keys show test2 -a --keyring-backend test)

