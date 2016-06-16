#!/bin/bash

echo "runing testcase $1 ..."

[ "$1" = "001" ] && exit 0
[ "$1" = "002" ] && exit 1

exit 0
