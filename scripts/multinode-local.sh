#!/bin/bash
CHAINID="saw-0"
KEYRING="test"
MONIKER="localtestnet"
LOGLEVEL="info"

VALIDATOR1="validator1"
VALIDATOR2="validator2"
VALIDATOR3="validator3"
BALANCE1="1000000000000usae"

SEED1="upper author page theme spring entire risk receive world film gasp prepare never depart cotton slow stay photo indicate citizen crystal defy glide position"
SEED2="display ranch board mixture morning culture age random middle enjoy obscure effort divorce permit burst shell citizen fun current join ankle early onion cake"
SEED3="candy spice suit select scan gossip silver lens dwarf subway spice sword reward humor fall fruit bleak garment van okay hat prosper motor velvet"

rm -rf $HOME/.saw/
killall screen
killall sawd
make install

sawd config keyring-backend $KEYRING
sawd config chain-id $CHAINID

# start a testnet
sawd init $VALIDATOR1 --home $HOME/.saw/node0 --chain-id $CHAINID
sawd init $VALIDATOR2 --home $HOME/.saw/node1 --chain-id $CHAINID
sawd init $VALIDATOR3 --home $HOME/.saw/node2 --chain-id $CHAINID

echo $SEED1 | sawd keys add $VALIDATOR1 --keyring-backend $KEYRING --home $HOME/.saw/node0 --recover
echo $SEED2 | sawd keys add $VALIDATOR2 --keyring-backend $KEYRING --home $HOME/.saw/node0 --recover
echo $SEED3 | sawd keys add $VALIDATOR3 --keyring-backend $KEYRING --home $HOME/.saw/node0 --recover
echo $SEED2 | sawd keys add $VALIDATOR2 --keyring-backend $KEYRING --home $HOME/.saw/node1 --recover
echo $SEED3 | sawd keys add $VALIDATOR3 --keyring-backend $KEYRING --home $HOME/.saw/node2 --recover

VAL1ADDRESS=$(sawd keys show $VALIDATOR1 --keyring-backend $KEYRING --home $HOME/.saw/node0 --output json | jq '.address')
VAL2ADDRESS=$(sawd keys show $VALIDATOR2 --keyring-backend $KEYRING --home $HOME/.saw/node0 --output json | jq '.address')
VAL3ADDRESS=$(sawd keys show $VALIDATOR3 --keyring-backend $KEYRING --home $HOME/.saw/node0 --output json | jq '.address')

cat $HOME/.saw/node0/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="usae"' > $HOME/.saw/node0/config/tmp_genesis.json && mv $HOME/.saw/node0/config/tmp_genesis.json $HOME/.saw/node0/config/genesis.json
cat $HOME/.saw/node0/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="usae"' > $HOME/.saw/node0/config/tmp_genesis.json && mv $HOME/.saw/node0/config/tmp_genesis.json $HOME/.saw/node0/config/genesis.json
cat $HOME/.saw/node0/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="usae"' > $HOME/.saw/node0/config/tmp_genesis.json && mv $HOME/.saw/node0/config/tmp_genesis.json $HOME/.saw/node0/config/genesis.json
cat $HOME/.saw/node0/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="usae"' > $HOME/.saw/node0/config/tmp_genesis.json && mv $HOME/.saw/node0/config/tmp_genesis.json $HOME/.saw/node0/config/genesis.json

echo "Addresses: " $VAL1ADDRESS $VAL2ADDRESS $VAL3ADDRESS

sawd add-genesis-account $VALIDATOR1 $BALANCE1 --keyring-backend $KEYRING --home $HOME/.saw/node0
sawd add-genesis-account $VALIDATOR2 $BALANCE1 --keyring-backend $KEYRING --home $HOME/.saw/node0
sawd add-genesis-account $VALIDATOR3 $BALANCE1 --keyring-backend $KEYRING --home $HOME/.saw/node0

# Sign genesis transaction
sawd gentx $VALIDATOR1 10000000usae --keyring-backend $KEYRING --chain-id $CHAINID --home $HOME/.saw/node0

# Collect genesis tx
sawd collect-gentxs --home $HOME/.saw/node0

# Run this to ensure everything worked and that the genesis file is setup correctly
sawd validate-genesis --home $HOME/.saw/node0

# change app.toml values

PEER=$(sawd tendermint show-node-id --home $HOME/.saw/node0)

# validator 1
toml set --toml-path $HOME/.saw/node0/config/app.toml minimum-gas-prices "0usae"
toml set --toml-path $HOME/.saw/node0/config/app.toml api.swagger true
toml set --toml-path $HOME/.saw/node0/config/app.toml rosetta.denom-to-suggest "usae"

