package security

import "golang.org/x/crypto/bcrypt"

// HashPassword hashea una contraseña usando bcrypt con el costo por defecto
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// CheckPasswordHash verifica si una contraseña coincide con su hash
// Retorna true si coinciden, false en caso contrario
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// CheckPassword mantiene la interfaz anterior para compatibilidad
// Retorna nil si coinciden, error si no
func CheckPassword(hash, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}