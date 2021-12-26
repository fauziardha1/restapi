pkgs = $(shell go list ./... | grep -v vendor | grep -v mocks)
reponame = $(shell basename `git rev-parse --show-toplevel`)
branch = $(shell git rev-parse --abbrev-ref HEAD)
# sonarqube_url := $(shell echo http://172.31.7.10:9000/dashboard?branch=${branch})

.PHONY: build

dep:
	@echo "RUNNING DEP..."
	@dep ensure -vendor-only -v

mod:
	@echo "go mod tidy..."
	@go mod tidy
	@echo "go mod vendor -v..."
	@go mod vendor -v
	@echo go fmt ./...
	@go fmt ./...

build:
	@echo "Reformatting..."
	@go fmt ./...
	@echo "BUILD restapi..."
	@go build -v -o restapi cmd/restapi/*.go

run:
	@echo "RUN restapi..."
	make build
	@./restapi --module=$(module)

test:
	@echo "RUN TESTING..."
	@go clean -testcache 
	@go test -v -p=2 -cover -race $(pkgs) -coverprofile coverage.out

sonar:
	make test
	@.cmd/sonar.sh $(reponame) $(branch)

lint:
	@echo "RUN Lint checker..."
	@golint ./cmd/... ./internal/...
