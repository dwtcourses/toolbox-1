#!/bin/bash

set -e

output=bin/toolbox

mkdir -p bin
GO111MODULE=on go build -o $output github.com/owenrumney/toolbox/cmd/toolbox