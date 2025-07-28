#!/usr/bin/env bash
set -euo pipefail

if [[ -z "${PG_DSN:-}" ]]; then
  echo "PG_DSN not set"; exit 1
fi

atlas migrate apply --dir file://migrations --url "$PG_DSN"
