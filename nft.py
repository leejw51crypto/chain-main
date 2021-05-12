#!/usr/bin/python3
import hashlib
import os
import datetime
import time
import json
import subprocess


def create_token(tokenid):
    cmd = f'chain-maind tx nft mint mydenomid {tokenid} -y --uri="https://ipfs.io/ipfs/QmURd2wtgsLSUzkLUXfv2y89jbKz3isPtwZoZJAoWytYUt?filename=uri.json"  --recipient=$S2 --from=$S1 --chain-id=testnet-central-0 --fees=2basecro --keyring-backend test --name=mytokenname'
    os.system(f". ./setup.sh && {cmd}")


def get_newtokenid():
    i = 0
    t0 = datetime.datetime.now().strftime("%d-%b-%Y (%H:%M:%S.%f)")
    t = t0.encode('utf-8')
    a = f'token{i}'
    m = hashlib.sha256()
    m.update(a.encode('utf-8'))
    m.update(t)
    tokenid = (f'mytoken{m.digest().hex()[:40]}')
    return tokenid


def get_sequence_account_number():
    cmd = f'. ./setup.sh && chain-maind query account $S1 --output json > tx.json'
    a = os.system(cmd)
    with open("tx.json", "rb") as f:
        b = json.load(f)
        ret = (b["sequence"],  b["account_number"])
    return ret


def get_tx():
    newtokenid = get_newtokenid()
    cmd = f'chain-maind tx nft mint mydenomid {newtokenid} -y --uri="https://ipfs.io/ipfs/QmURd2wtgsLSUzkLUXfv2y89jbKz3isPtwZoZJAoWytYUt?filename=uri.json"  --recipient=$S2 --from=$S1 --chain-id=testnet-central-0 --fees=2basecro --keyring-backend test --name=mytokenname --generate-only  > tx0.txt'
    os.system(f". ./setup.sh && {cmd}")
    with open("tx0.txt", "r") as a:
        ret = a.read()
    return json.loads(ret)


def write_txs_to_sign(total):
    ret0 = get_tx()
    with open("tx.txt", "w") as f:
        for i in range(0, total):
            ret = json.loads(json.dumps(ret0))
            id = ret["body"]["messages"][0]["id"]
            ret["body"]["messages"][0]["id"] = f'{id}0000{i+1}'
            f.write(f'{json.dumps(ret)}\n')


def batch_sign_txs():
    (mysequence, myaccountnumber) = get_sequence_account_number()
    cmd = f'. ./setup.sh && chain-maind tx sign-batch  tx.txt --chain-id $CHAINID --keyring-backend $KEYRING --from $S1 --sequence {mysequence} --account-number {myaccountnumber}  --signature-only=false > sign.txt'
    os.system(cmd)


def process_txs(total):

    txs = []
    with open("sign.txt", "r") as f:
        while True:
            jsondata = f.readline()
            if jsondata == "":
                break
            onetx = json.loads(jsondata)
            txs.append(onetx)

    for i in range(0, len(txs)):
        onetx = txs[i]
        with open("send.txt", "w") as f:
            f.write(f'{json.dumps(onetx)}\n')
        cmd = f'chain-maind tx broadcast send.txt --broadcast-mode async > /dev/null'
        os.system(cmd)


def make_massive(total):
    write_txs_to_sign(total)
    batch_sign_txs()
    process_txs(total)


print("massive nft generation")
s = 0
n = 350000
unit = 50
start_time = time.time()
i = 0
while True:
    make_massive(unit)
    (mysequence, myaccountnumber) = get_sequence_account_number()
    print(f'current sequence {mysequence} account {myaccountnumber}')
    i = int(mysequence)
    elasped = time.time() - start_time
    t0 = datetime.datetime.now().strftime("%d-%b-%Y (%H:%M:%S.%f)")
    report = f'{i+1}/{n} {(i+1) /n*100}% time: {t0}  elapsed {elasped} seconds'
    print(report)
    if i > n:
        break
