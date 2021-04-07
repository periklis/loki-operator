# Auto generated binary variables helper managed by https://github.com/bwplotka/bingo v0.4.0. DO NOT EDIT.
# All tools are designed to be build inside $GOBIN.
BINGO_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
GOPATH ?= $(shell go env GOPATH)
GOBIN  ?= $(firstword $(subst :, ,${GOPATH}))/bin
GO     ?= $(shell which go)

# Below generated variables ensure that every time a tool under each variable is invoked, the correct version
# will be used; reinstalling only if needed.
# For example for bingo variable:
#
# In your main Makefile (for non array binaries):
#
#include .bingo/Variables.mk # Assuming -dir was set to .bingo .
#
#command: $(BINGO)
#	@echo "Running bingo"
#	@$(BINGO) <flags/args..>
#
BINGO := $(GOBIN)/bingo-v0.4.0
$(BINGO): $(BINGO_DIR)/bingo.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/bingo-v0.4.0"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=bingo.mod -o=$(GOBIN)/bingo-v0.4.0 "github.com/bwplotka/bingo"

CFSSL := $(GOBIN)/cfssl-v1.5.0
$(CFSSL): $(BINGO_DIR)/cfssl.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/cfssl-v1.5.0"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=cfssl.mod -o=$(GOBIN)/cfssl-v1.5.0 "github.com/cloudflare/cfssl/cmd/cfssl"

CFSSLJSON := $(GOBIN)/cfssljson-v1.5.0
$(CFSSLJSON): $(BINGO_DIR)/cfssljson.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/cfssljson-v1.5.0"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=cfssljson.mod -o=$(GOBIN)/cfssljson-v1.5.0 "github.com/cloudflare/cfssl/cmd/cfssljson"

CONTROLLER_GEN := $(GOBIN)/controller-gen-v0.4.1
$(CONTROLLER_GEN): $(BINGO_DIR)/controller-gen.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/controller-gen-v0.4.1"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=controller-gen.mod -o=$(GOBIN)/controller-gen-v0.4.1 "sigs.k8s.io/controller-tools/cmd/controller-gen"

GOFUMPT := $(GOBIN)/gofumpt-v0.1.1
$(GOFUMPT): $(BINGO_DIR)/gofumpt.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/gofumpt-v0.1.1"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=gofumpt.mod -o=$(GOBIN)/gofumpt-v0.1.1 "mvdan.cc/gofumpt"

GOLANGCI_LINT := $(GOBIN)/golangci-lint-v1.38.0
$(GOLANGCI_LINT): $(BINGO_DIR)/golangci-lint.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/golangci-lint-v1.38.0"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=golangci-lint.mod -o=$(GOBIN)/golangci-lint-v1.38.0 "github.com/golangci/golangci-lint/cmd/golangci-lint"

KIND := $(GOBIN)/kind-v0.10.0
$(KIND): $(BINGO_DIR)/kind.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/kind-v0.10.0"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=kind.mod -o=$(GOBIN)/kind-v0.10.0 "sigs.k8s.io/kind"

KUSTOMIZE := $(GOBIN)/kustomize-v3.8.7
$(KUSTOMIZE): $(BINGO_DIR)/kustomize.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/kustomize-v3.8.7"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=kustomize.mod -o=$(GOBIN)/kustomize-v3.8.7 "sigs.k8s.io/kustomize/kustomize/v3"

OPERATOR_SDK := $(GOBIN)/operator-sdk-v1.5.0
$(OPERATOR_SDK): $(BINGO_DIR)/operator-sdk.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/operator-sdk-v1.5.0"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=operator-sdk.mod -o=$(GOBIN)/operator-sdk-v1.5.0 "github.com/operator-framework/operator-sdk/cmd/operator-sdk"
