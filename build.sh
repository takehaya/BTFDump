#!/bin/bash

export GO111MODULE=on

VERSION=$(git rev-parse --abbrev-ref HEAD)
BUILD=$(git rev-parse --short HEAD)
LDFLAGS="-ldflags -X=github.com/takehaya/btfdump/pkg/version.Version=$VERSION.$BUILD"

go build ${LDFLAGS} -o btfdump.$VERSION.$BUILD ./cmd/btfdump
