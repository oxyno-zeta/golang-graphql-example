TARGETS           ?= linux/amd64 darwin/amd64 windows/amd64 linux/386 linux/ppc64le linux/s390x linux/arm linux/arm64
BUILDX_PLATFORMS  ?= linux/arm64,linux/amd64,linux/s390x,linux/ppc64le
PROJECT_NAME	  := golang-graphql-example
PKG				  := github.com/oxyno-zeta/$(PROJECT_NAME)
IMG_NAME		  ?= $(PROJECT_NAME)
IMG_VERSION		  ?= latest
REGISTRY_IMG      ?= localhost

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
IMG       ?= $(PROJECT_NAME)

# Required for globs to work correctly
SHELL=/usr/bin/env bash

#  Version

GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
DATE	   = $(shell date +%F_%T%Z)

ifeq ($(GIT_TAG),)
BINARY_VERSION = ${GIT_SHA}
else
BINARY_VERSION = ${GIT_TAG}
LDFLAGS += -X ${PKG}/pkg/${PROJECT_NAME}/version.Metadata=
endif
LDFLAGS += -X ${PKG}/pkg/${PROJECT_NAME}/version.Version=${BINARY_VERSION}
LDFLAGS += -X ${PKG}/pkg/${PROJECT_NAME}/version.GitCommit=${GIT_COMMIT}
LDFLAGS += -X ${PKG}/pkg/${PROJECT_NAME}/version.BuildDate=${DATE}

