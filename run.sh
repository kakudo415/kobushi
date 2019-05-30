#!/bin/sh
cd `dirname $0`
git pull && make
env PORT=10300 bin/kobushi
