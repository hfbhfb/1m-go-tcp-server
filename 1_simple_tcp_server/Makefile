

App=1mapp
Version=0.0.1

all:
	export CGO_ENABLED=0 && go build -o runableapp
	docker build -t ${Version}:${Version}
	@echo "docker 化编译 , 可以借助docker分析工具检测 内存,网络,存储等信息"

build:
	export CGO_ENABLED=0 && go build -o runableapp

imgprepare:
	docker build -t 1with1m:v0.0.1 .


stopdocker:
	-docker stop 1mappcontain

startdocker:
	-docker run --name 1mappcontain --restart always -p 8972:8972 -d ${Version}:${Version}
