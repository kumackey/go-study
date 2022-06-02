#!/bin/sh

if [ "$1" = "-v" ]; then
  VERBOSE=TRUE
  shift
fi
for param in "$@"; do
  echo $param
done
