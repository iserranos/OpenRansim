export PATH := $(GOPATH)/bin:./scripts:$(PATH)
ifeq (run,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

all: fmt install build

install:
	@echo "Install dependencies"
	@go get -u github.com/spf13/cobra

run:
	ifeq (,$(wildcard OpenRansim))
	@echo "To run OpenRansim first run 'make all'"
	@for e in $(RUN_ARGS) ; do echo Running $$e module && ./OpenRansim $$e ; done
	else
	@echo "To run OpenRansim first run 'make all'"
	endif

build:
	@echo "Building OpenRansim"
	go build -o OpenRansim

fmt:
	@echo "Formatting all go code..."
	go fmt `go list ./... | grep -v vendor`

clean:
	@echo "Cleaning build..."
	@rm -rf OpenRansim
