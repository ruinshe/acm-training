GO := go
GOGET := $(GO) get
GOGET_FLAGS := -v
GOBUILD := $(GO) build
GOBUILD_FLAGS := -v
GOTEST := $(GO) test
GOTEST_FLAGS := -v
TARGET := acm-training
GOLINT := golint

COVERAGE_REPORT := coverage.txt
TEST_SOURCES := $(wildcard **/*.go) main.go
SOURCES := $(filter-out $(wilcard **/*_test.go), $(TEST_SOURCES))

.PHONY: all
all: $(TARGET)

get:
	$(GOGET) $(GOGET_FLAGS) ./...
	$(GOGET) $(GOGET_FLAGS) golang.org/x/lint/golint
	$(GOGET) $(GOGET_FLAGS) github.com/ory/go-acc

$(TARGET): $(SOURCES)
	@echo "Building the application..."
	@$(GOBUILD) $(GOBUILD_FLAGS) ./...
	@$(GOBUILD) -o $@ $(GOBUILD_FLAGS)

.PHONY: lint
lint: $(TEST_SOURCES)
	@echo "Perform linting..."
	@$(GOLINT) ./...

.PHONY: test
test: $(TEST_SOURCES) lint
	@echo "Perform testing..."
	@$(GOTEST) $(GOTEST_FLAGS) ./...

.PHONY: coverage
coverage: $(TEST_SOURCES) lint
	@echo "Generating coverage report..."
	@go-acc github.com/uestc-acm/acm-training/...

.PHONY: clean
clean:
	@echo "Perform cleanup..."
	@rm -f $(TARGET) $(COVERAGE_REPORT)
