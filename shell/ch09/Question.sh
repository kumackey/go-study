#!/bin/sh

Question() {
  if [ $# -lt 3 ]; then
    echo "Usage: Question question" \
      "default help message" 1>&2
    exit 1
  fi

  ANSWER=
  _DEFAULT=$2
  _QUESTION=
  _HELPMSG=$3

  if [ "$_DEFAULT" = "" ]; then
    _QUESTION="$1"
  else
    _QUESTION="$1 [$_DEFAULT]? "
  fi

  while true; do
    echo "$_QUESTION\c"
    read ANSWER
    case $(echo "$ANSWER" | tr [A-Z] [a-z]) in
    "") if [ "$_DEFAULT" != "" ]; then
      ANSWER=$_DEFAULT
      break
    fi ;;
    yes | y)
      ANSWER=yes
      break
      ;;
    no | n)
      ANSWER=no
      break
      ;;
    quit | q) exit 1 ;;
    +x | x) set $ANSWER ;;
    !*) eval $(expr "ANSWER" : "!\(.*\)") ;;
    "?")
      echo ""
      if [ "$_HELPMSG" = "" ]; then
        echo "No help available"
      else
        echo "$_HELPMSG"
      fi
      echo ""
      ;;
    *) break ;;
    esac
  done
}

Question "remove file" "yes" "the file is going to be removed"
echo $ANSWER