# validator2
toml set --toml-path $HOME/.saw/node1/config/app.toml api.address "tcp://localhost:1316"
toml set --toml-path $HOME/.saw/node1/config/app.toml grpc.address "0.0.0.0:9088"
toml set --toml-path $HOME/.saw/node1/config/app.toml grpc-address "0.0.0.0:9089"
toml set --toml-path $HOME/.saw/node1/config/app.toml minimum-gas-prices "0usae"
toml set --toml-path $HOME/.saw/node1/config/app.toml api.swagger true
toml set --toml-path $HOME/.saw/node1/config/app.toml rosetta.denom-to-suggest "usae"

# validator3
toml set --toml-path $HOME/.saw/node2/config/app.toml api.address "tcp://localhost:1315"
toml set --toml-path $HOME/.saw/node2/config/app.toml grpc.address "0.0.0.0:9086"
toml set --toml-path $HOME/.saw/node2/config/app.toml grpc-address "0.0.0.0:9087"
toml set --toml-path $HOME/.saw/node2/config/app.toml minimum-gas-prices "0usae"
toml set --toml-path $HOME/.saw/node2/config/app.toml api.swagger true
toml set --toml-path $HOME/.saw/node2/config/app.toml rosetta.denom-to-suggest "usae"

# change config.toml values

# validator1
toml set --toml-path $HOME/.saw/node0/config/config.toml p2p.allow_duplicate_ip true
# validator2
toml set --toml-path $HOME/.saw/node1/config/config.toml proxy_app tcp://127.0.0.1:26655
toml set --toml-path $HOME/.saw/node1/config/config.toml p2p.allow_duplicate_ip true
toml set --toml-path $HOME/.saw/node1/config/config.toml rpc.laddr tcp://127.0.0.1:26757
toml set --toml-path $HOME/.saw/node1/config/config.toml p2p.laddr tcp://0.0.0.0:26756
toml set --toml-path $HOME/.saw/node1/config/config.toml p2p.persistent_peers "$PEER@0.0.0.0:26656"


# validator3
toml set --toml-path $HOME/.saw/node2/config/config.toml proxy_app tcp://127.0.0.1:26652
toml set --toml-path $HOME/.saw/node2/config/config.toml p2p.allow_duplicate_ip true
toml set --toml-path $HOME/.saw/node2/config/config.toml rpc.laddr tcp://127.0.0.1:26857
toml set --toml-path $HOME/.saw/node2/config/config.toml p2p.laddr tcp://0.0.0.0:26856
toml set --toml-path $HOME/.saw/node2/config/config.toml p2p.persistent_peers "$PEER@0.0.0.0:26656"



# copy validator1 genesis file to validator2-3
cp $HOME/.saw/node0/config/genesis.json $HOME/.saw/node1/config/genesis.json
cp $HOME/.saw/node0/config/genesis.json $HOME/.saw/node2/config/genesis.json

echo "start all three validators"
screen -S $VALIDATOR1 -d -m sawd start --home=$HOME/.saw/node0
screen -S $VALIDATOR2 -d -m sawd start --home=$HOME/.saw/node1
screen -S $VALIDATOR3 -d -m sawd start --home=$HOME/.saw/node2

echo "Wait 10s for validator 1 start"
sleep 10

sawd tx staking create-validator \
    --amount=10000000usae \
    --keyring-backend $KEYRING \
    --chain-id $CHAINID \
    --home $HOME/.saw/node1\
    --commission-rate="0.10" \
    --commission-max-rate="0.20" \
    --commission-max-change-rate="0.01" \
    --min-self-delegation="1000000" \
    --gas=2000000 \
    --gas-prices="0.0025usae" \
    --from=$VALIDATOR2\
    --moniker=$VALIDATOR2\
    --pubkey=$(sawd tendermint show-validator --home=$HOME/.saw/node1)\
    -y

sawd tx staking create-validator \
    --amount=10000000usae \
    --keyring-backend $KEYRING \
    --chain-id $CHAINID \
    --home $HOME/.saw/node2\
    --commission-rate="0.10" \
    --commission-max-rate="0.20" \
    --commission-max-change-rate="0.01" \
    --min-self-delegation="1000000" \
    --gas=2000000 \
    --gas-prices="0.0025usae" \
    --from=$VALIDATOR3\
    --moniker=$VALIDATOR3\
    --pubkey=$(sawd tendermint show-validator --home=$HOME/.saw/node2)\
    -y

# echo $(sawd keys show $VALIDATOR1 -a --keyring-backend=test --home=$HOME/.saw/node0)
# echo $(sawd keys show $VALIDATOR1 -a --keyring-backend=test --home=$HOME/.saw/node1)
# echo $(sawd keys show $VALIDATOR1 -a --keyring-backend=test --home=$HOME/.saw/node2)
sleep 5
sawd q staking validators