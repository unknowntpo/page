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

## local/run: run service at local

## mock/gen: generate mock $(IFASE) implementation against interface inside internal/domain, e.g. make mock/gen IFASE=PageUsecase
mock/gen:
	mockgen -source ./internal/domain/page.go \
		-destination internal/domain/mock/$(IFASE).go \
		-package mock \
		$(IFASE)

## proto/gen: generate code from grpc proto
proto/gen:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./internal/api/page/grpc/page/page.proto