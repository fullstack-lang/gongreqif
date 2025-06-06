#!/bin/bash

# Store the current directory
ORIGINAL_DIR=$(pwd)

# Navigate to the specified directory
cd ../../fullstack-lang/gong

# Get the latest commit hash
LATEST_COMMIT=$(git rev-parse HEAD)

# Return to the original directory
cd "$ORIGINAL_DIR"

# Perform go get with the latest commit hash
echo "Updating module with latest commit: $LATEST_COMMIT"
go get github.com/fullstack-lang/gong@"$LATEST_COMMIT"

gongc go/models

echo "Module updated successfully!"