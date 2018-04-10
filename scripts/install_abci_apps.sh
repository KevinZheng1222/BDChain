#! /bin/bash

# get the abci commit used by bdc
COMMIT=$(bash scripts/dep_utils/parse.sh abci)
echo "Checking out vendored commit for abci: $COMMIT"

go get -d github.com/bdc/abci
cd "$GOPATH/src/github.com/bdc/abci" || exit
git checkout "$COMMIT"
make get_tools
make get_vendor_deps
make install
