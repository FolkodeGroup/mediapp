package handlers

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/FolkodeGroup/mediapp/internal/auth"
	"github.com/FolkodeGroup/mediapp/internal/security"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
)

// setDest asigna el valor v al destino dest[i], donde dest[i] es un puntero
// similar a lo que espera database/sql.Scan. Soporta los tipos usados en los tests.
func setDest(dest []interface{}, i int, v interface{}) {
	switch p := dest[i].(type) {
	case *uuid.UUID:
		*p = v.(uuid.UUID)
	case *string:
		*p = v.(string)
	case *int:
		switch vv := v.(type) {
		case int:
			*p = vv
		case int64:
			*p = int(vv)
		case int32:
			*p = int(vv)
		default:
			// no hacer nada
		}
	case *int64:
		switch vv := v.(type) {
		case int64:
			*p = vv
		case int:
			*p = int64(vv)
		case int32:
			*p = int64(vv)
		default:
			// no hacer nada
		}
	case *bool:
		*p = v.(bool)
	case *time.Time:
		*p = v.(time.Time)
	case **uuid.UUID:
		if v == nil {
			*p = nil
		} else {
			val := v.(uuid.UUID)
			*p = &val
		}
	case **time.Time:
		if v == nil {
			*p = nil
		} else {
			val := v.(time.Time)
			*p = &val
		}
	default:
		// Intentar asignación por reflexión como último recurso
		rv := reflect.ValueOf(dest[i])
		if rv.Kind() == reflect.Ptr && !rv.IsNil() {
			ev := rv.Elem()
			vv := reflect.ValueOf(v)
			if vv.IsValid() && vv.Type().ConvertibleTo(ev.Type()) {
				ev.Set(vv.Convert(ev.Type()))
			}
		}
	}
}

type mockRow struct {
	scanFunc func(dest ...interface{}) error
}

func (r mockRow) Scan(dest ...interface{}) error { return r.scanFunc(dest...) }

// Implementa la interfaz de pgx.Row
func (r mockRow) FieldDescriptions() []pgconn.FieldDescription { return nil }
func (r mockRow) Values() ([]interface{}, error)               { return nil, nil }
func (r mockRow) RawValues() [][]byte                          { return nil }
func (r mockRow) Err() error                                   { return nil }

type mockDB struct {
	queryRowFunc func(ctx context.Context, sql string, args ...interface{}) pgx.Row
	execFunc     func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}

func (m *mockDB) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return m.queryRowFunc(ctx, sql, args...)
}
func (m *mockDB) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return m.execFunc(ctx, sql, args...)
}

// TestLoginSuccess prueba el login exitoso
func TestLoginSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userID := uuid.New()
	rolID := 2
	password := "testpass123"
	hashedPassword, _ := security.HashPassword(password)

	mockdb := &mockDB{
		queryRowFunc: func(ctx context.Context, _sql string, args ...interface{}) pgx.Row {
			return mockRow{

				scanFunc: func(dest ...interface{}) error {
					setDest(dest, 0, userID)
					setDest(dest, 1, "Test User")
					setDest(dest, 2, "test@example.com")
					setDest(dest, 3, hashedPassword)
					setDest(dest, 4, rolID)
					setDest(dest, 5, nil)
					setDest(dest, 6, true)
					setDest(dest, 7, time.Now())
					setDest(dest, 8, int64(0))
					setDest(dest, 9, nil)
					return nil
				},
			}
		},
		execFunc: func(ctx context.Context, _sql string, args ...interface{}) (pgconn.CommandTag, error) {
			return pgconn.NewCommandTag("UPDATE 1"), nil
		},
	}

	logger := zap.NewNop()
	h := NewAuthHandler(logger, mockdb)
	h.generateToken = func(uid string, rid int) (string, error) { return "mocktoken", nil }
	// Inyectar verificador de contraseña para test
	h.verifyPassword = func(plain, hash string) bool { return plain == password }

	reqBody := map[string]string{"email": "test@example.com", "password": password}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.Login(ctx)

	if rec.Code != http.StatusOK {
		t.Fatalf("Se esperaba status 200, obtuvo %d", rec.Code)
	}
	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp["token"] != "mocktoken" {
		t.Errorf("Se esperaba token 'mocktoken', obtuvo %v", resp["token"])
	}
}

