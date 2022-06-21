GOC=go
LINUX64=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
#repository namesapce in docker hub.
NAMESPACE= docker.io/yxty/k8s
REPO=webhooker

VERSION=v0.1
#VERSION=0.6.02
#VERSION=0.4.01
#VERSION=0.3.41
#VERSION=0.3.32
#VERSION=0.1.27

IMAGE=$(NAMESPACE)/$(REPO)/$(VERSION)

all: push

build:
	$(LINUX64) $(GOC) build -o bin/webhooker

docker: build
	docker build -f Dockerfile -t $(IMAGE) .

run: docker
	docker run --rm -p 8080:8080 $(IMAGE)

push: docker
	docker push $(IMAGE)

pro: push
	kubectl set image deployment/webhooker webhooker=$(IMAGE)

dev: push
	kubectl set image deployment/webhooker-dev webhooker=$(IMAGE) 
clean: 
	rm bin/webhooker
	docker rmi -f $(IMAGE)

local: build
	bin/webhooker
