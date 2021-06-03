GOCMD=go
GOTEST=$(GOCMD) test
BINARY=playing_cards
VERSION?=0.1.0
SERVER_PORT?=3000
EXPORT_RESULT?=false

all: build

build:
	mkdir -p out/bin
	GO111MODULE=on $(GOCMD) build -mod vendor -o out/bin/$(BINARY) src/*.go

vendor:
	$(GOCMD) mod vendor

clean:
	rm -rf ./bin
	rm -rf ./out

test: ## Run the tests of the project
ifeq ($(EXPORT_RESULT), true)
	GO111MODULE=off go get -u github.com/jstemmer/go-junit-report
	$(eval OUTPUT_OPTIONS = | tee /dev/tty | go-junit-report -set-exit-code > junit-report.xml)
endif
	$(GOTEST) -v -race ./... $(OUTPUT_OPTIONS)

coverage: ## Run the tests of the project and export the coverage
	$(GOTEST) -cover -covermode=count -coverprofile=profile.cov ./...
	$(GOCMD) tool cover -func profile.cov
ifeq ($(EXPORT_RESULT), true)
	GO111MODULE=off go get -u github.com/AlekSi/gocov-xml
	GO111MODULE=off go get -u github.com/axw/gocov/gocov
	gocov convert profile.cov | gocov-xml > coverage.xml
endif

run: build
	./out/bin/$(BINARY)

