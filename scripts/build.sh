#!/bin/bash

set -e

if [[ ${GOOS} = "" ]]; then
    GOOS=$(uname | tr "[:upper:]" "[:lower:]")
fi

output=bin/f3-${GOOS}-amd64
if [[ "${CI}" = "" ]]; then
    output=toolbox
    version=""
else
    # default to $TRAVIS_TAG unless version is supplied
    if [[ $# -eq 0 ]]; then
	    version="${TRAVIS_TAG}"
    else
        version="$1"
    fi

    # last resort
    if [[ "$version" == "" ]]; then
      git fetch --tags
      version=`git describe --tags`
    fi

    if [[ "$version" == "" ]]; then
	    echo "Error: Cannot determine version"
	    exit 1
    fi
fi

mkdir -p bin
GO111MODULE=on go build -o $output -ldflags "-X github.com/owenrumney/toolbox/cmd/toolbox"