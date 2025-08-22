#!/bin/bash
# Script de prueba para endpoints principales de mediapp
API_URL="http://localhost:8080"

# 1. Registrar usuario admin (ajusta el UUID de consultorio según tu base)
echo "Registrando usuario admin..."
curl -s -X POST "$API_URL/register" -H "Content-Type: application/json" -d '{
  "nombre": "Admin Test",
  "email": "admin@example.com",
  "password": "admin1234",
  "rol_id": 1,
  "consultorio_id": "7623abc7-5197-4f52-a9da-68594dffcf77",
  "activo": true
}' | jq

echo "\nLogueando usuario admin..."
# 2. Login usuario admin
TOKEN=$(curl -s -X POST "$API_URL/login" -H "Content-Type: application/json" -d '{
  "email": "admin@example.com",
  "password": "admin1234"
}' | jq -r .token)
echo "Token: $TOKEN"

# 3. Listar pacientes
echo "\nListando pacientes..."
curl -s -X GET "$API_URL/api/v1/pacientes" -H "Authorization: Bearer $TOKEN" | jq

# 4. Crear paciente de ejemplo (ajusta el UUID de consultorio y usuario según corresponda)
echo "\nCreando paciente de ejemplo..."
curl -s -X POST "$API_URL/api/v1/pacientes" -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d '{
  "nombre": "Juan",
  "apellido": "Pérez",
  "fecha_nacimiento": "1990-01-01",
  "nro_credencial": "123456",
  "obra_social": "OSDE",
  "condicion_iva": "Monotributo",
  "plan": "210",
  "creado_por_usuario": "REEMPLAZAR_UUID_USUARIO",
  "consultorio_id": "7623abc7-5197-4f52-a9da-68594dffcf77"
}' | jq

# 5. Listar pacientes nuevamente
echo "\nListando pacientes tras alta..."
curl -s -X GET "$API_URL/api/v1/pacientes" -H "Authorization: Bearer $TOKEN" | jq

echo "\nPruebas finalizadas."
