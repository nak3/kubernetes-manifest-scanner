#!/bin/bash -e

APP_NAME=kubernetes-manifest-scanner
ORG_PATH="github.com/nak3"
REPO_PATH="${ORG_PATH}/kubernetes-manifest-scanner"

export GOPATH=${PWD}/gopath

if ! type godep > /dev/null 2>&1; then
  echo "You haven't installed godep... Please install first."
  echo "go get github.com/tools/godep"
fi

if [ ! -h gopath/src/${REPO_PATH} ]; then
  mkdir -p gopath/src/${ORG_PATH}
  ln -s ../../../.. gopath/src/${REPO_PATH} || exit 255
fi

echo "Getting dependencies into ${GOPATH}..."
go get github.com/nak3/jvmap
go get github.com/spf13/cobra
go get k8s.io/kubernetes/pkg/kubectl/cmd/util

if [ ! -e '.git' ]; then
  echo '.git directory not found'
  echo 'git init . && git add -A . && git commit -m "To test godep"'
  exit 1
fi

echo "Running godep save..."
godep save github.com/nak3/kubernetes-manifest-scanner/pkg
