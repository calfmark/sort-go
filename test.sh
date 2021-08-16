#!/bin/sh

cd `dirname ${0}` || exit 1
ROOTDIR=`pwd`
echo "ROOTDIR=$ROOTDIR"

find . -maxdepth 1 -type d -print | grep -v '^.$' | grep -v '.git' | grep -v 'test' | while read LINE
do
    echo -e "\n************************$LINE*******************************"
    cd $ROOTDIR && cd $LINE && go test -v
    cd $ROOTDIR && cd $LINE && go test -v -bench=. -benchmem
done