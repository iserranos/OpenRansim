export PATH := $(GOPATH)/bin:./scripts:$(PATH)
ifeq (run,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

.PHONY: all
all: fmt install build

.PHONY: install
install:
	@echo "Install dependencies"
	@go get ./...

.PHONY: run
run:
	ifeq (,$(wildcard OpenRansim))
	@echo "To run OpenRansim first run 'make all'"
	@for e in $(RUN_ARGS) ; do echo Running $$e module && ./OpenRansim $$e ; done
	else
	@echo "To run OpenRansim first run 'make all'"
	endif

.PHONY: build
build:
	@echo "Building OpenRansim"
	go build -o OpenRansim

.PHONY: fmt
fmt:
	@echo "Formatting all go code..."
	go fmt `go list ./... | grep -v vendor`

.PHONY: clean
clean:
	@echo "Cleaning build..."
	@rm -rf OpenRansim

