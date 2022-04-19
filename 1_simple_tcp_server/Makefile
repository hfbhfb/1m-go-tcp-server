
SERVER1=gz1

App=1mapp1
Version=0.0.1

GZProDir=/home/usera/projects/go/mod-pro/1m-go-tcp-server

all:gitcommitall fileprepare build imgprepare stopdocker startdocker
	@echo "docker 化编译 , 可以借助docker分析工具检测 内存,网络,存储等信息"

gitcommitall:
	-git commit -a -m "auto提交 ${App}" && git push gz master && git push origin master 
	@echo "提交git到服务"


fileprepare:
	ssh $(SERVER1) "cd ${GZProDir} && git pull origin master"
	@echo "赣州 git 获取提交"


build:
	ssh $(SERVER1) "cd ${GZProDir}/1_simple_tcp_server && export CGO_ENABLED=0 && go build -o runableapp "
	@echo "编译"


imgprepare:
	ssh $(SERVER1) "cd ${GZProDir}/1_simple_tcp_server && docker build -t ${App}:${Version} . "
	

stopdocker:
	-ssh $(SERVER1) "docker stop contain1m${App}"
	-ssh $(SERVER1) "docker rm contain1m${App}"

startdocker:
	-ssh $(SERVER1) "docker run --name contain1m${App} --restart always -p 8972:8972 -d ${App}:${Version}"




