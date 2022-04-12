#!/bin/sh

VAR1=value
VAR2=VAR1

echo $"$VAR2"
echo $$VAR2

eval echo \$$VAR2

eval "ls"