// TestLoginInvalidPassword prueba login con contraseña incorrecta
func TestLoginInvalidPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)
	userID := uuid.New()
	rolID := 2
	password := "testpass123"
	hashedPassword, _ := security.HashPassword(password)

	mockdb := &mockDB{
		queryRowFunc: func(ctx context.Context, _sql string, args ...interface{}) pgx.Row {
			return mockRow{
				scanFunc: func(dest ...interface{}) error {
					setDest(dest, 0, userID)
					setDest(dest, 1, "Test User")
					setDest(dest, 2, "test@example.com")
					setDest(dest, 3, hashedPassword)
					setDest(dest, 4, rolID)
					setDest(dest, 5, nil)
					setDest(dest, 6, true)
					setDest(dest, 7, time.Now())
					setDest(dest, 8, int64(0))
					setDest(dest, 9, nil)
					return nil
				},
			}
		},
		execFunc: func(ctx context.Context, _sql string, args ...interface{}) (pgconn.CommandTag, error) {
			return pgconn.NewCommandTag("UPDATE 1"), nil
		},
	}

	logger := zap.NewNop()
	h := NewAuthHandler(logger, mockdb)
	h.generateToken = func(uid string, rid int) (string, error) { return "mocktoken", nil }
	h.verifyPassword = func(plain, hash string) bool { return false } // forzar fallo

	reqBody := map[string]string{"email": "test@example.com", "password": "wrongpass"}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.Login(ctx)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("Se esperaba status 401, obtuvo %d", rec.Code)
	}
	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp["error"] == nil {
		t.Errorf("Se esperaba mensaje de error, obtuvo %v", resp)
	}
}

// TestLoginUserBlocked prueba login con usuario bloqueado
func TestLoginUserBlocked(t *testing.T) {
	gin.SetMode(gin.TestMode)
	userID := uuid.New()
	rolID := 2
	password := "testpass123"
	hashedPassword, _ := security.HashPassword(password)

	mockdb := &mockDB{
		queryRowFunc: func(ctx context.Context, _sql string, args ...interface{}) pgx.Row {
			return mockRow{
				scanFunc: func(dest ...interface{}) error {
					setDest(dest, 0, userID)
					setDest(dest, 1, "Test User")
					setDest(dest, 2, "test@example.com")
					setDest(dest, 3, hashedPassword)
					setDest(dest, 4, rolID)
					setDest(dest, 5, nil)
					setDest(dest, 6, true)
					setDest(dest, 7, time.Now())
					setDest(dest, 8, int64(6))
					setDest(dest, 9, nil)
					return nil
				},
			}
		},
		execFunc: func(ctx context.Context, _sql string, args ...interface{}) (pgconn.CommandTag, error) {
			return pgconn.NewCommandTag("UPDATE 1"), nil
		},
	}

	logger := zap.NewNop()
	h := NewAuthHandler(logger, mockdb)
	h.generateToken = func(uid string, rid int) (string, error) { return "mocktoken", nil }
	h.verifyPassword = func(plain, hash string) bool { return plain == password }

	reqBody := map[string]string{"email": "test@example.com", "password": password}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.Login(ctx)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("Se esperaba status 401, obtuvo %d", rec.Code)
	}
	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	// Solo debe contener el mensaje de bloqueo, no intentos_restantes
	if resp["error"] == nil {
		t.Errorf("Se esperaba mensaje de bloqueo, obtuvo %v", resp)
	}
	if _, ok := resp["intentos_restantes"]; ok {
		t.Errorf("No se esperaba 'intentos_restantes' en la respuesta de usuario bloqueado: %v", resp)
	}
}

// TestLoginUserNotFound prueba login con usuario no existente
func TestLoginUserNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockdb := &mockDB{
		queryRowFunc: func(ctx context.Context, _sql string, args ...interface{}) pgx.Row {
			return mockRow{
				scanFunc: func(dest ...interface{}) error {
					// Simula sql.ErrNoRows exactamente
					return sql.ErrNoRows
				},
			}
		},
		execFunc: func(ctx context.Context, _sql string, args ...interface{}) (pgconn.CommandTag, error) {
			return pgconn.NewCommandTag("UPDATE 0"), nil
		},
	}

	logger := zap.NewNop()
	h := NewAuthHandler(logger, mockdb)
	h.generateToken = func(uid string, rid int) (string, error) { return "mocktoken", nil }
	h.verifyPassword = func(plain, hash string) bool { return plain == "irrelevante" }

	reqBody := map[string]string{"email": "noexiste@example.com", "password": "irrelevante"}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.Login(ctx)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("Se esperaba status 401, obtuvo %d", rec.Code)
	}
	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp["error"] == nil {
		t.Errorf("Se esperaba mensaje de error, obtuvo %v", resp)
	}
}

