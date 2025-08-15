
package auth

import (
    "fmt"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "go.uber.org/zap"
)

// getJWTSecretKey obtiene la clave secreta de las variables de entorno
func getJWTSecretKey(logger *zap.Logger) []byte {
    secret := os.Getenv("JWT_SECRET_KEY")
    if secret == "" {
        logger.Warn("JWT_SECRET_KEY no encontrada, usando valor por defecto")
        return []byte("clave_secreta_por_defecto_muy_segura_12345")
    }
    if len(secret) < 32 {
        logger.Warn("JWT_SECRET_KEY debería tener al menos 32 caracteres para mayor seguridad")
    }
    return []byte(secret)
}

// Variable para la clave secreta
var jwtSecretKey []byte

// Inicializar el paquete auth
func Init(logger *zap.Logger) {
    jwtSecretKey = getJWTSecretKey(logger)
}

// CustomClaims estructura que incluye claims personalizados y estándar
type CustomClaims struct {
    UserID   string `json:"user_id"`
    UserRole string `json:"user_role"`
    jwt.RegisteredClaims
}

// GenerateToken crea y firma un nuevo token JWT
func GenerateToken(userID string, userRole string) (string, error) {
    if jwtSecretKey == nil {
        return "", fmt.Errorf("JWT no inicializado. Llama a auth.Init() primero")
    }

    expirationTime := time.Now().Add(24 * time.Hour)

    claims := &CustomClaims{
        UserID:   userID,
        UserRole: userRole,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
            Subject:   fmt.Sprintf("user:%s", userID),
            Issuer:    "mediapp-backend",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString(jwtSecretKey)
    if err != nil {
        return "", fmt.Errorf("error al firmar el token: %w", err)
    }

    return signedToken, nil
}

// ValidateToken valida un token JWT
func ValidateToken(tokenString string) (*CustomClaims, error) {
    if jwtSecretKey == nil {
        return nil, fmt.Errorf("JWT no inicializado. Llama a auth.Init() primero")
    }

    claims := &CustomClaims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
        }
        return jwtSecretKey, nil
    })

    if err != nil {
        return nil, fmt.Errorf("error al parsear el token: %w", err)
    }

    if !token.Valid {
        return nil, fmt.Errorf("token inválido")
    }

    return claims, nil
}