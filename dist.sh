#!/bin/bash
# build binary distributions for linux/amd64 and darwin/amd64
set -e

NAME="diter"
VERSION=1.0.0

TARGET_OS="windows linux darwin"
TARGET_ARCH="amd64"
DIST_DIR="./dist"

mkdir -p $DIST_DIR

dep ensure

for os in $TARGET_OS; do
  for arch in $TARGET_ARCH; do
    ext=
    if [ $os = windows ]; then
        ext=".exe"
    fi
    target="$DIST_DIR/$NAME-$VERSION.$os-$arch$ext"

    echo "building $target ..."
    GOOS=$os GOARCH=$arch \
        go build -ldflags="-s -w" -o $target || exit 1
  done
done
