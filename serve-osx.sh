#!/bin/bash
echo "Rebuilding"

if [[ `ps aux | grep ${PWD##*/}-main.osx | grep -v grep | wc -l` > 0 ]]
        then
        killall ${PWD##*/}-main.osx 2> /dev/null
fi

go build -o /tmp/${PWD##*/}-main.osx main.go && /tmp/${PWD##*/}-main.osx &