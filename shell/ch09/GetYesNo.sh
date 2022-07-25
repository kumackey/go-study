#!/bin/sh

GetYesNo() {
  _ANSWER=
  if [ $# -eq 0 ]; then
    echo "Usage: Get Yes No massage" 1>&2
    exit 1
  fi
  echo "$@"
  while :; do
    read _ANSWER
    case "$_ANSWER" in
    [yY] | yes | YES | Yes) return 0 ;;
    [nN] | no | NO | No) return 1 ;;
    *) echo "Please enter y or n." ;;
    esac
  done
}

GetYesNo "$@"
