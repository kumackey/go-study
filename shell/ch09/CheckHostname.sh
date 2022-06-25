#!/bin/sh

CheckHostname() {
  _PING=
  _HOST=${1:-$(hostname)}

  case $(uname -s) in
  FreeBSD | Linux | OSF1 | Darwin) _PING="ping -c1 $_HOST" ;;
  HP-UX) _PING="ping $_HOST 64 1" ;;
  *) return 1 ;;
  esac

  $_PING >/dev/null 2>&1
  return $?
}

#CheckHostname
#echo $?