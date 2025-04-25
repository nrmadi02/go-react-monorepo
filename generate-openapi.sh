#!/bin/bash
# Script untuk generate openapi types dari backend sesuai env

if [ -z "$VITE_BACKEND_URL" ]; then
  echo "VITE_BACKEND_URL belum diset di environment."
  exit 1
fi

npx openapi-typescript "$VITE_BACKEND_URL/openapi/swagger.yaml" -o ./common/schema/openapi.d.ts
