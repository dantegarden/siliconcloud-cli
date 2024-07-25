#!/usr/bin/env bash

VERSION=$1
REPO=dantegarden/siliconcloud-cli
LIST=(
    "siliconcloud-cli-macosx-$VERSION-amd64.tgz"
    "siliconcloud-cli-macosx-$VERSION-arm64.tgz"
    "siliconcloud-cli-linux-$VERSION-amd64.tgz"
    "siliconcloud-cli-linux-$VERSION-arm64.tgz"
    "siliconcloud-cli-windows-$VERSION-amd64.zip"
)

for filename in "${LIST[@]}"
do
    curl -fsSL -O \
        -H "Authorization: Bearer $GITHUB_TOKEN" \
        https://github.com/$REPO/releases/download/v"$VERSION"/"$filename"
    ls -l
    shasum -a 256 "$filename" >> SHASUMS256.txt
done

cat ./SHASUMS256.txt

