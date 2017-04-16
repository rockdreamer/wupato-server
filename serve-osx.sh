#!/bin/bash
echo "Rebuilding"
pkill wupato-server-main
if [[ `ps aux | grep ${PWD##*/}-main.osx | grep -v grep | wc -l` > 0 ]]
        then
        killall ${PWD##*/}-main.osx 2> /dev/null
fi

go build -o /tmp/${PWD##*/}-main.osx && /tmp/${PWD##*/}-main.osx &
