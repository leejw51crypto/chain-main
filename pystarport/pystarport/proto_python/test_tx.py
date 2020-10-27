#!/usr/bin/python3
from cosmospy import generate_wallet, Transaction,seed_to_privkey,privkey_to_address
import json


unit="cro"
cro_path="m/44'/394'/0'/0/{}"

wallet = generate_wallet(path=cro_path.format(1),hrp="cro")
seed= wallet["seed"]
print(f'wallet={wallet}')


privkey = seed_to_privkey(seed, path=cro_path.format(1))
b_addr=privkey_to_address(privkey, hrp=unit)
print(f'b={b_addr}')

privkey2 = seed_to_privkey(seed, path=cro_path.format(2))
b2_addr=privkey_to_address(privkey2, hrp=unit)
print(f'b2={b2_addr}')



tx = Transaction(
    privkey=privkey,
    account_num=11335,
    sequence=0,
    fee=1000,
    gas=70000,
    memo="",
    chain_id="chain-test",
    sync_mode="sync",
)
tx.add_transfer(
    recipient=b2_addr, amount=123
)
pushable_tx = tx.get_pushable()
a= json.dumps(pushable_tx, indent=4)
print(a)
