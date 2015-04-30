#!/bin/bash

exec 1> /tmp/1
exec 2>&1

echo "1" >&1
echo "2" >&2

exec 1> /tmp/2
exec 2>&1
echo "3" >&1
echo "4" >&2

exec 1>> /tmp/1
exec 2>&1
echo "1" >&1
echo "2" >&2

exec 1> /dev/tty
exec 2> /dev/tty

echo "echo /tmp/1 ..."
cat /tmp/1

echo "echo /tmp/2 ..."
cat /tmp/2

echo "ls /not_existed"
ls /not_existed
