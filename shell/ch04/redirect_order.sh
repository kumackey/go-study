#!/bin/sh

echo first
cat not_exist_file >/dev/null 2>&1

echo second
cat not_exist_file 2>&1 >/dev/null
