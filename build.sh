#!/bin/bash

cd lib/parser
goyacc  -o parser.go parser.y    
cd ../../
go build main.go
