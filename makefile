
.PHONY: all test clean build install gen_api

GOFLAGS ?= $(GOFLAGS:)

all: install test

build:
	go build $(GOFLAGS) ./.

install:
	glide i

bench: install
	go test -run=NONE -bench=. $(GOFLAGS) ./...

clean:
	go clean $(GOFLAGS) -i ./...

rundev: build
	bee generate docs
	./natsmicro

# 测试
test:
	go test -v ./services/...


TARGET=$(src)
# 根据模型(model)定义, 生成接口(services)代码
# make src="models/group.go" gen_api
gen_api:
	python3 gen_api.py $(src)
