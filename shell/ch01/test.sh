#!/bin/sh

if test -f abc; then
  echo "exists!"
else
  echo "not exists!"
fi

if [ -f adc ]; then
  echo "exists!"
fi
