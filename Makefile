BIN="./bin"
SRC=$(shell find . -name "*.go")
XC_OS="linux darwin"
XC_ARCH="amd64"
XC_PARALLEL="2"

IMAGE := jharshman/bitlybot
TAG := latest

ifeq (, $(shell which gox))
$(warning "could not find gox in $(PATH), run: go get github.com/mitchellh/gox")
endif

ifeq (, $(shell which skaffold))
$(warning "could not find skaffold in $(PATH)")
endif

.PHONY: fmt vet build all publish clean registrylogin deploy

default: all

all: fmt vet build

fmt:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

build:
	$(info ******************** building ********************)
	gox \
		-os=$(XC_OS) \
		-arch=$(XC_ARCH) \
		-parallel=$(XC_PARALLEL) \
		-output=$(BIN)/{{.Dir}}_{{.OS}}_{{.Arch}} \
		;

publish: registrylogin
	$(info ******************** publishing ********************)
	skaffold build

registrylogin:
	echo $(DOCKER_PASSWORD) | docker login -u $(DOCKER_USER) --password-stdin

deploy: registrylogin
	$(info ******************** build & deploy ********************)
	skaffold run

vet:
	$(info ******************** vetting ********************)
	go vet ./...

clean:
	rm -rf $(BIN)
