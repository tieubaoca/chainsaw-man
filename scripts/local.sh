#!/bin/bash
VALIDATOR="validator"
KEY1="tieubaoca01"
KEY2="tieubaoca02"
KEY3="tieubaoca03"
BALANCE1="1000000000usae"
BALANCE2="50000000usae"
CHAINID="toddler"
KEYRING="test"
MONIKER="localtestnet"
LOGLEVEL="info"


rm -rf $HOME/.saw*
make install
sawd config keyring-backend $KEYRING
sawd config chain-id $CHAINID
# init chain
sawd init $MONIKER --chain-id $CHAINID
sawd keys add $KEY1 --keyring-backend $KEYRING
sawd keys add $KEY2 --keyring-backend $KEYRING
sawd keys add $KEY3 --keyring-backend $KEYRING
sawd keys add $VALIDATOR --keyring-backend $KEYRING
# Change parameter token denominations to usae
cat $HOME/.saw/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="usae"' > $HOME/.saw/config/tmp_genesis.json && mv $HOME/.saw/config/tmp_genesis.json $HOME/.saw/config/genesis.json
cat $HOME/.saw/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="usae"' > $HOME/.saw/config/tmp_genesis.json && mv $HOME/.saw/config/tmp_genesis.json $HOME/.saw/config/genesis.json
cat $HOME/.saw/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="usae"' > $HOME/.saw/config/tmp_genesis.json && mv $HOME/.saw/config/tmp_genesis.json $HOME/.saw/config/genesis.json
cat $HOME/.saw/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="usae"' > $HOME/.saw/config/tmp_genesis.json && mv $HOME/.saw/config/tmp_genesis.json $HOME/.saw/config/genesis.json

# enable rest server and swagger
toml set --toml-path $HOME/.saw/config/app.toml api.swagger true
toml set --toml-path $HOME/.saw/config/app.toml api.enable true
# toml set --toml-path $HOME/.saw/config/app.toml api.address tcp://0.0.0.0:1350

# Allocate genesis accounts (cosmos formatted addresses) 
sawd add-genesis-account $KEY1 $BALANCE1 --keyring-backend $KEYRING
sawd add-genesis-account $KEY2 $BALANCE1 --keyring-backend $KEYRING
sawd add-genesis-account $KEY3 $BALANCE2 --keyring-backend $KEYRING
sawd add-genesis-account $VALIDATOR $BALANCE1 --keyring-backend $KEYRING

# Sign genesis transaction
sawd gentx $VALIDATOR 10000000usae --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
sawd collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
sawd validate-genesis

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
sawd start --pruning=nothing --log_level $LOGLEVEL --minimum-gas-prices=0.0001usae
# do it man :))))