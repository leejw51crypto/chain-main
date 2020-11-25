# ledger test 


## setup chain

### config
account address == validator address , only different hrp 
```
const (
	CoinType       = 394
	FundraiserPath = "44'/394'/0'/0/1"
)

var (
	AccountAddressPrefix   = "cro"
	AccountPubKeyPrefix    = "cropub"
	ValidatorAddressPrefix = "crocncl"
	ValidatorPubKeyPrefix  = "crocnclpub"
	ConsNodeAddressPrefix  = "crocnclcons"
	ConsNodePubKeyPrefix   = "crocnclconspub"
	HumanCoinUnit          = "cro"
	BaseCoinUnit           = "basecro" // 10^-8 AKA "carson"
	CroExponent            = 8
)
```

### prepare
go.sh
```
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

chain-maind gentx validator1 100000000cro  --chain-id test  --keyring-backend test

chain-maind collect-gentxs

chain-maind validate-genesis

```
### setup
#### setup.sh
$V1, $VV1 is the same address
$VV1: $V1's validator address  starting with crocncl
```
export V1=$(chain-maind keys show validator1 -a --keyring-backend test)
export V2=$(chain-maind keys show validator2 -a --keyring-backend test) 
export S1=$(chain-maind keys show signer1 -a --keyring-backend test) 
export S2=$(chain-maind keys show signer2 -a --keyring-backend test)
export VV1=crocncl1t8etx0hzfrpxl0yk8nawupvz6lhyh85nhsjqlt

echo V1= $V1
echo V2= $V2
echo S1= $S1
echo S2= $S2
echo VV1= $VV1
```
#### setup3.sh
$L1 is your ledger address
```
export CHAINID=test
export L1=cro1l320uqmk82nu432c0xuz54ktf4f9qmmqd8nhmr 
```

### show info
s1.sh
```
#!/bin/bash
chain-maind  query bank balances  $S1 --node tcp://mynode:26657
echo  "============================================================="
chain-maind query staking delegation $S1  $VV1
```
s2.sh
```
#!/bin/bash
. ./setup.sh
chain-maind  query bank balances  $S2
```

### send funds
#### s1->s2

send.sh
```
. ./setup.sh
chain-maind tx bank  send $S1 $S2 1basecro  --chain-id test --keyring-backend test

```
####  sendl.sh
send funds to the ledger address
```
. ./setup.sh
. ./setup3.sh
chain-maind tx bank  send $S1 $L1 10000basecro  --chain-id test --keyring-backend test
```


### staking
staking.sh
```
chain-maind tx staking delegate $VV1 1basecro --from $S1 --keyring-backend test --chain-id test
```

### unbonding
just for reference   
unbond.sh   
```
chain-maind tx staking unbond $VV1 1basecro --from $S1 --keyring-backend test --chain-id test

```


## modify config
enable cors && bind 0.0.0.0
### app.toml
cd $HOME/.chain-maind
cd config
vi app.toml
```
enabled-unsafe-cors = true
enable = true# Swagger defines if swagger documentation should automatically be registered.
swagger = true# Address defines the API server to listen on
address = "tcp://0.0.0.0:1317"# Address defines the gRPC server address to bind to.
address = "0.0.0.0:9090"
```
### config.toml
vi config.toml
```
# TCP or UNIX socket address for the RPC server to listen on
laddr = "tcp://0.0.0.0:26657"
cors_allowed_origins = ["*"]
```

## install ledger live
settting -> experimental feature-> developer mode  `on`

## install firmware
1. manager-> app catalog-> crypto.com chain 
2. done!!


## compile desktop chain wallet and run
1. install nodejs 14.x
2. npm -g yarn
3. git clone git@github.com:crypto-com/chain-desktop-wallet.git
   git checkout ledger-support
4. sudo apt install libudev-dev   (for linux, for mac skip)
5. yarn
6. yarn electron:dev 
7. yarn electron:windev (for windows)
8. open http://localhost:3000/ in chrome
   
## open chrome
1. enter url : chrome://version//flags
2. enable experimentail webplatform feature 
3. enable new udb backend


## create wallet first
as normal wallet

## create ledger wallet
1. click create wallet
2. use hardware wallet
3. choose ledger
4. choose devnet
```
nodeurl: http://yournodeip
indexing url: https://chain.crypto.com/explorer/api/v1/
derivation path: m/44'/394'/0'/0/0
validator prefix: crocncl
address prefix: cro
chainid: test
base denom: basecro
cro denom: cro
```
5. all set
   
## test
1. send
2. staking
   

## for electron
 1. turn on native-hid mode
vi ./src/service/LedgerService.ts

change to false
```
export const useWebusbForLedger = false;
```
2. recompile
```
yarn
yarn electron:dev
```
3. use opened electron-window  
testing procedure is the same  
