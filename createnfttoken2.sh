. ./setup.sh
chain-maind tx nft mint mydenomid mytokenid3 --uri="https://ipfs.io/ipfs/QmURd2wtgsLSUzkLUXfv2y89jbKz3isPtwZoZJAoWytYUt?filename=uri.json"  --recipient=$S2 --from=$S1 --chain-id=testnet-central-0 --fees=2basecro --keyring-backend test --name=mytokenname --generate-only

