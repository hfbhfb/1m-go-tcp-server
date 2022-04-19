
SERVER1=gz1

App=1_simple_client
Version=0.0.1

GZProDir=/home/usera/projects/go/mod-pro/1m-go-tcp-server/${App}

all:gitcommitall fileprepare build  workwithprocess
	@echo "docker 化编译 , 可以借助docker分析工具检测 内存,网络,存储等信息"

gitcommitall:
	-git commit -a -m "auto提交 ${App}" && git push gz master && git push origin master 
	@echo "提交git到服务"


fileprepare:
	ssh $(SERVER1) "cd ${GZProDir} && git pull origin master"
	@echo "赣州 git 获取提交"


build:
	ssh $(SERVER1) "cd ${GZProDir} && export CGO_ENABLED=0 && go build -o runableapp1m "
	@echo "编译"



workwithprocess:
	ssh $(SERVER1) "cd ${GZProDir} && nohup ./runableapp1m asclient"
	ssh $(SERVER1) "cd ${GZProDir} && nohup ./runableapp1m asclient"
	@echo "以进程方式打开"


work:imgprepare stopdocker startdocker stopdocker2 startdocker2
	@echo "启动"

imgprepare:
	ssh $(SERVER1) "cd ${GZProDir} && docker build -t ${App}:${Version} . "

stopdocker:
	-ssh $(SERVER1) "docker stop contain1m${App}"
	-ssh $(SERVER1) "docker rm contain1m${App}"
startdocker:
	-ssh $(SERVER1) "docker run --name contain1m${App} --restart always -d ${App}:${Version}"

stopdocker2:
	-ssh $(SERVER1) "docker stop contain1m2${App}"
	-ssh $(SERVER1) "docker rm contain1m2${App}"
startdocker2:
	-ssh $(SERVER1) "docker run --name contain1m2${App} --restart always -d ${App}:${Version}"





