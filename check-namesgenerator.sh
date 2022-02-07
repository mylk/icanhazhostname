#!/bin/bash

curl -s -o /tmp/names-generator.go https://raw.githubusercontent.com/docker/docker-ce/master/components/engine/pkg/namesgenerator/names-generator.go

# sha256sum -b /tmp/names-generator.go > checksum.txt
sha256sum -c checksum.txt &> /dev/null

if [ $? != 0 ]; then
    echo "You must update the name generator."
else
    echo "You have the latest version of the name generator!"
fi

rm /tmp/names-generator.go

