#!/bin/bash

set -e

private_key=$(openssl genpkey -algorithm Ed25519)
public_key=$(echo "$private_key" | openssl pkey -pubout)

private_key_base64=$(echo "$private_key" | base64 | tr -d '\n')
public_key_base64=$(echo "$public_key" | base64 | tr -d '\n')

if grep -q "^CRYPTO_PRIVATE_KEY=" .env; then
  sed -i '' "s|^CRYPTO_PRIVATE_KEY=.*|CRYPTO_PRIVATE_KEY=$private_key_base64|" .env
else
  echo "CRYPTO_PRIVATE_KEY=$private_key_base64" >> .env
fi

if grep -q "^CRYPTO_PUBLIC_KEY=" .env; then
  sed -i '' "s|^CRYPTO_PUBLIC_KEY=.*|CRYPTO_PUBLIC_KEY=$public_key_base64|" .env
else
  echo "CRYPTO_PUBLIC_KEY=$public_key_base64" >> .env
fi
