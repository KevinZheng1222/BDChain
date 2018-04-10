#! /bin/bash
set -eu

DOCKER_IMAGE=$1
NETWORK_NAME=$2
ID=$3
APP_PROXY=$4

set +u
NODE_FLAGS=$5
set -u

set +eu

echo "starting bdc peer ID=$ID"
# start bdc container on the network
# NOTE: $NODE_FLAGS should be unescaped (no quotes). otherwise it will be
# treated as one flag.
set -u
docker run -d \
	--net="$NETWORK_NAME" \
	--ip=$(test/p2p/ip.sh "$ID") \
	--name "local_testnet_$ID" \
	--entrypoint bdc \
	-e TMHOME="/go/src/github.com/bdc/bdc/test/p2p/data/mach$ID/core" \
	--log-driver=syslog \
	--log-opt syslog-address=udp://127.0.0.1:5514 \
	--log-opt syslog-facility=daemon \
	--log-opt tag="{{.Name}}" \
		"$DOCKER_IMAGE" node $NODE_FLAGS --log_level=debug --proxy_app="$APP_PROXY"
