#!/bin/bash
# Script para obtener el UUID del usuario admin desde la API
API_URL="http://localhost:8080"

# Obtener el UUID del usuario admin por email
ADMIN_EMAIL="admin@example.com"

UUID=$(curl -s "$API_URL/api/v1/usuarios?email=$ADMIN_EMAIL" | jq -r '.usuarios[0].id')
echo "UUID del usuario admin: $UUID"
