#!/bin/sh

echo "$1" | grep "$2" >/dev/null
if [ $? -eq 0 ]; then
  echo "Include it"
else
  echo "Not inclued it."
fi
