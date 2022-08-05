#!/bin/sh

Prompt() {
  if [ "$(echo -n)" = "-n" ]; then
    echo "${@-> }\c"
  else
    echo -n "${@-> }"
  fi
}

Prompt "hello"
