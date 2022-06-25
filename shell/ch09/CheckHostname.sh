#!/bin/sh

CheckHostname() {
  _PING=
  _HOST=${1:-$(hostname)}

  case $(uname -s) in
  FreeBSD) _PING="ping -c1 $_HOST" ;;
  Linux) _PING="ping -c1 $_HOST" ;;
  OSF1) _PING="ping -c1 $_HOST" ;;
  HP-UX) _PING="ping $_HOST 64 1" ;;
  Darwin) _PING="ping -c1 $_HOST" ;;
  *) return 1 ;;
  esac

  if [ $($_PING 2>&1 | grep -ci "Unknown host") -eq 0 ]; then
    return 1
  else
    return 0
  fi
}
