#!/bin/bash
set -e

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ]; then
  ARCH="amd64"
fi

URL="https://github.com/IrynaMazaieva/weather-cli/releases/latest/download/weather-cli-$OS"

echo "Downloading weather-cli for $OS..."
curl -L -o /usr/local/bin/weather-cli "$URL"
chmod +x /usr/local/bin/weather-cli

echo "Installation complete! Run 'weather-cli' to get started."
