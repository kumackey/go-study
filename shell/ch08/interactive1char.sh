#!/bin/sh

/bin/echo -n "Would you like to ... [y/n]?"
stty raw
ANSWER=$(dd bs=1 count=1 2>/dev/null)
ssty -raw
echo ""
case "$ANSWER" in
[yY]) FRAG=TRUE ;;
*) FRAG=FALSE ;;
esac

echo $FRAG
