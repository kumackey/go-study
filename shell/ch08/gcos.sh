#!/bin/sh

OLDIFS=$IFS
IFS=:
while read USER PASS UID GID GCOS REMINDER; do
  echo "$USER" "$GCOS"
done </etc/passwd
IFS=$OLDIFS
