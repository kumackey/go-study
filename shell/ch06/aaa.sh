#!/bin/sh

AFLAG=FALSE
BFLAG=FALSE
CFLAG=FALSE
VALUE=
OPT=

echo $#

while getopts abc: OPT; do
  case $OPT in
  a)
    AFLAG=TRUE
    ;;
  b)
    BFLAG=TRUE
    ;;
  c)
    CFLAG=TRUE
    VALUE=$OPTARG
    ;;
  \?)
    echo "Usage : $0 [-ab] [-c value] parameter" 1>&2
    exit 1
    ;;
  esac
done
shift $(expr $OPTIND - 1)

if [ "$AFLAG" = "TRUE" ]; then
  echo "option a is specified"
fi

if [ "$BFLAG" = "TRUE" ]; then
  echo "option b is specified"
fi

if [ "$CFLAG" = "TRUE" ]; then
  echo "option c is specified, and VALUE is $VALUE"
fi

echo "parameter is $1"
