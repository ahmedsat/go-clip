#!/bin/env bash

outBath=build/debug
out=go-clip-client

clear

mkdir -p $outBath

go mod tidy && 

go build -v -tags debug -o $outBath/$out  ./cmd && 

./$outBath/$out $*
