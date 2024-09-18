# Variables
BINARY_NAME=ascii-converter
BUILD_DIR=./build
SRC_DIR=./cmd/server
TEMPLATES_DIR=./templates

# Default target
all: build

# Build the Go project
build:
	@echo "Building the project..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)/main.go

# Run the project
run: build
	@echo "Running the project..."
	@$(BUILD_DIR)/$(BINARY_NAME)

# Clean build files
clean:
	@echo "Cleaning build directory..."
	@rm -rf $(BUILD_DIR)

# Format Go code
fmt:
	@echo "Formatting Go code..."
	@go fmt ./...

# Update dependencies
deps:
	@echo "Updating dependencies..."
	@go mod tidy

.PHONY: all build run clean fmt deps
