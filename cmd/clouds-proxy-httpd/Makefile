MAIN = *.go
APP = $(shell basename `pwd`)
VERSION	= 1.0.0
DEPLOY_HOST=
TIME = $(shell date "+%F %T")
GIT = $(shell git rev-parse HEAD)
PKG = github.com/luopengift/version

FLAG = "-X '${PKG}.VERSION=${VERSION}' -X '${PKG}.APP=${APP}' -X '${PKG}.TIME=${TIME}' -X '${PKG}.GIT=${GIT}'"
build: 
	go build -ldflags $(FLAG) -o ${APP} ${MAIN}
deploy: build
	for HOST in $(DEPLOY_HOST); \
	do \
		scp -i ~/.ssh/id_rsa -P 36000 ${APP} ops@$$HOST:/tmp; \
	done
run: build
	pkill $(APP)
	nohup ./$(APP) >nohup.out 2>&1 &
debug:
	go run *.go -conf conf.yml
