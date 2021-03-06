TF_PLUGIN_DIR=~/.terraform.d/plugins
PROVIDER_NAMESPACE=example.com/example/example

build: GOOS=$(shell go env GOOS)
build: GOARCH=$(shell go env GOARCH)
ifeq ($(OS),Windows_NT)  # is Windows_NT on XP, 2000, 7, Vista, 10...
build: DESTINATION=$(APPDATA)/terraform.d/plugins/$(PROVIDER_NAMESPACE)/0.0.1/$(GOOS)_$(GOARCH)
else
build: DESTINATION=$(HOME)/.terraform.d/plugins/$(PROVIDER_NAMESPACE)/0.0.1/$(GOOS)_$(GOARCH)
endif
build:
	@echo "==> Generating documentation"
	@go generate
	@echo "==> Generating example-server"
	cd example-server && swagger generate server -A pet -f ./swagger.yml
	@echo "==> Generating client"
	swagger generate client -f ./example-server/swagger.yml -A pet
	@echo "==> Installing plugin to $(DESTINATION)"
	@mkdir -p $(DESTINATION)
	go build -o $(DESTINATION)/terraform-provider-example_v0.0.1

run:
	@echo "==>starting server"
	PORT=9999 go run example-server/cmd/pet-server/main.go

testacc:
	TF_ACC=1 go test ./example -v $(TESTARGS) -timeout 1m

.PHONY: build
