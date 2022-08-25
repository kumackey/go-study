#!/bin/sh

CMDNAME=$(basename $0)
USAGE="Usage: $CMDNAME [-iv] string [filename]"
STRING=
FILENAME=
I=
L=-l

if [ "$OPTIND" = 1 ]; then
  while getopts iv OPT; do
    case $OPT in
    i)
      I=-i
      ;;
    v)
      L=
      ;;
    \?)
      echo "$USAGE" 1>&2
      exit 1
      ;;
    esac
  done
  shift $(expr $OPTIND - 1)
else
  USAGE="Usage: $CMDNAME [-i] [-v] string [filename]"
  while :; do
    case $1 in
    -i)
      I=-i
      shift
      ;;
    -v)
      L=
      shift
      ;;
    --)
      shift
      break
      ;;
    -*)
      echo "$USAGE" 1>&2
      exit 1
      ;;
    *)
      break
      ;;
    esac
  done
fi

if [ $# -lt 1 -o $# -gt 2 ]; then
  echo "$USAGE" 1>&2
  exit 1
fi

STRING=$1
FILENAME=${2:-"*"}

find . \( -type f -o -type l \) -name "$FILENAME" -print |
  xargs grep $I $L -- "$STRING" /dev/null

exit 0