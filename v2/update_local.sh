#! /bin/bash

### copy/update package into GOROOT
### set you GOROOT path here.

AGNI_PATH=$GOROOT/src/agnione/v2

sudo rm -r $AGNI_PATH 

sudo mkdir -p $AGNI_PATH  ### create the folder for AgniOne packages

sudo cp -r ./src $AGNI_PATH  ### copy/update packages

echo "AgniOne packages are update to $AGNI_PATH" 