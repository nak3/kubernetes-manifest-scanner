#!/bin/bash -e

APP_NAME="kubernetes-manifest-scanner"
ORG_PATH="github.com/nak3"
REPO_PATH="${ORG_PATH}/${APP_NAME}"

if [ ! -h gopath/src/${REPO_PATH} ]; then
        mkdir -p gopath/src/${ORG_PATH}
        ln -s ../../../.. gopath/src/${REPO_PATH} || exit 255
fi

export GOBIN=${PWD}/bin
export GOPATH=${PWD}/gopath

eval $(go env)

if [ ${GOOS} = "linux" ]; then
        echo "Building ${APP_NAME}..."
	# TODO
        #${GOBIN}/godep go build -o ${GOBIN}/${APP_NAME} ${REPO_PATH}/pkg
        go build -o ${GOBIN}/${APP_NAME} ${REPO_PATH}/pkg
else
        echo "Not on Linux - skipping ${APP_NAME} build"
fi
