#!/bin/sh

decimal2hex() {
  echo 16o $1 p | dc
}

for i in $*; do
  echo "$i is $(decimal2hex $i)."
done
