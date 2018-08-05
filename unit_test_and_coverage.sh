#!/usr/bin/env bash

SED_PREPROCESSED_PACKAGES=$(echo $(ls -d */) | sed -E -e 's/ehrproto\/|examples\///g' | sed -r s/\ /" github.com\/geekmdio\/noted\/"/g | sed 's/github.com\/geekmdio\/noted//' | sed 's/\///')

go test ${SED_PREPROCESSED_PACKAGES} -v -race -coverprofile=coverage.txt -covermode=atomic