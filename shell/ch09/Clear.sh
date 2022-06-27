#!/bin/sh

Clear() {
  { clear; } 2>/dev/null ||
    { tput clear; } 2>/dev/null ||
    for i in {1..100}; do
      echo
    done
}

Clear
