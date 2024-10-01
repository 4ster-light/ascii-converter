# Variables
BINARY_NAME=ascii-converter
BUILD_DIR=./build
SRC_DIR=./backend/server
FRONTEND_DIR=./frontend
FRONTEND_BUILD_DIR=$(FRONTEND_DIR)/dist

# Default target
all: build

# Build the Go project
build:
	echo "Building the project..."
	cd $(FRONTEND_DIR) && bun run build
	mkdir -p $(BUILD_DIR)
	cp -r $(FRONTEND_BUILD_DIR) $(BUILD_DIR)/frontend
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)/main.go

# Run the project
run: build
	echo "Running the project..."
	$(BUILD_DIR)/$(BINARY_NAME)

# Clean build files
clean:
	echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)
	rm -rf $(FRONTEND_BUILD_DIR)

# Format Go code
fmt:
	echo "Formatting Go code..."
	go fmt ./...

# Update dependencies
deps:
	echo "Updating dependencies..."
	go mod tidy

.PHONY: all build run clean fmt deps
