APP_NAME = mobile-security-service
ORG_NAME = aerogear
PKG = github.com/$(ORG_NAME)/$(APP_NAME)
APP_FILE=./cmd/mobile-security-service/main.go
APP_FILE_DIR=cmd/mobile-security-service
TOP_SRC_DIRS = pkg
PACKAGES     ?= $(shell sh -c "find $(TOP_SRC_DIRS) -name \\*_test.go \
                   -exec dirname {} \\; | sort | uniq")
BIN_DIR := $(GOPATH)/bin				   
BINARY ?= mobile-security-service
RELEASE_TAG = $(CIRCLE_TAG)

# This follows the output format for goreleaser
BINARY_LINUX_64 = ./dist/linux_amd64/$(BINARY)

IMAGE_REGISTRY=quay.io
IMAGE_LATEST_TAG = $(IMAGE_REGISTRY)/$(ORG_NAME)/$(APP_NAME):latest
IMAGE_MASTER_TAG = $(IMAGE_REGISTRY)/$(ORG_NAME)/$(APP_NAME):master
IMAGE_RELEASE_TAG = $(IMAGE_REGISTRY)/$(ORG_NAME)/$(APP_NAME):$(RELEASE_TAG)

LDFLAGS=-ldflags "-w -s -X main.Version=${TAG}"


# SERVER
# SERVER setup
.PHONY: setup
setup: setup-githooks
	@echo Installing application dependencies:
	dep ensure
	make build-swagger-api

.PHONY: build-swagger-api
build-swagger-api:
	@echo Installing Swagger dep:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
	@echo Updating Swagger api:
	cd $(APP_FILE_DIR); swagger generate spec -m -o ../../api/swagger.yaml

.PHONY: setup-githooks
setup-githooks:
	@echo Installing errcheck dependence:
	go get -u github.com/kisielk/errcheck
	@echo Setting up Git hooks:
	ln -sf $$PWD/.githooks/* $$PWD/.git/hooks/

# SERVER build/release
.PHONY: build
build: setup
	go build -o $(BINARY) $(APP_FILE)

.PHONY: build-linux
build-linux: setup
	env GOOS=linux GOARCH=amd64 go build -o $(BINARY_LINUX_64) $(APP_FILE)

.PHONY: build-image
build-image: build-linux
	docker build -t $(IMAGE_LATEST_TAG) --build-arg BINARY=$(BINARY_LINUX_64) .

.PHONY: build-release-image
build-release-image:
	docker build -t $(IMAGE_LATEST_TAG) -t $(IMAGE_RELEASE_TAG) --build-arg BINARY=$(BINARY_LINUX_64) .

.PHONY: build-master-image
build-master-image:
	docker build -t $(IMAGE_MASTER_TAG) --build-arg BINARY=$(BINARY_LINUX_64) .

.PHONY: push-release-image
push-release-image:
	@docker login --username $(QUAY_USERNAME) --password $(QUAY_PASSWORD) $(IMAGE_REGISTRY)
	docker push $(IMAGE_LATEST_TAG)
	docker push $(IMAGE_RELEASE_TAG)

.PHONY: push-master-image
push-master-image:
	@docker login --username $(QUAY_USERNAME) --password $(QUAY_PASSWORD) $(IMAGE_REGISTRY)
	docker push $(IMAGE_MASTER_TAG)

.PHONY: release
release: setup
	goreleaser --rm-dist

# SERVER test
.PHONY: test-all
test-all: test-unit
	make test-integration

.PHONY: test
test: test-unit

.PHONY: test-unit
test-unit:
	@echo Running tests:
	GOCACHE=off go test -cover \
	  $(addprefix $(PKG)/,$(PACKAGES))

.PHONY: test-integration
test-integration:
	@echo Running tests:
	GOCACHE=off go test -failfast -cover -tags=integration \
	  $(addprefix $(PKG)/,$(PACKAGES))

.PHONY: test-integration-cover
test-integration-cover:
	echo "mode: count" > coverage-all.out
	GOCACHE=off $(foreach pkg,$(PACKAGES),\
		go test -failfast -tags=integration -coverprofile=coverage.out -covermode=count $(addprefix $(PKG)/,$(pkg)) || exit 1;\
		tail -n +2 coverage.out >> coverage-all.out;)
	make cleanup-coverage-file

.PHONY: cleanup-coverage-file
cleanup-coverage-file:
	@echo "Cleaning up output of coverage report"
	./scripts/cleanup-coverage-file.sh

# SERVER misc
.PHONY: generate
generate:
	go generate $(APP_FILE)

.PHONY: errcheck
errcheck:
	@echo errcheck
	@errcheck -ignoretests $$(go list ./...)

.PHONY: vet
vet:
	@echo go vet
	go vet $$(go list ./... | grep -v /vendor/)

.PHONY: fmt
fmt:
	@echo go fmt
	go fmt $$(go list ./... | grep -v /vendor/)

.PHONY: clean
clean:
	-rm -f ${BINARY}
	-rm -rf .vendor-new
	-rm -rf vendor/

## UI
.PHONY:
ui-npm-install:
	cd ui && npm install

.PHONY:
ui-npm-ci:
	cd ui && npm ci

.PHONY: ui
ui: ui-npm-ci
	cd ui && npm run build

.PHONY: ui-check-code-style
ui-check-code-style: ui
	cd ui && npm run lint

.PHONY: ui-test-cover
ui-test-cover: ui-npm-ci
	cd ui && npm run coverage

.PHONY: serve
serve: build ui
	export STATIC_FILES_DIR=$(CURDIR)/ui/build; ./mobile-security-service -kubeconfig ~/.kube/config