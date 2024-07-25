#!/usr/bin/env bash

set -e

VERSION=$1

SILICONCLOUD="./out/siliconcloud"
REPO=datengarden/siliconcloud-cli

go build -ldflags "-X 'github.com/siliconflow/siliconcloud-cli/meta.Version=${VERSION}'" -o $SILICONCLOUD main.go

if [[ "$VERSION" == *"-beta" ]]; then
  echo "beta. skip."
else
  echo "${VERSION}" > out/version

  RELEASE_ID=$(curl -fsSL \
    -H "Accept: application/vnd.github+json" \
    -H "Authorization: Bearer $GITHUB_TOKEN" \
    -H "X-GitHub-Api-Version: 2022-11-28" \
    https://api.github.com/repos/$REPO/releases/tags/v"$VERSION" | jq '.["id"]')

  DATA='{"draft":false,"prerelease":false,"make_latest":true}'

  curl -fsSL \
    -X PATCH \
    -H "Accept: application/vnd.github+json" \
    -H "Authorization: Bearer $GITHUB_TOKEN" \
    -H "X-GitHub-Api-Version: 2022-11-28" \
    https://api.github.com/repos/$REPO/releases/"$RELEASE_ID" \
    -d "$DATA"
fi