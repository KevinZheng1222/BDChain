#! /bin/bash

# update the `tester` image by copying in the latest bdc binary

docker run --name builder tester true
docker cp $GOPATH/bin/bdc builder:/go/bin/bdc
docker commit builder tester
docker rm -vf builder

