#!/bin/sh

stty -echo
echo "Enter your password"
read PASSWORD
stty echo

echo $PASSWORD
