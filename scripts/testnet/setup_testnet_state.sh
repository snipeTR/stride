#!/bin/bash

set -eu

DEPLOYMENT_NAME="$1" # e.g. testnet
NETWORK_NAME="$2"    # e.g. stride
NUM_NODES="$3"    

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

STATE=$SCRIPT_DIR/state
VAL_PREFIX=val
PORT_ID=26656
STRIDE_CMD=build/strided
VAL_TOKENS=500000000ustrd
STAKE_TOKENS=300000000ustrd
DOMAIN=stridenet.co

echo "Cleaning state"
rm -rf $STATE
mkdir $STATE
touch $STATE/keys.txt

PEER_NODE_IDS=""
MAIN_ID=1 # Node responsible for genesis and persistent_peers
MAIN_NODE_NAME=""
MAIN_NODE_CMD=""
echo 'Initializing state for each node in the chain...'
for (( i=1; i <= $NUM_NODES; i++ )); do
    # Node names will be of the form: "stride-node1"
    node_name="${NETWORK_NAME}-node${i}"
    # Moniker is of the form: STRIDE_1
    moniker=$(printf "${NETWORK_NAME}_${i}" | awk '{ print toupper($0) }')

    # Create a state directory for the current node and initialize the chain
    mkdir -p $STATE/$node_name
    st_cmd="$STRIDE_CMD --home ${STATE}/$node_name"
    $st_cmd init $moniker --chain-id $NETWORK_NAME --overwrite 2> /dev/null

    # Update node networking configuration 
    sed -i -E "s|cors_allowed_origins = \[\]|cors_allowed_origins = [\"\*\"]|g" "${STATE}/${node_name}/config/config.toml"
    sed -i -E "s|127.0.0.1|0.0.0.0|g" "${STATE}/${node_name}/config/config.toml"
    # update the denom in the genesis file 
    sed -i -E 's|"stake"|"ustrd"|g' "${STATE}/${node_name}/config/genesis.json"

    # Get the endpoint and node ID
    endpoint="${node_name}.${DEPLOYMENT_NAME}.${DOMAIN}"
    node_id=$($st_cmd tendermint show-node-id)@$endpoint:$PORT_ID
    echo "Node ID: $node_id"

    # add a validator account
    val_acct="${VAL_PREFIX}${i}"
    $st_cmd keys add $val_acct --keyring-backend=test >> $STATE/keys.txt 2>&1
    val_addr=$($st_cmd keys show $val_acct --keyring-backend test -a)
    # Add this account to the current node
    $st_cmd add-genesis-account ${val_addr} $VAL_TOKENS
    # actually set this account as a validator on the current node 
    $st_cmd gentx $val_acct $STAKE_TOKENS --chain-id $NETWORK_NAME --keyring-backend test 2> /dev/null
    
    if [ $i -eq $MAIN_ID ]; then
        MAIN_NODE_NAME=$node_name
        MAIN_NODE_CMD=$st_cmd
        MAIN_NODE_ID=$node_id
    else
        # add this node's id to the list of peer nodes that will be used by the main node
        PEER_NODE_IDS="${node_id},${PEER_NODE_IDS}" 
        # also add this account and it's genesis tx to the main node
        $MAIN_NODE_CMD add-genesis-account ${val_addr} $VAL_TOKENS
        cp ${STATE}/${node_name}/config/gentx/*.json ${STATE}/${MAIN_NODE_NAME}/config/gentx/
    fi
done

# now we process gentx txs on the main node
$MAIN_NODE_CMD collect-gentxs 2> /dev/null

# wipe out the persistent peers for the main node (these are incorrectly autogenerated for each validator during collect-gentxs)
sed -i -E "s|persistent_peers = .*|persistent_peers = \"\"|g" "${STATE}/${MAIN_NODE_NAME}/config/config.toml"

# for all peer nodes....
for (( i=2; i <= $NUM_NODES; i++ )); do
    node_name="${NETWORK_NAME}-node${i}"
    # add the main node as a persistent peer
    sed -i -E "s|persistent_peers = .*|persistent_peers = \"${MAIN_NODE_ID}\"|g" "${STATE}/${node_name}/config/config.toml"
    # copy the main node's genesis to the peer nodes to ensure they all have the same genesis
    cp ${STATE}/${MAIN_NODE_NAME}/config/genesis.json ${STATE}/${node_name}/config/genesis.json
done