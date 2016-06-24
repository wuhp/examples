#!/bin/bash

set -e

docker-compose -f `dirname $0`/src/service.yml up -d
