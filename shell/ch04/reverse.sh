#!/bin/sh

ed - data.txt <<- !
g/^/m0
w
q
!
