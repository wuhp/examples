#!/bin/bash

cd `dirname $0`

rm -f main
find . -name \*.o | xargs rm -f
find . -name \*.a | xargs rm -f
find . -name \*.so | xargs rm -f
