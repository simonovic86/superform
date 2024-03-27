#!/bin/bash

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo "Go could not be found. Please install Go and try again."
    exit 1
fi

# Define the build output and bin directory
OUTPUT_NAME="superformcli"
BIN_DIR="./bin"

# Create bin directory if it does not exist
if [ ! -d "$BIN_DIR" ]; then
    mkdir -p "$BIN_DIR"
fi

# Build the Go program
echo "Building the tool..."
go build -o $OUTPUT_NAME cmd/superformcli/main.go

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo "Build successful! Moving the tool to $BIN_DIR"
    # Move the binary to the bin directory
    mv ./$OUTPUT_NAME $BIN_DIR/
else
    echo "Build failed."
    exit 1
fi
