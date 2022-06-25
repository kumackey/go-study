#!/bin/sh

IFS="=/"
for f in $(pwd); do
  echo $f
done
