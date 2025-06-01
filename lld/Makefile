# LLD Project Makefile
# Usage: make run parkinglot, make run stackoverflow, etc.

# Binary output directory
BIN_DIR := bin

.PHONY: run build clean list help

# Run command with problem name as argument
run:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "Usage: make run <problem>"; \
		echo "Available problems:"; \
		make list; \
		exit 1; \
	fi
	@PROBLEM=$(filter-out $@,$(MAKECMDGOALS)); \
	if [ ! -d "$$PROBLEM" ]; then \
		echo "Error: Problem '$$PROBLEM' not found!"; \
		make list; \
		exit 1; \
	fi; \
	echo "Building and running $$PROBLEM..."; \
	mkdir -p $(BIN_DIR); \
	go build -o $(BIN_DIR)/$$PROBLEM ./$$PROBLEM; \
	./$(BIN_DIR)/$$PROBLEM

# Build command with problem name as argument
build:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "Usage: make build <problem>"; \
		echo "Available problems:"; \
		make list; \
		exit 1; \
	fi
	@PROBLEM=$(filter-out $@,$(MAKECMDGOALS)); \
	if [ ! -d "$$PROBLEM" ]; then \
		echo "Error: Problem '$$PROBLEM' not found!"; \
		exit 1; \
	fi; \
	echo "Building $$PROBLEM..."; \
	mkdir -p $(BIN_DIR); \
	go build -o $(BIN_DIR)/$$PROBLEM ./$$PROBLEM

# Build all problems
build-all:
	@echo "Building all problems..."
	@mkdir -p $(BIN_DIR)
	@for dir in */; do \
		if [ -f "$$dir/main.go" ] || [ -f "$$dir"*.go ]; then \
			problem=$$(basename "$$dir"); \
			echo "Building $$problem..."; \
			go build -o $(BIN_DIR)/$$problem ./$$problem; \
		fi \
	done

# Clean binaries
clean:
	@echo "Cleaning binaries..."
	@rm -rf $(BIN_DIR)

# List available problems
list:
	@echo "Available problems:"
	@for dir in */; do \
		if [ -f "$$dir/main.go" ] || [ -f "$$dir"*.go ]; then \
			echo "  - $$(basename "$$dir")"; \
		fi \
	done

# Test a specific problem
test:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "Usage: make test <problem>"; \
		echo "Available problems:"; \
		make list; \
		exit 1; \
	fi
	@PROBLEM=$(filter-out $@,$(MAKECMDGOALS)); \
	if [ ! -d "$$PROBLEM" ]; then \
		echo "Error: Problem '$$PROBLEM' not found!"; \
		exit 1; \
	fi; \
	echo "Testing $$PROBLEM..."; \
	go test ./$$PROBLEM -v

# Test all problems
test-all:
	@echo "Testing all problems..."
	@go test ./... -v

# Help
help:
	@echo "LLD Project Commands:"
	@echo "  make run <problem>       - Build and run a specific problem"
	@echo "  make build <problem>     - Build a specific problem"
	@echo "  make test <problem>      - Test a specific problem"
	@echo "  make build-all           - Build all problems"
	@echo "  make test-all            - Test all problems"
	@echo "  make list                - List available problems"
	@echo "  make clean               - Clean all binaries"
	@echo "  make help                - Show this help"
	@echo ""
	@echo "Examples:"
	@echo "  make run parkinglot"
	@echo "  make run stackoverflow"
	@echo "  make run elevator"

# Treat problem names as targets to avoid make errors
%:
	@: