#!/bin/bash
DIRECTORY=$PWD
HANDLERS_PATH="handlers"
BIN_PATH="bin"

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

if [ "$2" != "" ]; then
    BIN_PATH="$2"
fi

export GO111MODULE=on
for CMD in `ls $HANDLERS_PATH`; do
    echo "Compiling $CMD"
    MAIN_PATH="$DIRECTORY/$HANDLERS_PATH/$CMD"
    DESTINY="$DIRECTORY/$BIN_PATH"
    env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o $DESTINY $MAIN_PATH
    cd $DIRECTORY
done

echo "READY!"
exit 0
