#!/bin/bash
set -e

# Config
REPO="NKS01X/Nikium"
BIN_NAME="nikium"

# Colors
RED='\033[31m'
GREEN='\033[32m'
BLUE='\033[34m'
RESET='\033[0m'

echo -e "${BLUE}Downloading Nikium...${RESET}"

# OS detect
OS_TYPE=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$OS_TYPE" in
    linux*) OS_PATTERN="linux" ;;
    darwin*) OS_PATTERN="darwin" ;;
    *) echo -e "${RED}OS $OS_TYPE not supported.${RESET}"; exit 1 ;;
esac

# Arch detect
ARCH_TYPE=$(uname -m)
case "$ARCH_TYPE" in
    x86_64|amd64) ARCH_PATTERN="(x86_64|amd64)" ;;
    arm64|aarch64) ARCH_PATTERN="(arm64|aarch64)" ;;
    i386|i686) ARCH_PATTERN="(i386|386)" ;;
    *) echo -e "${RED}Arch $ARCH_TYPE not supported.${RESET}"; exit 1 ;;
esac

# Query GitHub API
API_URL="https://api.github.com/repos/$REPO/releases/latest"
DOWNLOAD_URL=$(curl -s "$API_URL" | grep "browser_download_url" | grep -iE "$OS_PATTERN" | grep -iE "$ARCH_PATTERN" | grep "\.tar\.gz" | head -n 1 | cut -d '"' -f 4)

if [ -z "$DOWNLOAD_URL" ]; then
    echo -e "${RED}Cannot find release for $OS_TYPE $ARCH_TYPE.${RESET}"
    exit 1
fi

echo -e "${BLUE}Found release: $DOWNLOAD_URL${RESET}"

# Download & extract
TMP_DIR=$(mktemp -d)
curl -sL "$DOWNLOAD_URL" -o "$TMP_DIR/nikium.tar.gz"
tar -xzf "$TMP_DIR/nikium.tar.gz" -C "$TMP_DIR"

# Find extracted bin (Using -iname for case-insensitivity)
EXTRACTED_BIN=$(find "$TMP_DIR" -type f -iname "$BIN_NAME" | head -n 1)

if [ -z "$EXTRACTED_BIN" ]; then
    echo -e "${RED}Binary $BIN_NAME not found in archive.${RESET}"
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

# Move bin (Forces it to be lowercase 'nikium' in the system)
mv "$EXTRACTED_BIN" "$INSTALL_DIR/nikium"
chmod +x "$INSTALL_DIR/nikium"

# Clean
rm -rf "$TMP_DIR"

echo -e "${GREEN}Successfully installed!${RESET}"
echo "Nikium is located at $INSTALL_DIR/nikium"

# PATH warning
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo -e "${RED}Please add $INSTALL_DIR to your PATH.${RESET}"
fi