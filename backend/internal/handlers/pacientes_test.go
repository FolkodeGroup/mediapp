package handlers

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func makeCtx(method, path string, body []byte) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(method, path, bytes.NewBuffer(body))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	c.Request = req
	return c, w
}

func TestCreatePaciente_BindError(t *testing.T) {
	c, w := makeCtx("POST", "/api/v1/pacientes", []byte("{notjson}"))
	h := NewPacienteHandler(&mockPoolP{}, zap.NewNop())
	h.CreatePaciente(c)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 got %d body=%s", w.Code, w.Body.String())
	}
}

func TestInspectTables_NotAllowed(t *testing.T) {
	c, w := makeCtx("GET", "/api/v1/inspect/tables?table=hack", nil)
	q := c.Request.URL.Query()
	q.Add("table", "hack")
	c.Request.URL.RawQuery = q.Encode()
	h := NewPacienteHandler(&mockPoolP{}, zap.NewNop())
	h.InspectTables(c)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 got %d body=%s", w.Code, w.Body.String())
	}
}

func TestGetPacientes_NoPanic(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// provide a mock pool that returns an error on Query to avoid nil deref
	pool := &mockPoolP{queryFunc: func(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
		return nil, errors.New("no pool")
	}}
	h := NewPacienteHandler(pool, zap.NewNop())
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/pacientes", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	h.GetPacientes(ctx)
	// test passes if handler doesn't panic
}
