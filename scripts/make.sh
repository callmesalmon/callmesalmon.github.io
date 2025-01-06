#!/usr/bin/bash

# Needs to be run like this:
#     scripts/make.sh
for file in scripts/*; do
    case $file in
        *.go)
            go build $file
    esac
done
