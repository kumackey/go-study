#!/bin/sh

a=1
while test $a -lt 3; do
  echo $a
  a=$(expr $a + 1)
done
