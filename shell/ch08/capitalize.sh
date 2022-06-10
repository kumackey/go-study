#!/bin/sh

STRING="$1"
CHAR=$(expr "$STRING" : '\(.\).*')
REMINDER=$(expr "$STRING" : '.\(.*\)')
CHAR=$(echo "$CHAR" | tr '[a-z]' '[A-Z]')
echo $CHAR$REMINDER
