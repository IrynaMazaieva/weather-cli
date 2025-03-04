#!/bin/bash
set -e

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ]; then
  ARCH="amd64"
fi

URL="https://github.com/IrynaMazaieva/weather-cli/releases/latest/download/weather-cli-$OS"


INSTALL_DIR="$HOME/.local/bin"
mkdir -p "$INSTALL_DIR"

echo "Downloading weather-cli for $OS..."

curl -L -o "$INSTALL_DIR/weather-cli" "$URL"
chmod +x "$INSTALL_DIR/weather-cli"

echo "export PATH=\"$INSTALL_DIR:\$PATH\"" >> ~/.bashrc
source ~/.bashrc

echo "Installation complete! Run 'weather-cli' to get started."
