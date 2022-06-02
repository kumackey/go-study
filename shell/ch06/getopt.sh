#!/bin/sh

FLAG=
VALUE=
OPT=
set -- `getopt fv:  $*`
if [ $? != 0 ]; then
  echo "Usage: $0 [-f] [-v value]" 1>&2
  exit 1
fi

for OPT in $*; do
  case $OPT in
  -f)
    FLAG=TRUE
    shift
    ;;
  -v)
    VALUE=$2
    shift 2
    ;;
  --)
    shift
    break
    ;;
  esac
done
