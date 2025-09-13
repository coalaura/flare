#!/bin/bash
set -e

VERSION="v0.0.1"
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    armv7l) ARCH="arm" ;;
    *) printf "\033[0;31mUnsupported architecture: %s\033[0m\n" $ARCH$; exit 1 ;;
esac

printf "Downloading flare_${VERSION}_${ARCH}...\n"

sudo curl -fsSL "https://github.com/coalaura/flare/releases/download/${VERSION}/flare_${VERSION}_${ARCH}" -o /usr/local/bin/flare
sudo chmod +x /usr/local/bin/flare

printf "Testing installation...\n"

flare

printf "Setting up cron job...\n"

printf "0 */6 * * * /usr/local/bin/flare\n" > /etc/cron.d/flare

printf "\033[0;32mInstalled successfully!\033[0m\n"