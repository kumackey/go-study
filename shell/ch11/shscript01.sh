#!/bin/sh

ABC=0

for i in 1 2 3 4; do
  ABC=$(expr $ABC + $i)
done

echo $ABC
