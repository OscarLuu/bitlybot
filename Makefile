BIN="./bin"
SRC=$(shell find . -name "*.go")
XC_OS="linux darwin"
XC_ARCH="amd64"
XC_PARALLEL="2"

ifeq (, $(shell which gox))
$(warning "could not find gox in $(PATH), run: go get github.com/mitchellh/gox")
endif

.PHONY: fmt vet build all

default: all

all: fmt vet build

fmt:
	$(info ******************** checking formatting ********************)
	gofmt -d $(SRC)

build:
	$(info ******************** building ********************)
	gox \
		-os=$(XC_OS) \
		-arch=$(XC_ARCH) \
		-parallel=$(XC_PARALLEL) \
		-output=$(BIN)/{{.Dir}}_{{.OS}}_{{.Arch}} \
		;

vet:
	$(info ******************** vetting ********************)
	go vet ./...

clean:
	rm -rf $(BIN)
