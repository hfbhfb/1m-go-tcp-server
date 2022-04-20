#!/bin/bash



#ps -aux|grep runableapp|grep asserver|awk '{print $2}'
# 杀死进程
ps -aux|grep runableapp|grep asserver|awk '{print $2}'
ps -aux|grep runableapp|grep asserver|awk '{print $2}'|xargs kill



