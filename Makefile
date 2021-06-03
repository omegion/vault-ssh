export PATH := $(abspath ./vendor/bin):$(PATH)

BASE_PACKAGE_NAME	= github.com/omegion/vault-ssh
GIT_VERSION 		= $(shell git describe --tags --always 2> /dev/null || echo 0.0.0)
LDFLAGS           	= -ldflags "-X $(BASE_PACKAGE_NAME)/internal/info.Version=$(GIT_VERSION)"
BUFFER            	:= $(shell mktemp)
REPORT_DIR         	= dist/report
COVER_PROFILE      	= $(REPORT_DIR)/coverage.out
TARGETOS			= "darwin"
TARGETARCH			= "amd64"

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) go build $(LDFLAGS) -a -installsuffix cgo -o dist/vault-ssh main.go

.PHONY: lint
lint:
	@echo "Checking code style"
	gofmt -l . | tee $(BUFFER)
	@! test -s $(BUFFER)
	go vet ./...
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1
	@golangci-lint --version
	golangci-lint run --fix
	go get -u golang.org/x/lint/golint
	golint -set_exit_status ./...

.PHONY: test
test:
	@echo "Running unit tests"
	mkdir -p $(REPORT_DIR)
	go test -covermode=count -coverprofile=$(COVER_PROFILE) -tags test -failfast ./...
	go tool cover -html=$(COVER_PROFILE) -o $(REPORT_DIR)/coverage.html

.PHONY: cut-tag
cut-tag:
	@echo "Cutting $(version)"
	git tag $(version)
	git push origin $(version)

.PHONY: release
release: build
	@echo "Releasing $(GIT_VERSION)"
	docker build -t vault-ssh .
	docker tag vault-ssh:latest omegion/vault-ssh:$(GIT_VERSION)
	docker push omegion/vault-ssh:$(GIT_VERSION)

.PHONY: docker-image
docker-image:
	@echo "Building Docker Image"
	docker buildx build -t vault-ssh --platform linux/amd64,linux/arm64 . --output=type=docker
