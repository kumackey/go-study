#!/bin/sh

pwd
( cd ../; pwd; VALUE="value"; echo $VALUE )

pwd
echo $VALUE

{ cd ../; pwd; VALUE="value2"; echo $VALUE; }

pwd
echo $VALUE