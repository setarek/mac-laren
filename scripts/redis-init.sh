#!/bin/bash
while :
do
    redis-cli -h redis -p 6379 quit
    if [ $? -eq 0 ]; then
        cat /script/order_processor.lua | redis-cli -x FUNCTION LOAD REPLACE
        break
    else
        echo "server not ready, wait then retry..."
        sleep 3
    fi
done