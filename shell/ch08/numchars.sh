#!/bin/sh

case "$1" in
?) echo 1 ;;
*) echo $(expr "$1" : '.*') ;;
esac
