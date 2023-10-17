TARGETS           ?= linux/amd64 darwin/amd64 windows/amd64 linux/386 linux/ppc64le linux/s390x linux/arm linux/arm64
PROJECT_NAME	  := golang-graphql-example
PKG				  := github.com/oxyno-zeta/$(PROJECT_NAME)

# go option
GO        ?= go
# Uncomment to enable vendor
GO_VENDOR := # -mod=vendor
TAGS      :=
TESTS     := .
TESTFLAGS :=
LDFLAGS   := -w -s
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin
DISTDIR   := dist

# Required for globs to work correctly
SHELL=/usr/bin/env bash

#  Version

GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
DATE	   = $(shell date +%F_%T%Z)

BINARY_VERSION = ${GIT_SHA}
LDFLAGS += -X ${PKG}/pkg/${PROJECT_NAME}/version.Version=${BINARY_VERSION}
LDFLAGS += -X ${PKG}/pkg/${PROJECT_NAME}/version.GitCommit=${GIT_COMMIT}
LDFLAGS += -X ${PKG}/pkg/${PROJECT_NAME}/version.BuildDate=${DATE}

HAS_GORELEASER := $(shell command -v goreleaser;)
HAS_GIT := $(shell command -v git;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_CURL:=$(shell command -v curl;)
HAS_MOCKGEN:=$(shell command -v mockgen;)
HAS_GQLGEN:=$(shell command -v gqlgen;)
HAS_FIELDALIGNMENT:=$(shell command -v fieldalignment;)
HAS_GOTESTSUM:=$(shell command -v gotestsum;)
HAS_GOCOVER_COBERTURA:=$(shell command -v gocover-cobertura;)
# Uncomment to use gox instead of goreleaser
# HAS_GOX := $(shell command -v gox;)

.DEFAULT_GOAL := code/lint

#############
#   Build   #
#############

.PHONY: code/lint
code/lint: setup/dep/install
	golangci-lint run ./...

.PHONY: code/lint/gitlab
code/lint/gitlab: setup/dep/install
	golangci-lint run ./... --out-format code-climate:gl-code-quality-report.json,line-number

.PHONY: code/fieldalignment
code/fieldalignment: setup/dep/install
	fieldalignment -fix -test=false ./...

.PHONY: code/graphql/concat
code/graphql/concat:
	cat graphql/*.graphql > tools/graphql-inspector/validation.graphql
	@echo "you must commit this file for graphql non breaking change check"

.PHONY: code/graphql/no-break-check
code/graphql/no-break-check:  code/graphql/concat
	graphql-inspector diff 'git:origin/master:./tools/graphql-inspector/validation.graphql' './tools/graphql-inspector/validation.graphql'

.PHONY: code/graphql/generate
code/graphql/generate:
	gqlgen generate

.PHONY: code/generate
code/generate:
	$(GO) $(GO_VENDOR) generate ./...

.PHONY: code/graphql
code/graphql: code/graphql/generate code/graphql/concat

.PHONY: code/build
code/build: code/clean setup/dep/install
	$(GO) build $(GO_VENDOR) -o $(BINDIR)/$(PROJECT_NAME) $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' $(PKG)/cmd/${PROJECT_NAME}

# Uncomment to use gox instead of goreleaser
# .PHONY: code/build-cross
# code/build-cross: code/clean setup/dep/install
# 	CGO_ENABLED=0 GOFLAGS="-trimpath $(GO_VENDOR)" gox -output="$(DISTDIR)/bin/{{.OS}}-{{.Arch}}/{{.Dir}}" -osarch='$(TARGETS)' $(if $(TAGS),-tags '$(TAGS)',) -ldflags '$(LDFLAGS)' ${PKG}/cmd/${PROJECT_NAME}

.PHONY: code/build-cross
code/build-cross: code/clean setup/dep/install
ifdef HAS_GORELEASER
	goreleaser --snapshot --skip-publish
endif
ifndef HAS_GORELEASER
	curl -sL https://git.io/goreleaser | bash -s -- --snapshot --skip-publish
endif

.PHONY: code/clean
code/clean:
	@rm -rf $(BINDIR) $(DISTDIR)

#############
#  Release  #
#############

# Uncomment to use gox instead of goreleaser
# .PHONY: release/all
# release/all: code/clean setup/dep/install code/build-cross
# 	cp Dockerfile $(DISTDIR)/bin/linux-amd64

.PHONY: release/all
release/all: code/clean setup/dep/install
ifdef HAS_GORELEASER
	goreleaser
endif
ifndef HAS_GORELEASER
	curl -sL https://git.io/goreleaser | bash
endif

#############
#   Tests   #
#############

.PHONY: test/all
test/all: setup/dep/install setup/test/integration
	gotestsum --format testname --format-hide-empty-pkg  --junitfile junit.xml -- -p 1 $(GO_VENDOR) --tags=unit,integration -v -coverpkg=./pkg/... -covermode=count -coverprofile=c.out.tmp ./pkg/...

.PHONY: test/all/original
test/all/original: setup/dep/install setup/test/integration
	$(GO) test -p 1 $(GO_VENDOR) --tags=unit,integration -v -coverpkg=./pkg/... -covermode=count -coverprofile=c.out.tmp ./pkg/...

.PHONY: test/unit
test/unit: setup/dep/install
	$(GO) test $(GO_VENDOR) --tags=unit -v -coverpkg=./pkg/... -covermode=count -coverprofile=c.out.tmp ./pkg/...

.PHONY: test/integration
test/integration: setup/dep/install setup/test/integration
	$(GO) test -p 1 $(GO_VENDOR) --tags=integration -v -coverpkg=./pkg/... -covermode=count -coverprofile=c.out.tmp ./pkg/...

.PHONY: test/coverage
test/coverage:
	cat c.out.tmp | grep -v "mock_" | grep -v "generated" > c.out
	$(GO) tool cover -html=c.out -o coverage.html
	$(GO) tool cover -func c.out
	gocover-cobertura < c.out > coverage.xml

#############
#   Setup   #
#############

.PHONY: down/services
down/services:
	@echo "Down services"
	docker rm -f postgres || true
	docker rm -f opa || true
	docker rm -f keycloak || true
	docker rm -f rabbitmq || true

.PHONY: down/metrics-services
down/metrics-services:
	@echo "Down metrics services"
	docker rm -f prometheus || true
	docker rm -f grafana || true
	# docker rm -f jaeger || true
	docker rm -f tempo || true

.PHONY: down/dev-services
down/dev-services:
	@echo "Down dev services"
	docker rm -f pgadmin || true
	docker rm -f maildev || true

.PHONY: setup/dev-services
setup/dev-services: down/dev-services
	@echo "Setup dev services"
	docker run --rm --name pgadmin -p 8090:80 --link postgres:postgres -e 'PGADMIN_DEFAULT_EMAIL=user@domain.com' -e 'PGADMIN_DEFAULT_PASSWORD=SuperSecret' -d dpage/pgadmin4
	docker run --rm --name maildev -p 1080:1080 -p 1025:1025 -d maildev/maildev:1.1.0 --incoming-user fake --incoming-pass fakepassword

.PHONY: setup/metrics-services
setup/metrics-services: down/metrics-services
	@echo "Setup metrics services"
	docker run --rm -d --name prometheus -v $(CURDIR)/.local-resources/prometheus/prometheus.yml:/prometheus/prometheus.yml --network=host prom/prometheus:v2.18.0 --web.listen-address=:9191
	docker run --rm -d --name grafana -e GF_AUTH_ANONYMOUS_ENABLED=true -e GF_AUTH_ANONYMOUS_ORG_ROLE=Admin -e GF_AUTH_DISABLE_LOGIN_FORM=true -e GF_FEATURE_TOGGLES_ENABLE=traceqlEditor -v $(CURDIR)/.local-resources/grafana/datasources.yaml:/etc/grafana/provisioning/datasources/datasources.yaml --network=host grafana/grafana:10.0.0
	docker run --rm -d --name tempo  -v $(CURDIR)/.local-resources/tempo/tempo.yaml:/etc/tempo.yaml -p 14268:14268 -p 3200:3200 -p 4318:4318 -p 4317:4317 grafana/tempo:latest -config.file=/etc/tempo.yaml

.PHONY: setup/services
setup/services: down/services
	@echo "Setup services"
	docker run -d --rm --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e PGDATA=/var/lib/postgresql/data/pgdata -v $(CURDIR)/.run/postgres:/var/lib/postgresql/data postgres:14
	tar czvf .local-resources/opa/bundle.tar.gz --directory=.local-resources/opa/bundle example/
	docker run -d --rm --name opa -p 8181:8181 -v $(CURDIR)/.local-resources/opa/bundle.tar.gz:/bundle.tar.gz openpolicyagent/opa:0.25.2 run --server --log-level debug --log-format text --bundle /bundle.tar.gz
	docker run -d --rm --name keycloak -p 8088:8080 -e KEYCLOAK_IMPORT=/tmp/realm-export.json -v $(CURDIR)/.local-resources/keycloak/realm-export.json:/tmp/realm-export.json -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin quay.io/keycloak/keycloak:11.0.3
	# User: guest / Password: guest
	docker run --rm -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management

.PHONY: setup/test/integration
setup/test/integration:
	docker run --rm --link postgres:postgres --name postgres-integration-setup -e PGHOST=postgres -e PGUSER=postgres -e PGPASSWORD=postgres -e PGDATABASE=postgres postgres:12 psql -c 'CREATE DATABASE "postgres-integration"' || true

.PHONY: setup/dep/install
setup/dep/install:
ifndef HAS_GOLANGCI_LINT
	@echo "=> Installing golangci-lint tool"
ifndef HAS_CURL
	$(error You must install curl)
endif
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.54.2
endif
ifndef HAS_GIT
	$(error You must install Git)
endif
ifndef HAS_GQLGEN
	@echo "=> Installing gqlgen tool"
	$(GO) install github.com/99designs/gqlgen@v0.17.39
endif
ifndef HAS_MOCKGEN
	@echo "=> Installing mockgen tool"
	$(GO) install go.uber.org/mock/mockgen@v0.3.0
endif
ifndef HAS_GOTESTSUM
	@echo "=> Installing gotestsum tool"
	$(GO) install gotest.tools/gotestsum@v1.11.0
endif
ifndef HAS_GOCOVER_COBERTURA
	@echo "=> Installing gocover-cobertura tool"
	$(GO) install github.com/boumenot/gocover-cobertura@v1.2.0
endif
ifndef HAS_FIELDALIGNMENT
	@echo "=> Installing fieldalignment tool"
	$(GO) install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@v0.14.0
endif
# Uncomment to use gox instead of goreleaser
# ifndef HAS_GOX
# 	@echo "=> Installing gox"
# 	$(GO) install github.com/mitchellh/gox@v1.0.1
# endif
	$(GO) mod download all

.PHONY: setup/dep/tidy
setup/dep/tidy:
	$(GO) mod tidy

.PHONY: setup/dep/update
setup/dep/update:
	$(GO) get ./...

.PHONY: setup/dep/vendor
setup/dep/vendor:
	$(GO) mod vendor
