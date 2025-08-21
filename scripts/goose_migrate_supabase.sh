#!/bin/bash
# Script para exportar el PATH y ejecutar migraciones Goose en Supabase

# Asegura que $HOME/go/bin esté en el PATH
export PATH=$PATH:$HOME/go/bin

# String de conexión a Supabase (ajusta si cambian tus credenciales)
export SUPABASE_DB_URL="postgres://postgres.omtkjcrkvwwjpownndvh:MediApp2025-@aws-1-us-east-2.pooler.supabase.com:5432/postgres"

# Ejecutar migraciones
cd "$(dirname "$0")/.."
goose -dir ./backend/migrations postgres "$SUPABASE_DB_URL" up
