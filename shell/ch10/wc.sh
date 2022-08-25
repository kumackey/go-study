#!/bin/sh

CMDNAME=$(basename $0)
USAGE="Usage: $CMDNAME [-c][-l][-w] [file | directory ...]"
TMP=/tmp/Wc.$$
LINE=
DIR=
PATTERN=
COUNT_LINES=FALSE
COUNT_WORDS=FALSE
COUNT_CHARS=FALSE
LINES=0
WORDS=0
CHARS=0

trap 'rm -f /tmp/*.$$; exit 1' 1 2 3 15

while :; do
  case $1 in
  -c)
    COUNT_CHARS=TRUE
    shift
    ;;
  -l)
    COUNT_LINES=TRUE
    shift
    ;;
  -w)
    COUNT_WORDS=TRUE
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

if [ $COUNT_LINES = FALSE -a $COUNT_WORDS = FALSE -a $COUNT_CHARS = FALSE ]; then
  COUNT_LINES=TRUE
  COUNT_WORDS=TRUE
  COUNT_CHARS=TRUE
fi

for parm in "${@:-.}"; do
  if [ -d "$parm" ]; then
    DIR="$parm"
    PATTERN="*"
  else
    DIR=$(dirname "$parm")
    PATTERN=$(basename "$parm")
  fi

  for d in $(find $DIR -type d -print | sort); do
    wc $d/$PATTERN 2>/dev/null | grep -v " total$" >$TMP
    exec <$TMP
    while read LINE; do
      set -- $LINE
      if [ "$COUNT_LINES" = "TRUE" ]; then
        LINES=$(expr $LINES + $1)
        echo "  $1"
      fi
      shift

      if [ "$COUNT_WORDS" = "TRUE" ]; then
        WORDS=$(expr $WORDS + $1)
        echo "  $1"
      fi
      shift

      if [ "$COUNT_CHARS" = "TRUE" ]; then
        CHARS=$(expr $CHARS + $1)
        echo "  $1"
      fi
      shift

      echo "  $@"
    done
  done
done

if [ "$COUNT_LINES" = "TRUE" ]; then
  echo "  $LINES"
fi

if [ "$COUNT_WORDS" = "TRUE" ]; then
  echo "  $WORDS"
fi
if [ "$COUNT_CHARS" = "TRUE" ]; then
  echo "  $CHARS"
fi

echo "  Total"
rm -f /tmp/*.$$
exit 0
