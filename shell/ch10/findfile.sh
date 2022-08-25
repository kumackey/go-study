#!/bin/sh

CMDNAME=$(basename $0)
if [ $# -eq 0 ]; then
  echo "Usage: $CMDNAME file [directory ... ]" 1>&2
  exit 1
fi

NAME=$1
shift

find "${@:-.}" -name "$NAME" -print
