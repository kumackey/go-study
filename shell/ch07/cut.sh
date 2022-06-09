#!/bin/sh

STRING="stringstringstring..."
FIRST=1
LEN=5
echo "$STRING" | cut -c$FIRST-$LEN
