#!/bin/bash

set -e

for image in `cat $(dirname $0)/src/service.yml | grep "image:" | sed "s/^.*image: *//g"`
do
    docker pull ${image}
done
