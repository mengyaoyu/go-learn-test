#!/bin/bash

PROJECT_HOME=$(cd `dirname $0`; pwd)
PROJECT_VENDOR=$PROJECT_HOME/vendor

GOPATH=$PROJECT_HOME:$PROJECT_VENDOR

export GOPATH

go build src/main.go
