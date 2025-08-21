-- +goose Up
-- Eliminar tablas viejas si existen

-- Eliminar todas las tablas posibles (incluyendo nombres viejos y nuevos)
DROP TABLE IF EXISTS auditorias CASCADE;
DROP TABLE IF EXISTS consultorios CASCADE;
DROP TABLE IF EXISTS datos_personales CASCADE;
DROP TABLE IF EXISTS historia_clinica_version CASCADE;
DROP TABLE IF EXISTS historias_clinicas CASCADE;
DROP TABLE IF EXISTS pacientes CASCADE;
DROP TABLE IF EXISTS permisos CASCADE;
DROP TABLE IF EXISTS recetas_medicas CASCADE;
DROP TABLE IF EXISTS rol_permiso CASCADE;
DROP TABLE IF EXISTS roles CASCADE;
DROP TABLE IF EXISTS turnos CASCADE;
DROP TABLE IF EXISTS usuarios CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS patients CASCADE;

-- Crear tablas nuevas (copiadas de las migraciones normalizadas)


-- ORDEN CORRECTO DE CREACIÓN DE TABLAS

-- 1. roles
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    nombre_rol VARCHAR(50) NOT NULL UNIQUE
);

-- 2. permisos
CREATE TABLE IF NOT EXISTS permisos (
    id SERIAL PRIMARY KEY,
    nombre_permiso VARCHAR(100) NOT NULL UNIQUE
);

-- 3. consultorios
CREATE TABLE IF NOT EXISTS consultorios (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    direccion TEXT NOT NULL
);

-- 4. usuarios
CREATE TABLE IF NOT EXISTS usuarios (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    contrasena_hash TEXT NOT NULL,
    rol_id INT NOT NULL REFERENCES roles(id),
    consultorio_id UUID REFERENCES consultorios(id),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    creado_en TIMESTAMP NOT NULL DEFAULT NOW()
);

-- 5. pacientes
CREATE TABLE IF NOT EXISTS pacientes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    fecha_nacimiento DATE NOT NULL,
    nro_credencial VARCHAR(50),
    obra_social VARCHAR(100),
    condicion_iva VARCHAR(50),
    plan VARCHAR(100),
    creado_por_usuario UUID REFERENCES usuarios(id),
    consultorio_id UUID REFERENCES consultorios(id),
    creado_en TIMESTAMP NOT NULL DEFAULT NOW()
);

-- 6. rol_permiso
CREATE TABLE IF NOT EXISTS rol_permiso (
    rol_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    permiso_id INT NOT NULL REFERENCES permisos(id) ON DELETE CASCADE,
    PRIMARY KEY (rol_id, permiso_id)
);

-- 7. historias_clinicas
CREATE TABLE IF NOT EXISTS historias_clinicas (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    paciente_id UUID NOT NULL REFERENCES pacientes(id) ON DELETE CASCADE,
    usuario_id UUID NOT NULL REFERENCES usuarios(id),
    fecha_consulta TIMESTAMP NOT NULL
);

-- 8. historia_clinica_version
CREATE TABLE IF NOT EXISTS historia_clinica_version (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    historia_clinica_id UUID NOT NULL REFERENCES historias_clinicas(id) ON DELETE CASCADE,
    motivo_consulta TEXT,
    antecedentes TEXT,
    examen_fisico TEXT,
    diagnostico TEXT,
    tratamiento TEXT,
    usuario_id UUID NOT NULL REFERENCES usuarios(id),
    modificado_en TIMESTAMP NOT NULL DEFAULT NOW()
);

-- 9. datos_personales
CREATE TABLE IF NOT EXISTS datos_personales (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    paciente_id UUID NOT NULL REFERENCES pacientes(id) ON DELETE CASCADE,
    telefono_encriptado BYTEA,
    dni_encriptado BYTEA,
    direccion TEXT
);

-- 10. recetas_medicas
CREATE TABLE IF NOT EXISTS recetas_medicas (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    paciente_id UUID NOT NULL REFERENCES pacientes(id) ON DELETE CASCADE,
    usuario_id UUID NOT NULL REFERENCES usuarios(id),
    contenido TEXT,
    fecha_emision TIMESTAMP NOT NULL DEFAULT NOW(),
    firma_digital BOOLEAN DEFAULT FALSE
);

-- 11. turnos
CREATE TABLE IF NOT EXISTS turnos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    paciente_id UUID NOT NULL REFERENCES pacientes(id) ON DELETE CASCADE,
    usuario_id UUID NOT NULL REFERENCES usuarios(id),
    fecha TIMESTAMP NOT NULL,
    motivo TEXT
);

-- 12. auditorias
CREATE TABLE IF NOT EXISTS auditorias (
    id SERIAL PRIMARY KEY,
    usuario_id UUID NOT NULL REFERENCES usuarios(id),
    accion VARCHAR(100) NOT NULL,
    tabla_afectada VARCHAR(100) NOT NULL,
    fecha TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP TABLE IF EXISTS auditorias CASCADE;
DROP TABLE IF EXISTS consultorios CASCADE;
DROP TABLE IF EXISTS datos_personales CASCADE;
DROP TABLE IF EXISTS historia_clinica_version CASCADE;
DROP TABLE IF EXISTS historias_clinicas CASCADE;
DROP TABLE IF EXISTS pacientes CASCADE;
DROP TABLE IF EXISTS permisos CASCADE;
DROP TABLE IF EXISTS recetas_medicas CASCADE;
DROP TABLE IF EXISTS rol_permiso CASCADE;
DROP TABLE IF EXISTS roles CASCADE;
DROP TABLE IF EXISTS turnos CASCADE;
DROP TABLE IF EXISTS usuarios CASCADE;

