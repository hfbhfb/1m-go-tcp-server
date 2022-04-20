#!/bin/bash



#ps -aux|grep runableapp|grep asserver|awk '{print $2}'
# 启动进程
nohup ./runableapp asserver &
