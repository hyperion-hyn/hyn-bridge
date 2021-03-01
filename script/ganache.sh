#!/bin/bash

bootstrap_ganache() {
    # Init testing directory
    PROJECT_PATH=$1
    TESTING_DIR="$PROJECT_PATH/test_data"
    if [ -d $TESTING_DIR ] 
    then
        echo "Remove legacy testing data"
        rm -r $TESTING_DIR
    fi
    mkdir $TESTING_DIR

    # Run ganache
    echo "Starting gananche"
    ganache-cli -a 90 -e 10000000000000000 -p 7545 -h 127.0.0.1 -g 10000000000 -l 10000000 -i 5777 \
        --acctKeys=$TESTING_DIR/accounnt_private.text --db $TESTING_DIR/db \
        2>&1 > $TESTING_DIR/ganache-output-$(date +%Y-%M-%dT%H:%M:%S).txt &
}

shut_down_ganache() {
    PROJECT_PATH=$1
    TESTING_DIR="$PROJECT_PATH/test_data"

    # Kill ganache process
    PID=`ps -eaf | grep ganache-cli | grep -v grep | awk '{print $2}'`
    if [[ "" !=  "$PID" ]]; then
        echo "killing gananche $PID"
        kill -9 $PID
    fi

    # Clean test data
    rm -r $TESTING_DIR
}