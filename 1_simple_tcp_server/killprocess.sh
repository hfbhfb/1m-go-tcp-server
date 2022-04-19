#!/bin/bash



#ps -aux|grep runableapp1m|grep asserver|awk '{print $2}'
# 杀死进程
ps -aux|grep runableapp1m|grep asserver|awk '{print $2}'|xargs kill



