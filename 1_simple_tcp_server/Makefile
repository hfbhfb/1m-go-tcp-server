
SERVER1=gz1

App=1mapp
Version=0.0.1

GZProDir=/home/usera/projects/go/mod-pro/1m-go-tcp-server

all:fileprepare build imgprepare stopdocker startdocker
	@echo "docker 化编译 , 可以借助docker分析工具检测 内存,网络,存储等信息"


fileprepare:
	ssh $(SERVER1) "cd ${GZProDir} && git pull origin master"


build:
	ssh $(SERVER1) "cd ${GZProDir}/1_simple_tcp_server && export CGO_ENABLED=0 && go build -o runableapp "

imgprepare:
	ssh $(SERVER1) "cd ${GZProDir}/1_simple_tcp_server && docker build -t ${App}:${Version} . "
	

stopdocker:
	-docker stop appcontain1m

startdocker:
	-docker run --name appcontain1m --restart always -p 8972:8972 -d ${Version}:${Version}
