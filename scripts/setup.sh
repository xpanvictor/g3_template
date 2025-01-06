#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 github.com/username/repo-name"
    exit 1
fi

NEW_MODULE=$1
TEMPLATE_NAME="github.com/xpanvictor/g3_template.git"

# Replace module name in go.mod
sed -i '' "s|${TEMPLATE_NAME}|${NEW_MODULE}|g" go.mod

# Find and replace in all .go files
find . -type f -name "*.go" -exec sed -i '' "s|${TEMPLATE_NAME}|${NEW_MODULE}|g" {} +

echo "Project setup complete for ${NEW_MODULE}"
