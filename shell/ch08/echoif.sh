#!/bin/sh

if [ "$VERBOSE" = "TRUE" ]; then
  ECHO=echo
else
  ECHO=:
fi

$ECHO "Message"