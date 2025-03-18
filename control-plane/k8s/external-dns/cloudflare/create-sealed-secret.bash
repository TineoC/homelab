#!/bin/bash

kubectl create secret generic \
  -n kube-system \
  cloudflare-api-key \
  --from-literal=apiKey=$CLOUDFLARE_API_TOKEN \
  --from-literal=email=$CLOUDFLARE_EMAIL \
  --dry-run=client \
  -o yaml | \
  kubeseal \
  --controller-namespace kube-system \
  --format yaml > sealed-cloudflare-secret.yaml
kubectl apply -f sealed-cloudflare-secret.yaml -n kube-system