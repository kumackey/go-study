#!/bin/sh

if [ -n "$(find "$1" -newer "$2" -print)" ]; then
  echo "Yes: $1 is newer than $2"
else
  echo "No: $1 is not newer than $2"
fi
