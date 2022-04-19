#!/bin/bash



#ps -aux|grep runableapp1m|grep asserver|awk '{print $2}'
# 启动进程
cd /home/usera/projects/go/mod-pro/1m-go-tcp-server/1_simple_tcp_server
nohup ./runableapp1m  &
