#!/bin/bash
DIRECTORY=$PWD
HANDLERS_PATH="/handlers"
BIN_PATH="/dist"
FULL_BIN_PATH="$DIRECTORY/$BIN_PATH"

function if_not_exist_create {
    REQ_PATH=$1
    if [ ! -d $REQ_PATH ]; then
        echo "Creating $REQ_PATH"
        mkdir $REQ_PATH
    fi
}

if [ "$1" != "" ]; then
    HANDLERS_PATH="$1"
fi

if_not_exist_create "$FULL_BIN_PATH"

export GO111MODULE=on
for CMD in `ls handlers`; do
    echo "Compiling $CMD"
    MAIN_PATH="$DIRECTORY/$HANDLERS_PATH/$CMD"
    DESTINY="$FULL_BIN_PATH/$CMD"
    if_not_exist_create $DESTINY
    env GOARCH=amd64 GOOS=linux go build -o $DESTINY $MAIN_PATH
    cd $DIRECTORY
done

echo "READY!"
exit 0
