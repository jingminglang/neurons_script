#!/bin/bash

cd lib/parser
goyacc -o parser.go parser.y
cd ../../
mkdir -p dist
go build  -o dist
