#!/bin/bash

CONTAINER=openapitools/openapi-generator-cli:latest

# force docker usage
CONTAINER_TOOL=docker

# For now, we skip the dateTime parsing because of incompatible formats ( inventory does not produce fully compliant

function generate_client() {
  NAME=$1
  SPEC_URL=$2

  LOCAL_JSON="$NAME/api/openapi.json"
  curl $SPEC_URL -o "$LOCAL_JSON"

  HERE=$(pwd)
  # remove old code and doc
  find "$HERE/$NAME" -name '*.go' -o -name '*.md' -print0 | xargs -0 rm

  $CONTAINER_TOOL run --rm -v ${HERE}:/local:z --security-opt=label=disable $CONTAINER generate \
      -i "/local/$LOCAL_JSON" \
      -g go \
      --api-package $NAME \
      -p packageName=$NAME,isGoSubmodule=true \
      --git-host "github.com" --git-user-id RedHatInsights --git-repo-id patchman-clients \
      --type-mappings DateTime=string \
      -o /local/$NAME
}


generate_client inventory "https://ci.cloud.redhat.com/api/inventory/v1/openapi.json"
generate_client vmaas "https://webapp-vmaas-ci.5a9f.insights-dev.openshiftapps.com/api/openapi.json"
generate_client rbac "https://ci.cloud.redhat.com/api/rbac/v1/openapi.json"
