#!/bin/sh

STRING=abc/def/ghi
PATTERN=/
STR=`echo "$STRING" | sed -e "s%$PATTERN.*%%"`
echo $STR