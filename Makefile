GO := go
GOGET := $(GO) get
GOGET_FLAGS := -v
GOLIST := $(GO) list
GOBUILD := $(GO) build
GOBUILD_FLAGS := -v
GOTEST := ginkgo
GOTEST_FLAGS := -gcflags=-l
TARGET := acm-training
GOLINT := golint
GOLINT_FLAG := -set_exit_status

COVERAGE_REPORT := coverage.txt
TEST_SOURCES := $(shell find . -name "*.go")
SOURCES := $(filter-out $(shell find . -name "*_test.go"), $(TEST_SOURCES))

.PHONY: all
all: $(TARGET)

get:
	$(GOGET) $(GOGET_FLAGS) ./...
	$(GOGET) $(GOGET_FLAGS) golang.org/x/lint/golint
	$(GOGET) $(GOGET_FLAGS) github.com/ory/go-acc
	$(GOGET) $(GOGET_FLAGS) github.com/onsi/ginkgo/ginkgo
	$(GOGET) $(GOGET_FLAGS) github.com/onsi/gomega/...

$(TARGET): $(SOURCES)
	@echo "Building the application..."
	@$(GOBUILD) $(GOBUILD_FLAGS) ./...
	@$(GOBUILD) -o $@ $(GOBUILD_FLAGS)

.PHONY: lint
lint: $(TEST_SOURCES)
	@echo "Perform linting..."
	@$(GOLIST) ./... | grep -v /app/model/ | xargs -L1 $(GOLINT) $(GOLINT_FLAG)

.PHONY: test
test: $(TEST_SOURCES) lint
	@echo "Perform testing..."
	@$(GOTEST) $(GOTEST_FLAGS) ./...

.PHONY: coverage
coverage: $(TEST_SOURCES) lint
	@echo "Generating coverage report..."
	@go-acc ./...

.PHONY: coverage-report
coverage-report: coverage
	$(GO) tool cover --html=$(COVERAGE_REPORT)

.PHONY: clean
clean:
	@echo "Perform cleanup..."
	@rm -f $(TARGET) $(COVERAGE_REPORT)
