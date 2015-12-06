#!/bin/bash -e

APP_NAME="kubernetes-manifest-scanner"
ORG_PATH="github.com/nak3"
REPO_PATH="${ORG_PATH}/${APP_NAME}"

GOPATH=${PWD}/Godeps/_workspace:${PWD}/_output/local/go/

if [ ! -h _output/local/go/src/${REPO_PATH} ]; then
  mkdir -p _output/local/go/src/${ORG_PATH}
  ln -s ../../../../../.. _output/local/go/src/${REPO_PATH}
fi

echo "Building ${APP_NAME}..."
go build -o ${PWD}/bin/${APP_NAME} ${PWD}/_output/local/go/src/github.com/nak3/kubernetes-manifest-scanner/pkg/main.go

echo "Building bash_comp_autogenerator..."
go build -o ${PWD}/bin/bash_comp_autogenerator ${PWD}/_output/local/go/src/github.com/nak3/kubernetes-manifest-scanner/pkg/gen_bash_comp.go

echo "Build completed!

Next steps: set up your path

export PATH="${PWD}/bin/:$PATH"
source bash-comp/kuberenetes-manifest-scanner
"
