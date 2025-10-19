.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c

BINARY_NAME=nepse_go
SRC_DIR=./cmd/examples
BUILD_DIR=./bin
GO=go
GOFLAGS=-ldflags="-s -w"
TEST_FLAGS?=-parallel 4
SOURCES=$(shell find . -type f -name '*.go') go.mod go.sum

# Default target
all: build

# Build the binary with dependency tracking
$(BUILD_DIR)/$(BINARY_NAME): $(SOURCES)
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)/...

# Build target
build: $(BUILD_DIR)/$(BINARY_NAME)

# Run the application
run:
	@echo "Running $(BINARY_NAME)..."
	@$(GO) run $(SRC_DIR)/...

# Test the project
test:
	@echo "Running tests..."
	@$(GO) test ./... -v $(TEST_FLAGS)

# Format the code
fmt:
	@echo "Formatting code..."
	@$(GO) fmt ./...

# Vet the code
vet:
	@echo "Vetting code..."
	@$(GO) vet ./...

# Clean the build directory
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@$(GO) mod tidy

# Build and run
start: build
	@echo "Starting $(BINARY_NAME)..."
	@$(BUILD_DIR)/$(BINARY_NAME)

# Phony targets
.PHONY: all build run test fmt vet clean deps start
