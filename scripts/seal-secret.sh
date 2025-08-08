#!/bin/bash
# scripts/seal-secret.sh

if [ $# -ne 2 ]; then
    echo "Usage: $0 <input-secret.yaml> <output-sealed-secret.yaml>"
    exit 1
fi

INPUT_FILE=$1
OUTPUT_FILE=$2

if [ ! -f "$INPUT_FILE" ]; then
    echo "Error: Input file $INPUT_FILE not found"
    exit 1
fi

CONTROLLER_NAME="sealed-secrets"
CONTROLLER_NAMESPACE="sealed-secrets"
FORMAT="yaml"

echo "Sealing secret from $INPUT_FILE to $OUTPUT_FILE"
kubeseal \
  --controller-name $CONTROLLER_NAME \
  --controller-namespace $CONTROLLER_NAMESPACE \
  --format $FORMAT \
  -f "$INPUT_FILE" \
  -w "$OUTPUT_FILE"

if [ $? -eq 0 ]; then
    echo "✅ Successfully sealed secret"
    echo "⚠️  Remember to delete the original secret file: $INPUT_FILE"
else
    echo "❌ Failed to seal secret"
    exit 1
fi