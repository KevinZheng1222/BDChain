#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

debora run -- bash -c "cd \$GOPATH/src/github.com/bdc/bdc; killall bdc; killall logjack"
debora run -- bash -c "cd \$GOPATH/src/github.com/bdc/bdc; bdc unsafe_reset_priv_validator; rm -rf ~/.bdc/data; rm ~/.bdc/config/genesis.json; rm ~/.bdc/logs/*"
debora run -- bash -c "cd \$GOPATH/src/github.com/bdc/bdc; git pull origin develop; make"
debora run -- bash -c "cd \$GOPATH/src/github.com/bdc/bdc; mkdir -p ~/.bdc/logs"
debora run --bg --label bdc -- bash -c "cd \$GOPATH/src/github.com/bdc/bdc; bdc node 2>&1 | stdinwriter -outpath ~/.bdc/logs/bdc.log"
debora run --bg --label logjack    -- bash -c "cd \$GOPATH/src/github.com/bdc/bdc; logjack -chopSize='10M' -limitSize='1G' ~/.bdc/logs/bdc.log"
printf "Done\n"
