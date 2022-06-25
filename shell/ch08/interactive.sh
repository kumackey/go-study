#!/bin/sh

/bin/echo -n "Would you like to ... [y/n]?"
read ANSWER
case $ANSWER in
y | yes) FRAG=TRUE ;;
*) FRAG=FALSE ;;
esac

echo $FRAG
