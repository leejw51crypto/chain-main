#!/usr/bin/python3
from cosmospy import generate_wallet, Transaction,seed_to_privkey,privkey_to_address
import json


unit="cro"
cro_path="m/44'/394'/0'/0/{}"

#wallet = generate_wallet(path=cro_path.format(1),hrp="cro")
#seed= wallet["seed"]
#print(f'wallet={wallet}')
seed= 'like address into shine excess paddle second pact short century victory skull twenty gossip book curious foster fault source total muffin rocket champion enact'

privkey = seed_to_privkey(seed, path=cro_path.format(1))
b_addr=privkey_to_address(privkey, hrp=unit)
print(f'b= {b_addr}')

privkey2 = seed_to_privkey(seed, path=cro_path.format(2))
b2_addr=privkey_to_address(privkey2, hrp=unit)
print(f'b2= {b2_addr}')



tx = Transaction(
    privkey=privkey,
    account_num=10,
    sequence=0,
    fee=0,
    gas=0,
    memo="",
    chain_id="test",
    sync_mode="sync",
    fee_denom="basecro",
)
tx.add_transfer(
    recipient=b2_addr, amount=123, denom="basecro"
)
pushable_tx = tx.get_pushable()
print(pushable_tx)

a= json.loads(pushable_tx)
print(json.dumps(a, indent=4))

f = open("tx.txt", "wb")
f.write(bytes(pushable_tx,'utf-8'))
f.close()