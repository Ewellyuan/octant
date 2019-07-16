# Copyright (c) 2019 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

SHELL=/bin/bash
BUILD_TIME=$(shell date -u +%Y-%m-%dT%T%z)
GIT_COMMIT=$(shell git rev-parse --short HEAD)

LD_FLAGS= '-X "main.buildTime=$(BUILD_TIME)" -X main.gitCommit=$(GIT_COMMIT)'
GO_FLAGS= -ldflags=$(LD_FLAGS)
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install

VERSION ?= v0.3.0

.PHONY: version
version:
	@echo $(VERSION)

# Run all tests
.PHONY: test
test: generate
	@echo "-> $@"
	@env go test -v ./internal/... ./pkg/...

# Run govet
.PHONY: vet
vet:
	@echo "-> $@"
	@env go vet  ./internal/... ./pkg/...

octant-dev:
	@mkdir -p ./build
	@env $(GOBUILD) -o build/octant $(GO_FLAGS) -v ./cmd/octant

octant-docker:
	@env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o /octant $(GO_FLAGS) -v ./cmd/octant

generate:
	@echo "-> $@"
	@go generate -v ./pkg/... ./internal/...

go-install:
	@env GO111MODULE=on $(GOINSTALL) github.com/GeertJohan/go.rice
	@env GO111MODULE=on $(GOINSTALL) github.com/GeertJohan/go.rice/rice
	@env GO111MODULE=on $(GOINSTALL) github.com/asticode/go-astilectron-bundler
	@env GO111MODULE=on $(GOINSTALL) github.com/golang/mock/gomock
	@env GO111MODULE=on $(GOINSTALL) github.com/golang/mock/mockgen
	@env GO111MODULE=on $(GOINSTALL) github.com/golang/protobuf/protoc-gen-go

# Remove all generated go files
.PHONY: clean
clean:
	@rm -rf ./internal/octant/fake
	@rm -rf ./internal/kubeconfig/fake
	@rm -rf ./internal/link/fake
	@rm -rf ./internal/event/fake
	@rm -rf ./internal/config/fake
	@rm -rf ./internal/api/fake
	@rm -rf ./internal/portforward/fake
	@rm -rf ./internal/objectstore/fake
	@rm -rf ./internal/queryer/fake
	@rm -rf ./internal/cluster/fake
	@rm -rf ./internal/module/fake
	@rm -rf ./internal/modules/overview/printer/fake
	@rm -rf ./pkg/store/fake
	@rm -rf ./pkg/plugin/fake
	@rm -rf ./pkg/plugin/api/fake
	@rm ./pkg/icon/rice-box.go

web-deps:
	@cd web; npm ci

web-build: web-deps
	@cd web; npm run build
	@go generate ./web

web-test: web-deps
	@cd web; npm run test:headless

ui-server:
	OCTANT_DISABLE_OPEN_BROWSER=false OCTANT_LISTENER_ADDR=localhost:3001 $(GOCMD) run ./cmd/octant/main.go $(OCTANT_FLAGS)

ui-client:
	@cd web; API_BASE=http://localhost:3001 npm run start

gen-electron:
	@GOCACHE=${HOME}/cache/go-build astilectron-bundler -v -c configs/electron/bundler.json

.PHONY: changelogs
changelogs:
	hacks/changelogs.sh

.PHONY: release
release:
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push --follow-tags

.PHONY: ci
ci: test vet web-test web-build octant-dev

.PHONY: ci-quick
ci-quick:
	@cd web; npm run build
	@go generate ./web
	make octant-dev

install-test-plugin:
	mkdir -p ~/.config/vmdash/plugins
	go build -o ~/.config/vmdash/plugins/pluginstub github.com/vmware/octant/cmd/pluginstub

.PHONY:
build-deps:
