# -----------------------------------------------------------------
#
#        ENV VARIABLE
#
# -----------------------------------------------------------------

# go env vars
GO=$(firstword $(subst :, ,$(GOPATH)))
# list of pkgs for the project without vendor
PKGS=$(shell go list ./... | grep -v /vendor/)
MAIN_FILE=main.go
export GO15VENDOREXPERIMENT=1


# -----------------------------------------------------------------
#        Main targets
# -----------------------------------------------------------------

help:
	@echo
	@echo "----- BUILD ------------------------------------------------------------------------------"
	@echo "all                  clean and build the project"
	@echo "clean                clean the project"
	@echo "dependencies         download the dependencies"
	@echo "build                build all libraries and binaries"
	@echo "----- TESTS && LINT ----------------------------------------------------------------------"
	@echo "test                 test all packages"
	@echo "format               format all packages"
	@echo "lint                 lint all packages"
	@echo "----- SERVERS AND DEPLOYMENTS ------------------------------------------------------------"
	@echo "start                start process on localhost"
	@echo "----- OTHERS -----------------------------------------------------------------------------"
	@echo "help                 print this message"

all: clean build

clean:
	@go clean
	@rm -Rf .tmp
	@rm -Rf .DS_Store
	@rm -Rf *.log
	@rm -Rf *.out
	@rm -Rf *.lock
	@rm -Rf *.mem
	@rm -Rf *.test
	@rm -Rf build

dependencies:
	@echo
	@echo "----- DOWNLOADING -------------------------------------------------------------------------"
	@go get -u github.com/gorilla/mux
	@go get -u github.com/gorilla/context
	@go get -u gopkg.in/mgo.v2
	@go get -u github.com/tools/godep
	@go get -u github.com/golang/lint/golint
	@echo "----- DONE --------------------------------------------------------------------------------"

build: format
	@go build -o ifi main.go

format:
	@go fmt $(PKGS)

lint:
	@golint controllers/...
	@golint models/...
	@golint routes/...
	@golint tests/...
	@golint utils/...
	@golint ./.
	@go vet $(PKGS)

start:
	@go run $(MAIN_FILE)

