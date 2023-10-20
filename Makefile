REPO_URL ?= "quay.io/avillega/hello-pod"
VERSION ?= $(shell cat VERSION)

all: clean build

build:
	@go build -o hello-pod ./src

clean:
	@go clean -modcache

container-build: clean
	@podman build -t $(REPO_URL):$(VERSION) -f src/Dockerfile .

push: container-build
	@podman push $(REPO_URL):$(VERSION)
