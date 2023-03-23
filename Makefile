.PHONY: all run clean help

APP = bluebell

## win: 编译打包win
.PHONY: win
win:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/bluebell
build:
	go build -o ${APP}

## 编译win，linux，mac平台
.PHONY: all
all: tool win

run:
	@go run ./main.go conf/config.yaml

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: tool
gotool:
	go fmt ./  #格式化代码
	go vet ./  #检查

## test: Run unit test.
.PHONY: test
test:
	@$(MAKE) go.test

## 清理二进制文件
clean:
	@if [ -f ./bin/${APP}-win64.exe ] ; then rm ./bin/${APP}-win64.exe; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make win - 编译 Go 代码, 生成windows二进制文件"
	@echo "make tidy - 执行go mod tidy"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除编译的二进制文件"
	@echo "make all - 编译多平台的二进制文件"