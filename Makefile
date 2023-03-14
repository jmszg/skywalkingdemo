ENVVAR=CGO_ENABLED=0 LD_FLAGS=-s
GOOS?=linux
TAG?=v1.1.0
REGISTRY?=registry.cn-hangzhou.aliyuncs.com/zgjhub
export GOPROXY=https://goproxy.cn,direct

build:
	rm -rf ./build/amd64
	$(ENVVAR) GOOS=$(GOOS) go build -o build/amd64/skywalkingdemo-loadgenerate ./cmd/loadgenerate
	$(ENVVAR) GOOS=$(GOOS) go build -o build/amd64/skywalkingdemo-server1 ./cmd/server1
	$(ENVVAR) GOOS=$(GOOS) go build -o build/amd64/skywalkingdemo-server2 ./cmd/server2
	$(ENVVAR) GOOS=$(GOOS) go build -o build/amd64/skywalkingdemo-server3 ./cmd/server3

make-image:
	docker build -t ${REGISTRY}/skywalkingdemo-loadgenerate:${TAG} -f ./cmd/loadgenerate/Dockerfile .
	docker build -t ${REGISTRY}/skywalkingdemo-server1:${TAG} -f ./cmd/server1/Dockerfile .
	docker build -t ${REGISTRY}/skywalkingdemo-server2:${TAG} -f ./cmd/server2/Dockerfile .
	docker build -t ${REGISTRY}/skywalkingdemo-server3:${TAG} -f ./cmd/server3/Dockerfile .

push-image:
	docker push  ${REGISTRY}/skywalkingdemo-loadgenerate:${TAG}
	docker push  ${REGISTRY}/skywalkingdemo-server1:${TAG}
	docker push  ${REGISTRY}/skywalkingdemo-server2:${TAG}
	docker push  ${REGISTRY}/skywalkingdemo-server3:${TAG}


