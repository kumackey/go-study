#!/bin/sh

DownShift() {
  echo "$@" | tr '[A-Z]' '[a-z]'
}

DownShift AJSJS KKK iuiuUI
