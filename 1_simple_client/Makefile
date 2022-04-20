
SERVER1=gz1

App=1_simple_client
Version=0.0.1

GZProDir=/home/usera/projects/go/mod-pro/1m-go-tcp-server/${App}

all:gitcommitall fileprepare build  workwithdocker
	@echo "docker 化编译 , 可以借助docker分析工具检测 内存,网络,存储等信息"

gitcommitall:
	-git commit -a -m "auto提交 ${App}" && git push gz master && git push origin master 
	@echo "提交git到服务"


fileprepare:
	ssh $(SERVER1) "cd ${GZProDir} && git pull origin master"
	@echo "赣州 git 获取提交"


build:
	ssh $(SERVER1) "cd ${GZProDir} && export CGO_ENABLED=0 && go build -o runableapp "
	@echo "编译"


workwithdocker:imgprepare stopdocker startdocker stopdocker2 startdocker2 stopdocker3 startdocker3 stopdocker4 startdocker4 stopdocker5 startdocker5 stopdocker6 startdocker6 stopdocker7 startdocker7 stopdocker8 startdocker8 stopdocker9 startdocker9 stopdocker10 startdocker10 stopdocker11 startdocker11 stopdocker12 startdocker12 stopdocker13 startdocker13
	@echo "启动"
# stopdocker14 startdocker14
# stopdocker15 startdocker15
# stopdocker16 startdocker16
# stopdocker17 startdocker17

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


stopdocker3:
	-ssh $(SERVER1) "docker stop contain1m3${App}"
	-ssh $(SERVER1) "docker rm contain1m3${App}"
startdocker3:
	-ssh $(SERVER1) "docker run --name contain1m3${App} --restart always -d ${App}:${Version}"



stopdocker4:
	-ssh $(SERVER1) "docker stop contain1m4${App}"
	-ssh $(SERVER1) "docker rm contain1m4${App}"
startdocker4:
	-ssh $(SERVER1) "docker run --name contain1m4${App} --restart always -d ${App}:${Version}"


stopdocker5:
	-ssh $(SERVER1) "docker stop contain1m5${App}"
	-ssh $(SERVER1) "docker rm contain1m5${App}"
startdocker5:
	-ssh $(SERVER1) "docker run --name contain1m5${App} --restart always -d ${App}:${Version}"


stopdocker6:
	-ssh $(SERVER1) "docker stop contain1m6${App}"
	-ssh $(SERVER1) "docker rm contain1m6${App}"
startdocker6:
	-ssh $(SERVER1) "docker run --name contain1m6${App} --restart always -d ${App}:${Version}"


stopdocker7:
	-ssh $(SERVER1) "docker stop contain1m7${App}"
	-ssh $(SERVER1) "docker rm contain1m7${App}"
startdocker7:
	-ssh $(SERVER1) "docker run --name contain1m7${App} --restart always -d ${App}:${Version}"


stopdocker8:
	-ssh $(SERVER1) "docker stop contain1m8${App}"
	-ssh $(SERVER1) "docker rm contain1m8${App}"
startdocker8:
	-ssh $(SERVER1) "docker run --name contain1m8${App} --restart always -d ${App}:${Version}"


stopdocker9:
	-ssh $(SERVER1) "docker stop contain1m9${App}"
	-ssh $(SERVER1) "docker rm contain1m9${App}"
startdocker9:
	-ssh $(SERVER1) "docker run --name contain1m9${App} --restart always -d ${App}:${Version}"


stopdocker10:
	-ssh $(SERVER1) "docker stop contain1m10${App}"
	-ssh $(SERVER1) "docker rm contain1m10${App}"
startdocker10:
	-ssh $(SERVER1) "docker run --name contain1m10${App} --restart always -d ${App}:${Version}"


stopdocker11:
	-ssh $(SERVER1) "docker stop contain1m11${App}"
	-ssh $(SERVER1) "docker rm contain1m11${App}"
startdocker11:
	-ssh $(SERVER1) "docker run --name contain1m11${App} --restart always -d ${App}:${Version}"





stopdocker12:
	-ssh $(SERVER1) "docker stop contain1m12${App}"
	-ssh $(SERVER1) "docker rm contain1m12${App}"
startdocker12:
	-ssh $(SERVER1) "docker run --name contain1m12${App} --restart always -d ${App}:${Version}"

stopdocker13:
	-ssh $(SERVER1) "docker stop contain1m13${App}"
	-ssh $(SERVER1) "docker rm contain1m13${App}"
startdocker13:
	-ssh $(SERVER1) "docker run --name contain1m13${App} --restart always -d ${App}:${Version}"

stopdocker14:
	-ssh $(SERVER1) "docker stop contain1m14${App}"
	-ssh $(SERVER1) "docker rm contain1m14${App}"
startdocker14:
	-ssh $(SERVER1) "docker run --name contain1m14${App} --restart always -d ${App}:${Version}"

stopdocker15:
	-ssh $(SERVER1) "docker stop contain1m15${App}"
	-ssh $(SERVER1) "docker rm contain1m15${App}"
startdocker15:
	-ssh $(SERVER1) "docker run --name contain1m15${App} --restart always -d ${App}:${Version}"

stopdocker16:
	-ssh $(SERVER1) "docker stop contain1m16${App}"
	-ssh $(SERVER1) "docker rm contain1m16${App}"
startdocker16:
	-ssh $(SERVER1) "docker run --name contain1m16${App} --restart always -d ${App}:${Version}"

stopdocker17:
	-ssh $(SERVER1) "docker stop contain1m17${App}"
	-ssh $(SERVER1) "docker rm contain1m17${App}"
startdocker17:
	-ssh $(SERVER1) "docker run --name contain1m17${App} --restart always -d ${App}:${Version}"
