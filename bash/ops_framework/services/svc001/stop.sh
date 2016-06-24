#!/bin/bash

docker-compose -f `dirname $0`/src/service.yml stop
