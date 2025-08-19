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

// CheckPassword verifica si una contraseña coincide con su hash
func CheckPassword(hash, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}