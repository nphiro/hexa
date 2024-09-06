#!/bin/bash

set -e

cert_pem=$(openssl req -x509 -newkey rsa:2048 -nodes -sha256 -subj '/CN=localhost' -keyout >(cat) -out >(cat) 2>/dev/null)

cert=$(echo "$cert_pem" | sed -n '/-----BEGIN CERTIFICATE-----/,/-----END CERTIFICATE-----/p')
key=$(echo "$cert_pem" | sed -n '/-----BEGIN PRIVATE KEY-----/,/-----END PRIVATE KEY-----/p')

cert_pem_base64=$(echo "$cert" | base64 | tr -d '\n')
key_pem_base64=$(echo "$key" | base64 | tr -d '\n')

if grep -q "^TLS_CERT=" .env; then
  sed -i '' "s|^TLS_CERT=.*|TLS_CERT=$cert_pem_base64|" .env
else
  echo "TLS_CERT=$cert_pem_base64" >> .env
fi

if grep -q "^TLS_KEY=" .env; then
  sed -i '' "s|^TLS_KEY=.*|TLS_KEY=$key_pem_base64|" .env
else
  echo "TLS_KEY=$key_pem_base64" >> .env
fi
