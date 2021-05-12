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


def massive_test():
    print("massive nft generation")
    s = 0
    n = 350000
    # n = 10
    start_time = time.time()
    for i in range(s, n):
        a = f'token{i}'
        m = hashlib.sha256()
        t0 = datetime.datetime.now().strftime("%d-%b-%Y (%H:%M:%S.%f)")
        t = t0.encode('utf-8')
        m.update(a.encode('utf-8'))
        m.update(t)
        tokenid = (f'mytoken{m.digest().hex()[:40]}')
        create_token(tokenid)
        elasped = time.time() - start_time
        report = f'{i+1}/{n} {(i+1) /n*100}% time: {t0}  tokenid {tokenid} elapsed {elasped} seconds'


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
    f = open("tx.json", "rb")
    b = json.load(f)
    ret = (b["sequence"],  b["account_number"])
    return ret


def get_tx():
    newtokenid = get_newtokenid()
    cmd = f'chain-maind tx nft mint mydenomid {newtokenid} -y --uri="https://ipfs.io/ipfs/QmURd2wtgsLSUzkLUXfv2y89jbKz3isPtwZoZJAoWytYUt?filename=uri.json"  --recipient=$S2 --from=$S1 --chain-id=testnet-central-0 --fees=2basecro --keyring-backend test --name=mytokenname --generate-only  > tx0.txt'
    os.system(f". ./setup.sh && {cmd}")
    a = open("tx0.txt", "r")
    ret = a.read()
    a.close()
    return json.loads(ret)


def write_txs_to_sign():
    ret0 = get_tx()
    f = open("tx.txt", "w")
    for i in range(0, total):
        ret = json.loads(json.dumps(ret0))
        id = ret["body"]["messages"][0]["id"]
        ret["body"]["messages"][0]["id"] = f'{id}0000{i+1}'
        f.write(f'{json.dumps(ret)}\n')
    f.close()


def batch_sign_txs():
    (mysequence, myaccountnumber) = get_sequence_account_number()
    cmd = f'. ./setup.sh && chain-maind tx sign-batch  tx.txt --chain-id $CHAINID --keyring-backend $KEYRING --from $S1 --sequence {mysequence} --account-number {myaccountnumber} > sign.txt'
    a = os.system(cmd)


def get_auth_info():
    (mysequence, myaccountnumber) = get_sequence_account_number()
    cmd = f'. ./setup.sh && chain-maind tx sign tx.txt --chain-id $CHAINID --keyring-backend $KEYRING --from $S1 --sequence {mysequence} --account-number {myaccountnumber} > sign1.txt'
    a = os.system(cmd)
    f = open("sign1.txt", "r")
    ret = json.loads(f.read())["auth_info"]
    f.close()
    return ret


def fill_final_tx():
    f = open("tx.txt", "r")
    alltx = json.loads(f.readline())
    f.close()
    alltx["body"]["messages"].clear()

    f = open("tx.txt", "r")
    for i in range(0, total):
        onetx = json.loads(f.readline())["body"]["messages"][0]
        alltx["body"]["messages"].append(onetx)
    f.close()

    auth_info = get_auth_info()
    print(auth_info)
    signer_info = auth_info["signer_infos"][0]

    f = open("sign.txt", "r")
    for i in range(0, total):
        onesig0 = json.loads(f.readline())["signatures"][0]
        onesig = onesig0["data"]["single"]["signature"]
        alltx["signatures"].append(onesig)
        oneauth = json.loads(json.dumps(signer_info))
        oneauth["public_key"] = onesig0["public_key"]
        oneauth["sequence"] = onesig0["sequence"]

        if len(alltx["auth_info"]["signer_infos"]) == 0:
            alltx["auth_info"]["signer_infos"].append(oneauth)
    f.close()

    alltx["auth_info"]["fee"]["amount"][0]["amount"] = str(2*total)
    print(json.dumps(alltx, indent=4))
    f = open("finaltx.txt", "w")
    f.write(json.dumps(alltx))
    f.write("\n")
    f.close()


total = 2
print("massive nft test")
write_txs_to_sign()
batch_sign_txs()
fill_final_tx()
