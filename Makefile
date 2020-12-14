TF_PLUGIN_DIR=~/.terraform.d/plugins
PROVIDER_NAMESPACE=example.com/example/example

default: build

build: GOOS=$(shell go env GOOS)
build: GOARCH=$(shell go env GOARCH)
ifeq ($(OS),Windows_NT)  # is Windows_NT on XP, 2000, 7, Vista, 10...
build: DESTINATION=$(APPDATA)/terraform.d/plugins/$(PROVIDER_NAMESPACE)/0.0.1/$(GOOS)_$(GOARCH)
else
build: DESTINATION=$(HOME)/.terraform.d/plugins/$(PROVIDER_NAMESPACE)/0.0.1/$(GOOS)_$(GOARCH)
endif
build:
	@echo "==> Installing plugin to $(DESTINATION)"
	@mkdir -p $(DESTINATION)
	go build -o $(DESTINATION)/terraform-provider-example_v0.0.1

.PHONY: build
