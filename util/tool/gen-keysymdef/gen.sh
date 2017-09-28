#!/bin/sh
gopath=`go env GOPATH`
./gen-keysymdef 2>log |gofmt > $gopath/src/github.com/linuxdeepin/go-x11-client/util/keysyms/auto.go
