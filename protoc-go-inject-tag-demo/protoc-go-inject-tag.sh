#!/bin/bash

pb_go_files=`find . -name "*.pb.go"`
# shellcheck disable=SC2068
for file in ${pb_go_files[@]}
do
    if ! command -v protoc-go-inject-tag > /dev/null 2>&1; then
        `cd ~ && go get github.com/favadi/protoc-go-inject-tag`
    fi
    protoc-go-inject-tag -input=${file}
done