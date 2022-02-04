#!/bin/bash

curl -s -o /tmp/names-generator-master.go https://raw.githubusercontent.com/docker/docker-ce/master/components/engine/pkg/namesgenerator/names-generator.go

diff -q /tmp/names-generator-master.go names-generator.go.orig

if [ $? != 0 ]; then
    echo "You must update the name generator."
else
    echo "You have the latest version of the name generator!"
fi

rm /tmp/names-generator-master.go
