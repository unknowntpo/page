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