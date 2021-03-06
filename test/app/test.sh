#! /bin/bash
set -ex

#- kvstore over socket, curl
#- counter over socket, curl
#- counter over grpc, curl
#- counter over grpc, grpc

# TODO: install everything

export PATH="$GOBIN:$PATH"
export TMHOME=$HOME/.bdc_app

function kvstore_over_socket(){
    rm -rf $TMHOME
    bdc init
    echo "Starting kvstore_over_socket"
    abci-cli kvstore > /dev/null &
    pid_kvstore=$!
    bdc node > bdc.log &
    pid_bdc=$!
    sleep 5

    echo "running test"
    bash kvstore_test.sh "KVStore over Socket"

    kill -9 $pid_kvstore $pid_bdc
}

# start bdc first
function kvstore_over_socket_reorder(){
    rm -rf $TMHOME
    bdc init
    echo "Starting kvstore_over_socket_reorder (ie. start bdc first)"
    bdc node > bdc.log &
    pid_bdc=$!
    sleep 2
    abci-cli kvstore > /dev/null &
    pid_kvstore=$!
    sleep 5

    echo "running test"
    bash kvstore_test.sh "KVStore over Socket"

    kill -9 $pid_kvstore $pid_bdc
}


function counter_over_socket() {
    rm -rf $TMHOME
    bdc init
    echo "Starting counter_over_socket"
    abci-cli counter --serial > /dev/null &
    pid_counter=$!
    bdc node > bdc.log &
    pid_bdc=$!
    sleep 5

    echo "running test"
    bash counter_test.sh "Counter over Socket"

    kill -9 $pid_counter $pid_bdc
}

function counter_over_grpc() {
    rm -rf $TMHOME
    bdc init
    echo "Starting counter_over_grpc"
    abci-cli counter --serial --abci grpc > /dev/null &
    pid_counter=$!
    bdc node --abci grpc > bdc.log &
    pid_bdc=$!
    sleep 5

    echo "running test"
    bash counter_test.sh "Counter over GRPC"

    kill -9 $pid_counter $pid_bdc
}

function counter_over_grpc_grpc() {
    rm -rf $TMHOME
    bdc init
    echo "Starting counter_over_grpc_grpc (ie. with grpc broadcast_tx)"
    abci-cli counter --serial --abci grpc > /dev/null &
    pid_counter=$!
    sleep 1
    GRPC_PORT=36656
    bdc node --abci grpc --rpc.grpc_laddr tcp://localhost:$GRPC_PORT > bdc.log &
    pid_bdc=$!
    sleep 5

    echo "running test"
    GRPC_BROADCAST_TX=true bash counter_test.sh "Counter over GRPC via GRPC BroadcastTx"

    kill -9 $pid_counter $pid_bdc
}

cd $GOPATH/src/github.com/bdc/bdc/test/app

case "$1" in 
    "kvstore_over_socket")
    kvstore_over_socket
    ;;
"kvstore_over_socket_reorder")
    kvstore_over_socket_reorder
    ;;
    "counter_over_socket")
    counter_over_socket
    ;;
"counter_over_grpc")
    counter_over_grpc
    ;;
    "counter_over_grpc_grpc")
    counter_over_grpc_grpc
    ;;
*)
    echo "Running all"
    kvstore_over_socket
    echo ""
    kvstore_over_socket_reorder
    echo ""
    counter_over_socket
    echo ""
    counter_over_grpc
    echo ""
    counter_over_grpc_grpc
esac

