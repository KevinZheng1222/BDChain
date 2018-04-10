#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

debora run -- bash -c "cd \$GOPATH/src/github.com/bdc/bdc; killall bdc"
debora run -- bash -c "cd \$GOPATH/src/github.com/bdc/bdc; bdc unsafe_reset_priv_validator; rm -rf ~/.bdc/data"
debora run -- bash -c "cd \$GOPATH/src/github.com/bdc/bdc; git pull origin develop; make"
debora run -- bash -c "cd \$GOPATH/src/github.com/bdc/bdc; mkdir -p ~/.bdc/logs"
debora run --bg --label bdc -- bash -c "cd \$GOPATH/src/github.com/bdc/bdc; bdc node 2>&1 | stdinwriter -outpath ~/.bdc/logs/bdc.log"
printf "\n\nSleeping for a minute\n"
sleep 60
debora download bdc "logs/async$1"
debora run -- bash -c "cd \$GOPATH/src/github.com/bdc/bdc; killall bdc"
