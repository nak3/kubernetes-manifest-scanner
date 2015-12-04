#!/bin/bash -e

APP_NAME="kubernetes-manifest-scanner"
ORG_PATH="github.com/nak3"
REPO_PATH="${ORG_PATH}/${APP_NAME}"

GOBIN=${PWD}/bin
GOPATH=${PWD}/Godeps/_workspace

eval $(go env)

if [ ${GOOS} = "linux" ]; then
        echo "Building ${APP_NAME}..."
       # ${GOBIN}/godep go build -o ${GOBIN}/${APP_NAME} ${REPO_PATH}/pkg
       go build -o ${GOBIN}/${APP_NAME} ./pkg/main.go
else
        echo "Not on Linux - skipping ${APP_NAME} build"
fi
