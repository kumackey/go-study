#!/bin/sh

for FILE in *; do
  SIZE=$(wc -c <$FILE)
  echo $SIZE $FILE | awk '{printf("%5s Bytes %s\n", $1, $2)}'
done
