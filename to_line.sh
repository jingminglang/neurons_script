#!/bin/bash


c=`cat $1`
echo $c |sed 's/"/\\"/g'