// TestLoginQueryError simula un error inesperado de la consulta y espera 500
func TestLoginQueryError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockdb := &mockDB{
		queryRowFunc: func(ctx context.Context, _sql string, args ...interface{}) pgx.Row {
			return mockRow{scanFunc: func(dest ...interface{}) error {
				return errors.New("db query failed")
			}}
		},
		execFunc: func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
			return pgconn.NewCommandTag("UPDATE 0"), nil
		},
	}

	logger := zap.NewNop()
	h := NewAuthHandler(logger, mockdb)
	h.generateToken = func(uid string, rid int) (string, error) { return "", nil }
	h.verifyPassword = func(plain, hash string) bool { return false }

	reqBody := map[string]string{"email": "x@example.com", "password": "p"}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.Login(ctx)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("Se esperaba status 500, obtuvo %d", rec.Code)
	}
}

// TestLoginBindError prueba que un body inválido produce 400
func TestLoginBindError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	logger := zap.NewNop()
	h := NewAuthHandler(logger, nil)

	// Body sin password
	reqBody := map[string]string{"email": "test@example.com"}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.Login(ctx)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("Se esperaba status 400, obtuvo %d", rec.Code)
	}
}

// TestLoginTokenGenerationError fuerza fallo en la generación de token y espera 500
func TestLoginTokenGenerationError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userID := uuid.New()
	rolID := 2
	password := "testpass123"
	hashedPassword, _ := security.HashPassword(password)

	mockdb := &mockDB{
		queryRowFunc: func(ctx context.Context, _sql string, args ...interface{}) pgx.Row {
			return mockRow{scanFunc: func(dest ...interface{}) error {
				setDest(dest, 0, userID)
				setDest(dest, 1, "Test User")
				setDest(dest, 2, "test@example.com")
				setDest(dest, 3, hashedPassword)
				setDest(dest, 4, rolID)
				setDest(dest, 5, nil)
				setDest(dest, 6, true)
				setDest(dest, 7, time.Now())
				setDest(dest, 8, int64(0))
				setDest(dest, 9, nil)
				return nil
			}}
		},
		execFunc: func(ctx context.Context, _sql string, args ...interface{}) (pgconn.CommandTag, error) {
			return pgconn.NewCommandTag("UPDATE 1"), nil
		},
	}

	logger := zap.NewNop()
	h := NewAuthHandler(logger, mockdb)
	h.generateToken = func(uid string, rid int) (string, error) { return "", errors.New("token fail") }
	h.verifyPassword = func(plain, hash string) bool { return plain == password }

	reqBody := map[string]string{"email": "test@example.com", "password": password}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.Login(ctx)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("Se esperaba status 500, obtuvo %d", rec.Code)
	}
}

// TestLoginExecUpdateAttemptsError simula error en Exec al incrementar intentos (debe devolver 401 y no 500)
func TestLoginExecUpdateAttemptsError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userID := uuid.New()
	rolID := 2
	password := "testpass123"
	hashedPassword, _ := security.HashPassword(password)

	mockdb := &mockDB{
		queryRowFunc: func(ctx context.Context, _sql string, args ...interface{}) pgx.Row {
			return mockRow{scanFunc: func(dest ...interface{}) error {
				setDest(dest, 0, userID)
				setDest(dest, 1, "Test User")
				setDest(dest, 2, "test@example.com")
				setDest(dest, 3, hashedPassword)
				setDest(dest, 4, rolID)
				setDest(dest, 5, nil)
				setDest(dest, 6, true)
				setDest(dest, 7, time.Now())
				setDest(dest, 8, int64(0))
				setDest(dest, 9, nil)
				return nil
			}}
		},
		execFunc: func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
			// Simular fallo en cualquier UPDATE
			if strings.Contains(sql, "UPDATE usuarios") {
				return pgconn.NewCommandTag("UPDATE 0"), errors.New("exec failed")
			}
			return pgconn.NewCommandTag("OK"), nil
		},
	}

	logger := zap.NewNop()
	h := NewAuthHandler(logger, mockdb)
	h.generateToken = func(uid string, rid int) (string, error) { return "", nil }
	h.verifyPassword = func(plain, hash string) bool { return false }

	reqBody := map[string]string{"email": "test@example.com", "password": "wrong"}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.Login(ctx)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("Se esperaba status 401, obtuvo %d", rec.Code)
	}
}

