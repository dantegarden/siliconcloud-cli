#!/usr/bin/env bash

VERSION=$1
REPO=dantegarden/siliconcloud-cli
URL="https://api.github.com/repos/$REPO/releases"

LIST=(
    "siliconcloud-cli-macosx-$VERSION-amd64.tgz"
    "siliconcloud-cli-macosx-$VERSION-arm64.tgz"
    "siliconcloud-cli-linux-$VERSION-amd64.tgz"
    "siliconcloud-cli-linux-$VERSION-arm64.tgz"
    "siliconcloud-cli-windows-$VERSION-amd64.zip"
)

for filename in "${LIST[@]}"
do
    ASSET_ID=`curl -X GET -H "Content-Type: application/json" -H "Authorization: token $GITHUB_TOKEN" -H "Accept: application/vnd.github.v3.raw" $URL | jq ".[0].assets | map(select(.name == \"$filename\"))[0] | .id"`
    echo ASSET_ID=$ASSET_ID
    wget -q --auth-no-challenge --header='Accept:application/octet-stream' https://$GITHUB_TOKEN:@api.github.com/repos/$REPO/releases/assets/$ASSET_ID  -O $filename
    ls -l siliconcloud-cli*
    shasum -a 256 "$filename" >> SHASUMS256.txt
done

cat ./SHASUMS256.txt

