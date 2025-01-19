#! /bin/bash

### copy/update package into GOROOT
### set you GOROOT path here.
#AGNI_PATH=/usr/local/go/src/agnione/v1

AGNI_PATH=$GOROOT/src/agnione/v1

rm -r $AGNI_PATH 

mkdir -p $AGNI_PATH  ### create the folder for AgniOne packages

cp -r ./src $AGNI_PATH  ### copy/update packages