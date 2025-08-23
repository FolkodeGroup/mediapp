package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// mockPoolP implements PoolTX for pacientes tests
type mockPoolP struct {
	execFunc     func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	queryFunc    func(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	queryRowFunc func(ctx context.Context, sql string, args ...interface{}) pgx.Row
	pingFunc     func(ctx context.Context) error
}

func (m *mockPoolP) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	if m.execFunc != nil {
		return m.execFunc(ctx, sql, args...)
	}
	return pgconn.NewCommandTag(""), nil
}
func (m *mockPoolP) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	if m.queryFunc != nil {
		return m.queryFunc(ctx, sql, args...)
	}
	return nil, nil
}
func (m *mockPoolP) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	if m.queryRowFunc != nil {
		return m.queryRowFunc(ctx, sql, args...)
	}
	return mockRowP{scanFunc: func(dest ...interface{}) error { return pgx.ErrNoRows }}
}
func (m *mockPoolP) Ping(ctx context.Context) error {
	if m.pingFunc != nil {
		return m.pingFunc(ctx)
	}
	return nil
}
func (m *mockPoolP) Stat() *pgxpool.Stat { return &pgxpool.Stat{} }

// mockRowsP implements pgx.Rows for our tests
type mockRowsP struct {
	idx       int
	rowsCount int
	scanFunc  func(dest ...interface{}) error
}

func (r *mockRowsP) Next() bool {
	if r.idx < r.rowsCount {
		r.idx++
		return true
	}
	return false
}
func (r *mockRowsP) Scan(dest ...interface{}) error {
	if r.scanFunc != nil {
		return r.scanFunc(dest...)
	}
	return nil
}
func (r *mockRowsP) Close()                                       {}
func (r *mockRowsP) Err() error                                   { return nil }
func (r *mockRowsP) FieldDescriptions() []pgconn.FieldDescription { return nil }
func (r *mockRowsP) Values() ([]interface{}, error)               { return nil, nil }
func (r *mockRowsP) RawValues() [][]byte                          { return nil }
func (r *mockRowsP) CommandTag() pgconn.CommandTag                { return pgconn.NewCommandTag("SELECT 1") }
func (r *mockRowsP) Conn() *pgx.Conn                              { return nil }

type mockRowP struct {
	scanFunc func(dest ...interface{}) error
}

func (r mockRowP) Scan(dest ...interface{}) error { return r.scanFunc(dest...) }

// TestGetPacientes_Success verifies GetPacientes returns list
func TestGetPacientes_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	now := time.Now()
	rows := &mockRowsP{rowsCount: 1}
	rows.scanFunc = func(dest ...interface{}) error {
		if p, ok := dest[0].(*[16]byte); ok {
			u := uuid.New()
			copy(p[:], u[:])
		}
		if p, ok := dest[1].(*string); ok {
			*p = "Nombre"
		}
		if p, ok := dest[2].(*string); ok {
			*p = "Apellido"
		}
		if p, ok := dest[3].(*time.Time); ok {
			*p = now
		}
		if p, ok := dest[4].(**string); ok {
			val := "123"
			*p = &val
		}
		if p, ok := dest[5].(**string); ok {
			*p = nil
		}
		if p, ok := dest[6].(**string); ok {
			*p = nil
		}
		if p, ok := dest[7].(**string); ok {
			*p = nil
		}
		if p, ok := dest[8].(*[16]byte); ok {
			u := uuid.New()
			copy(p[:], u[:])
		}
		if p, ok := dest[9].(*[16]byte); ok {
			u := uuid.New()
			copy(p[:], u[:])
		}
		if p, ok := dest[10].(*time.Time); ok {
			*p = now
		}
		return nil
	}

	pool := &mockPoolP{queryFunc: func(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
		return rows, nil
	}}

	logger := zap.NewNop()
	h := NewPacienteHandler(pool, logger)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/pacientes", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.GetPacientes(ctx)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp["total"].(float64) != 1 {
		t.Fatalf("expected total 1, got %v", resp["total"])
	}
}

// TestGetPaciente_Success verifies GetPaciente by id
func TestGetPaciente_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	now := time.Now()
	row := mockRowP{scanFunc: func(dest ...interface{}) error {
		if p, ok := dest[0].(*[16]byte); ok {
			u := uuid.New()
			copy(p[:], u[:])
		}
		if p, ok := dest[1].(*string); ok {
			*p = "Nombre"
		}
		if p, ok := dest[2].(*string); ok {
			*p = "Apellido"
		}
		if p, ok := dest[3].(*time.Time); ok {
			*p = now
		}
		if p, ok := dest[4].(**string); ok {
			val := "123"
			*p = &val
		}
		if p, ok := dest[5].(**string); ok {
			*p = nil
		}
		if p, ok := dest[6].(**string); ok {
			*p = nil
		}
		if p, ok := dest[7].(**string); ok {
			*p = nil
		}
		if p, ok := dest[8].(*[16]byte); ok {
			u := uuid.New()
			copy(p[:], u[:])
		}
		if p, ok := dest[9].(*[16]byte); ok {
			u := uuid.New()
			copy(p[:], u[:])
		}
		if p, ok := dest[10].(*time.Time); ok {
			*p = now
		}
		return nil
	}}

	pool := &mockPoolP{queryRowFunc: func(ctx context.Context, sql string, args ...interface{}) pgx.Row { return row }}
	logger := zap.NewNop()
	h := NewPacienteHandler(pool, logger)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/pacientes/1", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.GetPaciente(ctx)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

// TestCreateUpdateDeletePaciente flow
func TestCreateUpdateDeletePacienteFlow(t *testing.T) {
	gin.SetMode(gin.TestMode)
	pool := &mockPoolP{
		execFunc: func(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
			if strings.HasPrefix(strings.TrimSpace(strings.ToUpper(sql)), "INSERT") {
				return pgconn.NewCommandTag("INSERT 1"), nil
			}
			return pgconn.NewCommandTag("UPDATE 1"), nil
		},
	}
	logger := zap.NewNop()
	h := NewPacienteHandler(pool, logger)

	// Create
	reqBody := map[string]string{"nombre": "X", "apellido": "Y", "fecha_nacimiento": "2000-01-01", "consultorio_id": uuid.New().String()}
	b, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/pacientes", strings.NewReader(string(b)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	h.CreatePaciente(ctx)
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}

	// Update
	req2, _ := http.NewRequest(http.MethodPut, "/api/v1/pacientes/1", strings.NewReader(string(b)))
	req2.Header.Set("Content-Type", "application/json")
	rec2 := httptest.NewRecorder()
	ctx2, _ := gin.CreateTestContext(rec2)
	ctx2.Request = req2
	h.UpdatePaciente(ctx2)
	if rec2.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec2.Code)
	}

	// Delete
	req3, _ := http.NewRequest(http.MethodDelete, "/api/v1/pacientes/1", nil)
	rec3 := httptest.NewRecorder()
	ctx3, _ := gin.CreateTestContext(rec3)
	ctx3.Request = req3
	h.DeletePaciente(ctx3)
	if rec3.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec3.Code)
	}
}
