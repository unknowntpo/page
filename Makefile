# check installation of githooks and display help message when typing make
all: help

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## mock/gen: generate mock $(IFASE) implementation against interface inside internal/domain, e.g. make mock/gen IFASE=PageUsecase
mock/gen:
	mockgen -source ./domain/page.go \
		-destination ./domain/mock/$(IFASE).go \
		-package mock \
		$(IFASE)

## proto/gen: generate code from grpc proto
proto/gen:
	buf generate

## redis/setup: set up development environment
redis/setup:
	docker run -d --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:6.2.6-v6

## redis/flush: wipe out data in redis
redis/flush:
	docker exec -it redis-stack redis-cli -c 'FLUSHALL'

## redis/down: delete redis container
redis/down:
	docker rm -f redis-stack

TESTPKG ?= ./...

## test: run unit tests
test:
	go test $(if $(VERBOSE),-v) -p 1 \
	 -count 1 $(TESTPKG) \
	 -cover \
	 $(if $(FOCUS), -ginkgo.focus '$(FOCUS)')

## build: build the server binary
build:
	go build -o bin/server ./cmd/server

## run/server: build and run the server binary
run/server: build
	./bin/server

## run/client: build and run the client binary
run/client:
	go run ./cmd/client/client.go

