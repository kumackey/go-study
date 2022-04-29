#!/bin/sh

FLAG=FALSE
VALUE=
OPT=
while getopts fv: OPT; do
  case $OPT in
  f)
    FLAG=TRUE
    ;;
  v)
    VALUE=$OPTARG
    ;;
  \?)
    echo "Usage : $0 [-f] [-v value]" 1>&2
    exit 1
    ;;
  esac
done
shift $(expr $OPTIND - 1)

echo FLAG is "$FLAG", OPT is "$VALUE"
