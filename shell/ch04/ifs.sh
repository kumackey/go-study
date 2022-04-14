#!/bin/sh

OLDIFS=$IFS
IFS=
while read LINES; do
  echo $LINES
done <data.txt
IFS=$OLDIFS
