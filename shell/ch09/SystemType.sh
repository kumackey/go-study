#!/bin/sh

SystemType() {
  _HOSTNAME=$(hostname | sed 's/\..*//')
  case $(uname -s) in
  Darwin) echo Darwin ;;
  esac
}
