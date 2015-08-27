#!/bin/bash

ps -ef | grep nginx | grep -v grep| awk '{print $2}' | xargs kill -9
echo "Start nginx server ... "
nginx -c /Users/lihaoquan/GoProjects/Playground/src/beelike/conf/beelike_nginx.conf
echo "Start nginx is running... "

if [ ! -f beelike ]; then
   echo "Please input 'bee run' "
   exit 0
fi

#beelike










