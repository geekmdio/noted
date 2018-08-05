#!/usr/bin/env bash

PACKAGES=$(echo $(ls -d */) | sed -E -e 's/ehrproto\/|examples\///g' | sed -r s/\ /" github.com\/geekmdio\/noted\/"/g | sed 's/github.com\/geekmdio\/noted//' | sed 's/\///')

go test ${PACKAGES} -v -race -coverprofile=coverage.txt -covermode=atomic