HAS_GORELEASER := $(shell command -v goreleaser;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_CURL:=$(shell command -v curl;)
HAS_MOCKGEN:=$(shell command -v mockgen;)
HAS_GQLGEN:=$(shell command -v gqlgen;)
HAS_FIELDALIGNMENT:=$(shell command -v fieldalignment;)
HAS_GOTESTSUM:=$(shell command -v gotestsum;)
HAS_GOCOVER_COBERTURA:=$(shell command -v gocover-cobertura;)
# Uncomment to use gox instead of goreleaser
# HAS_GOX := $(shell command -v gox;)

#
## Tool versions
#

# ? Note: Go install versions are inline because renovate can manage them like that.

# renovate: datasource=github-tags depName=golangci/golangci-lint
GOLANGCI_LINT_VERSION := "v1.63.4"
# renovate: datasource=docker depName=dpage/pgadmin4
PGADMIN_VERSION := "9.0.0"
# renovate: datasource=docker depName=maildev/maildev
MAILDEV_VERSION := "2.2.1"
# renovate: datasource=docker depName=prom/prometheus
PROMETHEUS_VERSION := "v3.1.0"
# renovate: datasource=docker depName=grafana/grafana
GRAFANA_VERSION := "11.4.0"
# renovate: datasource=docker depName=grafana/tempo
TEMPO_VERSION := "2.7.0"
# renovate: datasource=docker depName=postgres
POSTGRESQL_VERSION := "16"
# renovate: datasource=docker depName=openpolicyagent/opa
OPA_VERSION := "1.0.0"
# renovate: datasource=docker depName=quay.io/keycloak/keycloak
KEYCLOAK_VERSION := "26.1.0"
# renovate: datasource=docker depName=rabbitmq
RABBITMQ_VERSION := "4.0.5"

DEFAULT_CONTAINER_RUNTIME ?= podman

.DEFAULT_GOAL := code/lint

#############
#   Build   #
#############

.PHONY: code/lint
code/lint: setup/dep/install setup/dep/lint/install
	golangci-lint run ./...

.PHONY: code/lint/gitlab
code/lint/gitlab: setup/dep/install setup/dep/lint/install
	golangci-lint run ./... --out-format code-climate:gl-code-quality-report.json,line-number

.PHONY: code/fieldalignment
code/fieldalignment: setup/dep/install setup/dep/gen/install
	fieldalignment -fix -test=false ./...

.PHONY: code/graphql/concat
code/graphql/concat:
	cat graphql/*.graphql > tools/graphql-inspector/validation.graphql
	@echo "you must commit this file for graphql non breaking change check"

.PHONY: code/graphql/no-break-check
code/graphql/no-break-check:  code/graphql/concat
	graphql-inspector diff 'git:origin/master:./tools/graphql-inspector/validation.graphql' './tools/graphql-inspector/validation.graphql'

.PHONY: code/graphql/generate
code/graphql/generate: setup/dep/gen/install
	gqlgen generate

.PHONY: code/generate
code/generate: setup/dep/gen/install
	$(GO) $(GO_VENDOR) generate ./...

.PHONY: code/graphql
code/graphql: code/graphql/generate code/graphql/concat

.PHONY: code/build
code/build: code/clean setup/dep/install
	CGO_ENABLED=0 $(GO) build $(GO_VENDOR) -o $(BINDIR)/$(PROJECT_NAME) $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' $(PKG)/cmd/${PROJECT_NAME}

# Uncomment to use gox instead of goreleaser
# .PHONY: code/build-cross
# code/build-cross: code/clean setup/dep/install setup/dep/build/install
# 	CGO_ENABLED=0 GOFLAGS="-trimpath $(GO_VENDOR)" gox -output="$(DISTDIR)/bin/{{.OS}}-{{.Arch}}/{{.Dir}}" -osarch='$(TARGETS)' $(if $(TAGS),-tags '$(TAGS)',) -ldflags '$(LDFLAGS)' ${PKG}/cmd/${PROJECT_NAME}

.PHONY: code/build-cross
code/build-cross: code/clean setup/dep/install
ifdef HAS_GORELEASER
	goreleaser -p 2 --snapshot
endif
ifndef HAS_GORELEASER
	curl -sL https://git.io/goreleaser | bash -s -- -p 2 --snapshot
endif

.PHONY: code/clean
code/clean:
	@rm -rf $(BINDIR) $(DISTDIR)

#############
#  Release  #
#############

.PHONY: release/docker/multiarch
release/docker/multiarch:
ifeq ("$(DEFAULT_CONTAINER_RUNTIME)", "docker")
	- $(DEFAULT_CONTAINER_RUNTIME) buildx create --name project-v3-builder
	$(DEFAULT_CONTAINER_RUNTIME) buildx use project-v3-builder
	$(DEFAULT_CONTAINER_RUNTIME) buildx build --platform=$(BUILDX_PLATFORMS) --push --tag $(REGISTRY_IMG)/$(IMG_NAME):$(IMG_VERSION) -f Dockerfile.multiarch .
	- $(DEFAULT_CONTAINER_RUNTIME) buildx rm project-v3-builder
else
	- $(DEFAULT_CONTAINER_RUNTIME) manifest create $(IMG_NAME):$(IMG_VERSION)
	$(DEFAULT_CONTAINER_RUNTIME) build --platform=$(BUILDX_PLATFORMS) --manifest $(IMG_NAME):$(IMG_VERSION) -f Dockerfile.multiarch .
	$(DEFAULT_CONTAINER_RUNTIME) manifest push $(IMG_NAME):$(IMG_VERSION) $(REGISTRY_IMG)/$(IMG_NAME):$(IMG_VERSION)
endif

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
test/all: setup/dep/install setup/dep/test/install setup/test/integration
	gotestsum --format testname --format-hide-empty-pkg  --junitfile junit.xml -- -p 1 $(GO_VENDOR) --tags=unit,integration -v -coverpkg=./pkg/... -covermode=count -coverprofile=c.out.tmp ./pkg/...

.PHONY: test/all/original
test/all/original: setup/dep/install setup/dep/test/install setup/test/integration
	$(GO) test -p 1 $(GO_VENDOR) --tags=unit,integration -v -coverpkg=./pkg/... -covermode=count -coverprofile=c.out.tmp ./pkg/...

.PHONY: test/unit
test/unit: setup/dep/install setup/dep/test/install
	$(GO) test $(GO_VENDOR) --tags=unit -v -coverpkg=./pkg/... -covermode=count -coverprofile=c.out.tmp ./pkg/...

.PHONY: test/integration
test/integration: setup/dep/install setup/dep/test/install setup/test/integration
	$(GO) test -p 1 $(GO_VENDOR) --tags=integration -v -coverpkg=./pkg/... -covermode=count -coverprofile=c.out.tmp ./pkg/...

.PHONY: test/coverage
test/coverage: setup/dep/test/install
	cat c.out.tmp | grep -v "mock_" | grep -v "generated" | grep -v "sql-for-tests\.go" > c.out
	$(GO) tool cover -html=c.out -o coverage.html
	$(GO) tool cover -func c.out
	gocover-cobertura < c.out > coverage.xml

#############
#   Setup   #
#############

.PHONY: down/services
down/services:
	@echo "Down services"
	$(DEFAULT_CONTAINER_RUNTIME) rm -f postgres || true
	$(DEFAULT_CONTAINER_RUNTIME) rm -f opa || true
	$(DEFAULT_CONTAINER_RUNTIME) rm -f keycloak || true
	$(DEFAULT_CONTAINER_RUNTIME) rm -f rabbitmq || true

.PHONY: down/metrics-services
down/metrics-services:
	@echo "Down metrics services"
	$(DEFAULT_CONTAINER_RUNTIME) rm -f prometheus || true
	$(DEFAULT_CONTAINER_RUNTIME) rm -f grafana || true
	$(DEFAULT_CONTAINER_RUNTIME) rm -f tempo || true

.PHONY: down/dev-services
down/dev-services:
	@echo "Down dev services"
	$(DEFAULT_CONTAINER_RUNTIME) rm -f pgadmin || true
	$(DEFAULT_CONTAINER_RUNTIME) rm -f maildev || true

.PHONY: setup/dev-services
setup/dev-services: down/dev-services
	@echo "Setup dev services"
	$(DEFAULT_CONTAINER_RUNTIME) run --rm --name pgadmin -p 8090:80 --link postgres:postgres -e 'PGADMIN_DEFAULT_EMAIL=user@domain.com' -e 'PGADMIN_DEFAULT_PASSWORD=SuperSecret' -d dpage/pgadmin4:$(PGADMIN_VERSION)
	$(DEFAULT_CONTAINER_RUNTIME) run --rm --name maildev -p 1080:1080 -p 1025:1025 -d maildev/maildev:$(MAILDEV_VERSION) --incoming-user fake --incoming-pass fakepassword

.PHONY: setup/metrics-services
setup/metrics-services: down/metrics-services
	@echo "Setup metrics services"
	$(DEFAULT_CONTAINER_RUNTIME) run --rm -d --name prometheus -v $(CURDIR)/.local-resources/prometheus/prometheus.yml:/prometheus/prometheus.yml --network=host prom/prometheus:$(PROMETHEUS_VERSION) --web.listen-address=:9191
	$(DEFAULT_CONTAINER_RUNTIME) run --rm -d --name grafana -e GF_AUTH_ANONYMOUS_ENABLED=true -e GF_AUTH_ANONYMOUS_ORG_ROLE=Admin -e GF_AUTH_DISABLE_LOGIN_FORM=true -e GF_FEATURE_TOGGLES_ENABLE=traceqlEditor -v $(CURDIR)/.local-resources/grafana/datasources.yaml:/etc/grafana/provisioning/datasources/datasources.yaml --network=host grafana/grafana:$(GRAFANA_VERSION)
	$(DEFAULT_CONTAINER_RUNTIME) run --rm -d --name tempo  -v $(CURDIR)/.local-resources/tempo/tempo.yaml:/etc/tempo.yaml -p 14268:14268 -p 3200:3200 -p 4318:4318 -p 4317:4317 grafana/tempo:$(TEMPO_VERSION) -config.file=/etc/tempo.yaml

.PHONY: setup/services
setup/services: down/services
	@echo "Setup services"
	mkdir -p .run/postgres || true
	$(DEFAULT_CONTAINER_RUNTIME) run -d --rm --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e PGDATA=/var/lib/postgresql/data/pgdata -v $(CURDIR)/.run/postgres:/var/lib/postgresql/data postgres:$(POSTGRESQL_VERSION)
	tar czvf .local-resources/opa/bundle.tar.gz --directory=.local-resources/opa/bundle example/
	$(DEFAULT_CONTAINER_RUNTIME) run -d --rm --name opa -p 8181:8181 -v $(CURDIR)/.local-resources/opa/bundle.tar.gz:/bundle.tar.gz openpolicyagent/opa:$(OPA_VERSION) run --server --log-level debug --log-format text --bundle /bundle.tar.gz
	$(DEFAULT_CONTAINER_RUNTIME) run -d --rm --name keycloak -p 8088:8080 -v $(CURDIR)/.local-resources/keycloak/realm-export.json:/opt/keycloak/data/import/realm-export.json -e KEYCLOAK_ADMIN=admin -e KEYCLOAK_ADMIN_PASSWORD=admin -e KC_HTTP_RELATIVE_PATH=/auth quay.io/keycloak/keycloak:$(KEYCLOAK_VERSION) start-dev --import-realm
	# User: guest / Password: guest
	$(DEFAULT_CONTAINER_RUNTIME) run --rm -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:$(RABBITMQ_VERSION)-management

.PHONY: setup/test/integration
setup/test/integration:
	$(DEFAULT_CONTAINER_RUNTIME) run --rm --network=host --name postgres-integration-setup -e PGHOST=localhost -e PGUSER=postgres -e PGPASSWORD=postgres -e PGDATABASE=postgres postgres:$(POSTGRESQL_VERSION) psql -c 'CREATE DATABASE "postgres-integration"' || true

.PHONY: setup/dep/all/install
setup/dep/all/install: setup/dep/lint/install setup/dep/gen/install setup/dep/test/install setup/dep/build/install

.PHONY: setup/dep/lint/install
setup/dep/lint/install: setup/dep/install
ifndef HAS_GOLANGCI_LINT
	@echo "=> Installing golangci-lint tool"
ifndef HAS_CURL
	$(error You must install curl)
endif
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)
endif

.PHONY: setup/dep/gen/install
setup/dep/gen/install:
ifndef HAS_GQLGEN
	@echo "=> Installing gqlgen tool"
	$(GO) install github.com/99designs/gqlgen@v0.17.63
endif
ifndef HAS_MOCKGEN
	@echo "=> Installing mockgen tool"
	$(GO) install go.uber.org/mock/mockgen@v0.5.0
endif
ifndef HAS_FIELDALIGNMENT
	@echo "=> Installing fieldalignment tool"
	$(GO) install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@v0.29.0
endif

.PHONY: setup/dep/test/install
setup/dep/test/install: setup/dep/install
ifndef HAS_GOTESTSUM
	@echo "=> Installing gotestsum tool"
	$(GO) install gotest.tools/gotestsum@v1.12.0
endif
ifndef HAS_GOCOVER_COBERTURA
	@echo "=> Installing gocover-cobertura tool"
	$(GO) install github.com/boumenot/gocover-cobertura@v1.3.0
endif

.PHONY: setup/dep/build/install
setup/dep/build/install: setup/dep/install
# Uncomment to use gox instead of goreleaser
# ifndef HAS_GOX
# 	@echo "=> Installing gox"
# 	$(GO) install github.com/mitchellh/gox@v1.0.1
# endif

.PHONY: setup/dep/install
setup/dep/install:
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
