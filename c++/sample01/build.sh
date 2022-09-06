#!/bin/bash

ROOT=$(cd `dirname $0`; pwd)

cd ${ROOT}/cal
g++ cal.cpp -c -o cal.o
ar rcs libcal.a cal.o

cd ${ROOT}/combine
g++ combine.cpp -fPIC -c -o combine.o
g++ -shared -o libcombine.so combine.o
 
cd ${ROOT}
g++ inc.cpp -c -o inc.o
g++ main.cpp -Ical -Icombine -c -o main.o
g++ main.o inc.o -Lcal -Lcombine -lcal -lcombine -o main

### export LD_LIBRARY_PATH=`pwd`/combine
### ./main 
