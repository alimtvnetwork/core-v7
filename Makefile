BinariesDirectory = ./bin
WindowsBinariesDirectory = bin
MainDirectory = cmd/main
ConfigDirectory = ./configs
ConfigDirectoryForWindows = configs
GoVersion=v1.24
MyApp=cli
GOCONVEY_PORT ?= 8080

.PHONY: clean lint changelog snapshot release
.PHONY: build
.PHONY: deps

target: run-main
all: run-main

# Cross-platform executable check
UNAME_S := $(shell uname -s 2>/dev/null || echo Windows)
ifeq ($(UNAME_S),Windows)
    SHELL := cmd.exe
    RM = del /Q
    MKDIR = mkdir
    SEP = \\
else
    RM = rm -f
    MKDIR = mkdir -p
    SEP = /
endif

EXECUTABLES = git go
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec) 2>/dev/null || where $(exec) 2>nul),some string,$(error "No $(exec) in PATH")))

VERSION ?= $(shell git describe --tags `git rev-list --tags --max-count=1` 2>/dev/null || echo "v0.0.0")
BINARY = cli
MAIN = main.go

BUILDDIR = build
GITREV = $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILDTIME = $(shell date +'%FT%TZ%z' 2>/dev/null || echo "unknown")
GO_BUILDER_VERSION=latest

run-main:
	go run cmd/main/*.go

run-server:
	go run cmd/server/*.go

run-client:
	go run cmd/client/*.go

run-sample:
	go run cmd/sample/*.go

build:
	@$(MKDIR) $(BUILDDIR)
	go build -o $(BUILDDIR)$(SEP)$(BINARY) ./$(MainDirectory)/...
	@echo "Build $(BINARY) done."
	@echo "Run \"$(BUILDDIR)$(SEP)$(BINARY)\" to start $(BINARY)."

run:
	./$(BUILDDIR)/$(BINARY)

build-run: build run

run-tests:
	cd tests && go test -v ./...

run-all-tests:
	go test -v ./...

run-pkg-tests:
	@echo "Usage: make run-pkg-tests PKG=<package>"
	cd tests && go test -v ./integratedtests/$(PKG)/...

goconvey:
	@which goconvey > /dev/null 2>&1 || (echo "Installing GoConvey..." && go install github.com/smartystreets/goconvey@latest)
	@echo "Starting GoConvey on http://localhost:$(GOCONVEY_PORT)"
	cd tests && goconvey -port $(GOCONVEY_PORT)

vet:
	go vet ./...

fmt:
	gofmt -w -s .

tidy:
	go mod tidy

clean:
	$(RM) -r $(BUILDDIR)

run-linux-docker:
	cd scripts && chmod +x ./docker-run-linux.sh
	cd scripts && sh ./docker-run-linux.sh

run-direct:
	"$(BinariesDirectory)/main"

linux-run:
	cd "$(BinariesDirectory)" && ./main

cat-ssh:
	cat ~/.ssh/id_rsa.pub

ssh-sample:
	echo "ssh-keygen -t rsa -b 4096 -C 'Your email'"

modify-authorized-keys:
	sudo vim ~/.ssh/authorized_keys

git-clean-get:
	git reset --hard
	git clean -df
	git status
	git pull

run-docker-build-linux:
	./bin/cli-linux-amd64

run-docker-build-windows:
	./bin/cli-windows-amd64.exe

export-private-key:
	export PRIVATE_KEY=$$(cat ~/.ssh/id_rsa | base64)

export-go-version:
	export GO_BUILDER_VERSION=$(GoVersion)

export-current-binaries:
	echo 'export PATH=$$PATH:$$PWD/bin'

gopath-export:
	export GOPATH="/root/go"

all-exports: export-private-key export-go-version export-current-binaries gopath-export

changelog:
	git-chglog $(VERSION) > CHANGELOG.md

debug: all-exports
	# Example : https://t.ly/pJiQ, https://t.ly/JEjg, https://t.ly/5Rre, https://goreleaser.com/install/
	docker run --rm --privileged \
		-e PRIVATE_KEY=$(PRIVATE_KEY) \
		-v $$PWD:/usr/src/$(MyApp) \
		-v $$GOPATH:/go \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-w /usr/src/$(MyApp) \
		ghcr.io/gythialy/golang-cross:$(GO_BUILDER_VERSION) --snapshot --rm-dist

basic-snapshot:
	cd scripts && chmod +x ./docker-deploy.sh
	cd scripts && sh ./docker-deploy.sh

snapshot: all-exports
	# Example : https://t.ly/pJiQ, https://t.ly/JEjg, https://t.ly/5Rre
	docker run --rm -it --privileged \
		-e PRIVATE_KEY=$(PRIVATE_KEY) \
		-v $(CURDIR):/usr/src/myapp \
		-v $(GOPATH):/go \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-w /usr/src/myapp \
		ghcr.io/gythialy/golang-cross:$(GO_BUILDER_VERSION) bash

release: changelog
	docker run --rm -it --privileged \
		-e GITHUB_TOKEN=$(GITHUB_TOKEN) \
		-e PRIVATE_KEY=$(PRIVATE_KEY) \
		-v $(CURDIR):/golang-cross-example \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v $(GOPATH)/src:/go/src \
		-w /golang-cross-example \
		ghcr.io/gythialy/golang-cross:$(GO_BUILDER_VERSION) --rm-dist --release-notes=CHANGELOG.md
