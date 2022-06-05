#!/bin/bash

for c in `./bin/tsuru | grep "Available commands" -A 500 | cut -f3 -d' ' | sort -u`
do
    cat ./docs/source/reference.rst | grep "$c" >/dev/null 2>&1
    RESULT=$?
    if [ $RESULT -eq 1 ]
    then
        echo "${c} is not documented"
    fi
done
