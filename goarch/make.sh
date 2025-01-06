#!/usr/bin/bash

cd src && \
go build main.go
mv main /usr/local/bin/goarch

cd ..
