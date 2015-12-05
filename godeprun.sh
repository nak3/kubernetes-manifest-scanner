#!/bin/bash -e

APP_NAME=kubernetes-manifest-scanner
ORG_PATH="github.com/nak3"
REPO_PATH="${ORG_PATH}/kubernetes-manifest-scanner"

export GOBIN=${PWD}/bin
export GOPATH=${PWD}/gopath

if [ ! -h gopath/src/${REPO_PATH} ]; then
  mkdir -p gopath/src/${ORG_PATH}
  ln -s ../../../.. gopath/src/${REPO_PATH} || exit 255
fi

eval $(go env)

if [ ${GOOS} = "linux" ]; then
  echo "Getting dependencies..."
  go get github.com/tools/godep
  go get github.com/nak3/jvmap
  go get github.com/spf13/cobra
  go get k8s.io/kubernetes/pkg/kubectl/cmd/util
else
  echo "Not on Linux - skipping $APP_NAME build"
fi


if [ ! -e '.git' ]; then
  echo '.git directory not found'
  echo 'git init . && git add -A . && git commit -m "To test godep"'
  exit 1
fi

# echo "Making git repository..."
#
echo "Running godep save..."
${GOBIN}/godep save github.com/nak3/kubernetes-manifest-scanner/pkg
