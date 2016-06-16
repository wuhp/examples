#!/bin/bash

ROOT_DIR=`dirname $0`

usage() {
    echo "Usage:"
    echo "  run.sh [testsuit] [testcase]"
    echo "  run.sh info"
}

print_testcase() {
    cat $1 | sed "s/#.*$//g" | grep -v "^ *$" | sed "s/^/  /g"
}

print_test_info() {
    for suit in `ls ${ROOT_DIR}/testsuits`
    do
        echo "${suit}:"
        list_file=${ROOT_DIR}/testsuits/${suit}/testcase.list
        [ -f ${list_file} ] && print_testcase ${list_file}
    done
}

testsuit() {
    [ ! -d ${ROOT_DIR}/testsuits/${1} ] && echo "testsuit ${1} not found" && return 1

    ${ROOT_DIR}/testsuits/${1}/init.sh > /dev/null
    if [ $? -ne 0 ]; then
        echo "testsuit ${1} init error"
        ${ROOT_DIR}/testsuits/${1}/clear.sh > /dev/null 2>&1
        return 1
    fi

    result=0
    if [ "$2" != "" ]; then
        ${ROOT_DIR}/testsuits/${1}/run.sh $2 > /dev/null
        result=$?
        if [ $result -eq 0 ]; then
            echo "[   OK   ]  -->  $1 - $2"
        else
            echo "[ Failed ]  -->  $1 - $2"
        fi
    else
        list_file=${ROOT_DIR}/testsuits/${1}/testcase.list
        if [ -f ${list_file} ]; then
            for testcase in `print_testcase ${list_file}`
            do
                ${ROOT_DIR}/testsuits/${1}/run.sh ${testcase} > /dev/null
                ret=$?
                [ $ret -ne 0 ] && result=$ret
                if [ $ret -eq 0 ]; then
                    echo "[   OK   ]  -->  $1 - ${testcase}"
                else
                    echo "[ Failed ]  -->  $1 - ${testcase}"
                fi
            done
        fi
    fi
    
    ${ROOT_DIR}/testsuits/${1}/clear.sh > /dev/null 2>&1
    return $result
}

all() {
    result=0
    for suit in `ls ${ROOT_DIR}/testsuits`
    do
        testsuit ${suit}
        ret=$?
        [ $ret -ne 0 ] && result=$ret
    done
    return $result
}

#
# main
#

[ "$1" = "-h" ] || [ "$1" = "--help" ] && usage && exit 0
[ "$1" = "info" ] && print_test_info && exit 0

if [ "$1" != "" ]; then
    testsuit $1 $2
    exit $?
fi

all
exit $?
