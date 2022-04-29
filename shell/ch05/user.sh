#!/bin/sh

USER=$(id | sed 's/uid=.*(\(.*\)) gid=.*/\1/')
USERID=$(id | sed -e 's/uid=//' -e 's/(.*//')

echo user $USER
echo uid $USERID
