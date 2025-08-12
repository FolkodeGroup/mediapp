-- +goose Up
CREATE TABLE patients (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    dni VARCHAR(20) NOT NULL,
    medical_record_id VARCHAR(50) NOT NULL,
    birth_date DATE NOT NULL,
    gender CHAR(1) NOT NULL,
    email VARCHAR(255) NOT NULL
);

CREATE UNIQUE INDEX idx_patients_dni ON patients(dni);
CREATE UNIQUE INDEX idx_patients_medical_record_id ON patients(medical_record_id);

-- +goose Down
DROP TABLE IF EXISTS patients;
