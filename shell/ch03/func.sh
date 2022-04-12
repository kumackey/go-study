#!/bin/sh

lsl() {
  ls -l $*
}

lsl ~/

posdisplay() {
  echo $0 $2 $3 $4
}

posdisplay aaa bbb ccc ddd
# $0は関数名とかではなく、./func.shとなる
