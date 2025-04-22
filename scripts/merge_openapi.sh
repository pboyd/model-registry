#!/bin/bash

set -e

cd "$(pwd)/$(dirname "$0")/.."
YQ=${YQ:-$(realpath -e "bin/yq")}

CHECK=false
while [[ $# -gt 0 ]]; do
    case "$1" in
        --check)
            CHECK=true
            shift
            ;;
        -h|--help)
            echo "Usage: $0 [--check]"
            echo "  --check: Check for differences in the generated OpenAPI spec."
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

OUT_FILE=api/openapi/model-registry.yaml
if [[ "$CHECK" == "true" ]]; then
    OUT_FILE=$(mktemp --suffix=.yaml)
    trap 'rm "$OUT_FILE"' EXIT
fi

# Merge the src files together.
$YQ eval-all '. as $item ireduce ({}; . * $item )' api/openapi/src/*.yaml >"$OUT_FILE"

# Re-order the keys in the generated file.
$YQ eval -i '
    {
        "openapi": .openapi,
        "info": .info,
        "servers": .servers,
        "paths": .paths,
        "components": .components,
        "security": .security,
        "tags": .tags
    } |
        sort_keys(.paths) |
        sort_keys(.components.schemas) |
        sort_keys(.components.responses)
' "$OUT_FILE"

if [[ "$CHECK" == "true" ]]; then
    exec diff -u api/openapi/model-registry.yaml $OUT_FILE
fi
