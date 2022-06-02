#!/bin/sh

while [ $# -gt 0 ]
do
  case $1 in
  -f) FRAG=TRUE
    shift
    ;;
  -v) VALUE=$2
    shift 2
    ;;
  -v*) VALUE=`echo "$1" | sed 's/^..//'`
    shift
    ;;
  --) shift
    break
    ;;
  -*) echo "Usage: $0 [-f] [-v value]" 1>&2
    exit 1
    ;;
    *) break
      ;;
    esac
done

echo "$FRAG"
echo "$VALUE"