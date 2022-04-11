#!/bin/sh

STRING=abc
case "$STRING" in
ABC) ehco "STRING is ABC" ;;
abc) echo "STRING is abc" ;;
xyz) echo "STRING is xyz" ;;
esac

case "$STRING" in
def | ghi) ;;         # どちらか
[Yy]*) ;;             # Yかyで始まる
[!Yy]*) echo hello ;; # Yでもyでも始まらない
\*) ;;                # *である
"*") ;;               # *である
esac
