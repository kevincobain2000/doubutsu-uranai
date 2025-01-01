#!/bin/sh

# Variables
APP_NAME="doubutsu-uranai"
# Pick the latest tag from git
VERSION=$(git describe --tags --abbrev=0)
OUTPUT_DIR="dist"

rm -rf "$OUTPUT_DIR"

# List of target platforms and architectures
TARGETS="darwin amd64
darwin arm64
freebsd amd64
linux amd64
linux arm
linux arm64
"

# Create output directory
mkdir -p "$OUTPUT_DIR"

echo "Building binaries for $APP_NAME version $VERSION..."

go install mvdan.cc/garble@latest

# Loop through each target
echo "$TARGETS" | while read -r PLATFORM ARCH; do
    # Skip empty lines
    if [ -z "$PLATFORM" ] || [ -z "$ARCH" ]; then
        continue
    fi

    OUTPUT_FILE="$OUTPUT_DIR/${APP_NAME}-${PLATFORM}-${ARCH}"

    # Set environment variables and build
    if ! GOOS=$PLATFORM GOARCH=$ARCH garble build -o "$OUTPUT_FILE"; then
        echo "Failed to build for $PLATFORM $ARCH"
        exit 1
    fi

    echo "Built: $OUTPUT_FILE"
done

# Generate checksums
CHECKSUM_FILE="$OUTPUT_DIR/${APP_NAME}_${VERSION}_checksums.txt"
echo "Generating checksums..."

if ! shasum -a 256 "$OUTPUT_DIR"/* > "$CHECKSUM_FILE"; then
    echo "Failed to generate checksums"
    exit 1
fi

echo "Checksums written to $CHECKSUM_FILE"
echo "Build and packaging completed!"
