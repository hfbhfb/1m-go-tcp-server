
SERVER1=gz1

App=1_simple_tcp_server
Version=0.0.1

GZProDir=/home/usera/projects/go/mod-pro/1m-go-tcp-server/${App}


all:gitcommitall fileprepare build workwithprocess
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


workwithprocess:killprocess
	-ssh $(SERVER1) "bash runprocess.sh "
	@echo "以进程方式打开"


killprocess:
	-ssh $(SERVER1) "cd ${GZProDir} && bash killprocess.sh "
	@echo "杀死进程"


work:imgprepare stopdocker startdocker 
	@echo "启动"

imgprepare:
	ssh $(SERVER1) "cd ${GZProDir} && docker build -t ${App}:${Version} . "
	

stopdocker:
	-ssh $(SERVER1) "docker stop contain1m${App}"
	-ssh $(SERVER1) "docker rm contain1m${App}"

startdocker:
	-ssh $(SERVER1) "docker run --name contain1m${App} --restart always -p 8972:8972 -d ${App}:${Version}"




