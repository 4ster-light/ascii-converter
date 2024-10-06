# Variables
BINARY_NAME=ascii-converter
BUILD_DIR=./build
SRC_DIR=./backend/server
FRONTEND_DIR=./frontend
FRONTEND_BUILD_DIR=$(FRONTEND_DIR)/dist

all:
	@./build.sh

run:
	@./build.sh
	@echo "Running the project..."
	@$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning build directory..."
	@rm -rf $(BUILD_DIR)
	@rm -rf $(FRONTEND_BUILD_DIR)

fmt:
	@echo "Formatting Go code..."
	@go fmt ./...

deps:
	@echo "Updating dependencies..."
	@go mod tidy

.PHONY: all run clean fmt deps
