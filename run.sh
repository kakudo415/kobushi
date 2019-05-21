#!/bin/sh
git pull && make
env PORT=10300 bin/kobushi