#!/bin/bash -e

APP_NAME="kubernetes-manifest-scanner"
ORG_PATH="github.com/nak3"
REPO_PATH="${ORG_PATH}/${APP_NAME}"

GOBIN=${PWD}/bin
GOPATH=${PWD}/Godeps/_workspace:${PWD}/_output/local/go/

if [ ! -h _output/local/go/src/${REPO_PATH} ]; then
  mkdir -p _output/local/go/src/${ORG_PATH}
  ln -s ../../../../../.. _output/local/go/src/${REPO_PATH}
fi

eval $(go env)

if [ ${GOOS} = "linux" ]; then
  echo "Building ${APP_NAME}..."
  go build -o ${GOBIN}/${APP_NAME} ${PWD}/_output/local/go/src/github.com/nak3/kubernetes-manifest-scanner/pkg/main.go
  echo "Building bash_comp..."
  go build -o ${GOBIN}/bash_comp_autogenerater ${PWD}/_output/local/go/src/github.com/nak3/kubernetes-manifest-scanner/pkg/gen_bash_comp.go
else
  echo "Not on Linux - skipping ${APP_NAME} build"
fi
