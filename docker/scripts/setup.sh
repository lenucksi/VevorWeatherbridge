#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

echo "=== Creating data directories ==="
mkdir -p /mnt/llama-docker/{models,agentmemory,llama-swap,llamacpp-cache}

echo "=== Starting llama.cpp stack ==="
docker compose -f compose.llamacpp.yaml up -d

echo ""
echo "=== Models are downloaded automatically on first start ==="
echo "    (--hf-repo fetches from HuggingFace)"
echo ""
echo "=== URLs ==="
echo "  llama-swap (API):     http://localhost:8080/v1"
echo "  Agentmemory viewer:   http://localhost:3113"
echo ""
echo "=== Daily workflow ==="
echo "  docker compose -f compose.llamacpp.yaml up -d     # start"
echo "  docker compose -f compose.llamacpp.yaml down      # stop"
  echo "  rsync -a /mnt/llama-docker/ backup/                 # backup"
echo ""
echo "=== Logs ==="
echo "  docker compose -f compose.llamacpp.yaml logs -f"
