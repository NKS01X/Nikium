#!/bin/sh
set -e

# Config
REPO="NKS01X/Nikium"
BIN_NAME="nikium"

# Colors
RED='\033[31m'
GREEN='\033[32m'
BLUE='\033[34m'
RESET='\033[0m'

echo "${BLUE}Downloading Nikium...${RESET}"

# OS detect
OS_TYPE=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$OS_TYPE" in
    linux*) OS_PATTERN="linux" ;;
    darwin*) OS_PATTERN="darwin" ;;
    *) echo "${RED}OS $OS_TYPE not supported.${RESET}"; exit 1 ;;
esac

# Arch detect
ARCH_TYPE=$(uname -m)
case "$ARCH_TYPE" in
    x86_64|amd64) ARCH_PATTERN="(x86_64|amd64)" ;;
    arm64|aarch64) ARCH_PATTERN="(arm64|aarch64)" ;;
    i386|i686) ARCH_PATTERN="(i386|386)" ;;
    *) echo "${RED}Arch $ARCH_TYPE not supported.${RESET}"; exit 1 ;;
esac

# Query GitHub API
API_URL="https://api.github.com/repos/$REPO/releases/latest"
DOWNLOAD_URL=$(curl -s "$API_URL" | grep "browser_download_url" | grep -iE "$OS_PATTERN" | grep -iE "$ARCH_PATTERN" | grep "\.tar\.gz" | head -n 1 | cut -d '"' -f 4)

if [ -z "$DOWNLOAD_URL" ]; then
    echo "${RED}Cannot find release for $OS_TYPE $ARCH_TYPE.${RESET}"
    exit 1
fi

echo "${BLUE}Found release: $DOWNLOAD_URL${RESET}"

# Download & extract
TMP_DIR=$(mktemp -d)
curl -sL "$DOWNLOAD_URL" -o "$TMP_DIR/nikium.tar.gz"
tar -xzf "$TMP_DIR/nikium.tar.gz" -C "$TMP_DIR"

# Find extracted bin
EXTRACTED_BIN=$(find "$TMP_DIR" -type f -name "$BIN_NAME" | head -n 1)

if [ -z "$EXTRACTED_BIN" ]; then
    echo "${RED}Binary $BIN_NAME not found in archive.${RESET}"
    rm -rf "$TMP_DIR"
    exit 1
fi

# Pick install dir
if [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
else
    INSTALL_DIR="$HOME/.local/bin"
    mkdir -p "$INSTALL_DIR"
fi

# Move bin
mv "$EXTRACTED_BIN" "$INSTALL_DIR/$BIN_NAME"
chmod +x "$INSTALL_DIR/$BIN_NAME"

# Clean
rm -rf "$TMP_DIR"

echo "${GREEN}Successfully installed!${RESET}"
echo "Nikium is located at $INSTALL_DIR/$BIN_NAME"

# PATH warning
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo "${RED}Add $INSTALL_DIR to your PATH.${RESET}"
fi
