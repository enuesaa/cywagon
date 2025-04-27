#!/bin/bash

# Replace `${VERSION}` with the latest value in GitHub Actions.
VERSION="${VERSION}"; # like `v0.0.3`

if [[ "$VERSION" != v* ]]; then
  echo "Error: Incorrect install script";
  exit 1;
fi

# strip `v`
VERSION_NUMBER="${VERSION#v}"; # like `0.0.3`

OS=$(uname | tr '[:upper:]' '[:lower:]');
ARCH=$(uname -m);

if [ "$ARCH" = "x86_64" ]; then
  ARCH="amd64";
elif [ "$ARCH" = "aarch64" ]; then
  ARCH="arm64";
elif [[ "$ARCH" == arm* ]]; then
  ARCH="arm64";
else
  echo "Error: Unsupported arch $ARCH";
  exit 1;
fi

URL="https://github.com/enuesaa/cywagon/releases/download/v${VERSION_NUMBER}/cywagon_${VERSION_NUMBER}_${OS}_${ARCH}.tar.gz";
TMP_OUT_DIR="cywagon_${VERSION_NUMBER}_${OS}_${ARCH}";

echo "Downloading $URL";
echo "";

if [ -d "$TMP_OUT_DIR" ]; then
  echo "Error: Installation stopped because the directory ${TMP_OUT_DIR} already exists";
  exit 1;
fi

mkdir $TMP_OUT_DIR;

# download
curl -L "$URL" | tar -xz -C $TMP_OUT_DIR;
mv $TMP_OUT_DIR/cywagon .;

rm -rf $TMP_OUT_DIR;

echo "";
echo "Download complete!";
echo "";
echo "Next, run the following command:"
echo "  mv cywagon /usr/local/bin/cywagon";