// TestLoginResetAttemptsExecError simula fallo al resetear intentos tras login exitoso; debe devolver 200
func TestLoginResetAttemptsExecError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userID := uuid.New()
	rolID := 2
	password := "testpass123"
	hashedPassword, _ := security.HashPassword(password)

	mockdb := &mockDB{
		queryRowFunc: func(ctx context.Context, sql string, args ...interface{}) pgx.Row {
			return mockRow{scanFunc: func(dest ...interface{}) error {
				setDest(dest, 0, userID)
				setDest(dest, 1, "Test User")
				setDest(dest, 2, "test@example.com")
				setDest(dest, 3, hashedPassword)
				setDest(dest, 4, rolID)
				setDest(dest, 5, nil)
				setDest(dest, 6, true)
				setDest(dest, 7, time.Now())
				setDest(dest, 8, int64(0))
				setDest(dest, 9, nil)
				return nil
			}}
		},
		execFunc: func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
			// Error intencional al resetear intentos (contiene "SET intentos_fallidos = 0")
			if strings.Contains(sql, "intentos_fallidos = 0") {
				return pgconn.NewCommandTag("UPDATE 0"), errors.New("reset failed")
			}
			return pgconn.NewCommandTag("UPDATE 1"), nil
		},
	}

	logger := zap.NewNop()
	h := NewAuthHandler(logger, mockdb)
	h.generateToken = func(uid string, rid int) (string, error) { return "mocktoken", nil }
	h.verifyPassword = func(plain, hash string) bool { return plain == password }

	reqBody := map[string]string{"email": "test@example.com", "password": password}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.Login(ctx)

	if rec.Code != http.StatusOK {
		t.Fatalf("Se esperaba status 200, obtuvo %d", rec.Code)
	}
}

// TestRegisterSuccess prueba registro exitoso
func TestRegisterSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	consultorio := uuid.New()
	mockdb := &mockDB{
		queryRowFunc: func(ctx context.Context, _sql string, args ...interface{}) pgx.Row {
			return mockRow{scanFunc: func(dest ...interface{}) error { return sql.ErrNoRows }}
		},
		execFunc: func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
			return pgconn.NewCommandTag("INSERT 1"), nil
		},
	}

	logger := zap.NewNop()
	h := NewAuthHandler(logger, mockdb)

	reqBody := map[string]interface{}{
		"nombre":         "Nuevo",
		"email":          "nuevo@example.com",
		"password":       "password",
		"rol_id":         1,
		"consultorio_id": consultorio.String(),
		"activo":         true,
	}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.Register(ctx)

	if rec.Code != http.StatusCreated {
		t.Fatalf("Se esperaba status 201, obtuvo %d", rec.Code)
	}
}

// TestRegisterInvalidConsultorioUUID valida que request con uuid inválido devuelve 400
func TestRegisterInvalidConsultorioUUID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	logger := zap.NewNop()
	h := NewAuthHandler(logger, nil)

	reqBody := map[string]interface{}{
		"nombre":         "Nuevo",
		"email":          "nuevo@example.com",
		"password":       "password",
		"rol_id":         1,
		"consultorio_id": "not-uuid",
		"activo":         true,
	}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.Register(ctx)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("Se esperaba status 400, obtuvo %d", rec.Code)
	}
}

// ProtectedEndpoint tests
func TestProtectedEndpointMissingToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := zap.NewNop()
	h := NewAuthHandler(logger, nil)

	req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.ProtectedEndpoint(ctx)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("Se esperaba status 401, obtuvo %d", rec.Code)
	}
}

func TestProtectedEndpointInvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := zap.NewNop()
	h := NewAuthHandler(logger, nil)

	req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.ProtectedEndpoint(ctx)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("Se esperaba status 401, obtuvo %d", rec.Code)
	}
}

func TestProtectedEndpointValidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := zap.NewNop()
	auth.Init(logger)

	h := NewAuthHandler(logger, nil)

	uid := uuid.New()
	token, err := auth.GenerateToken(uid.String(), 2)
	if err != nil {
		t.Fatalf("No se pudo generar token en test: %v", err)
	}

	req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.ProtectedEndpoint(ctx)

	if rec.Code != http.StatusOK {
		t.Fatalf("Se esperaba status 200, obtuvo %d", rec.Code)
	}

	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp["user_id"] != uid.String() {
		t.Errorf("user_id esperado %s, obtenido %v", uid.String(), resp["user_id"])
	}
}
