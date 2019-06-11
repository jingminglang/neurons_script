#!/bin/bash


cat $1 |row2l |sed 's/"/\\"/g'
