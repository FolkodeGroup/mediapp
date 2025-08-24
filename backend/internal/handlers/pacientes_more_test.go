package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

// TestSupabaseConnection success path: ping ok and counts return 0
func TestTestSupabaseConnection_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	pool := &mockPoolP{
		pingFunc: func(ctx context.Context) error { return nil },
		queryRowFunc: func(ctx context.Context, sql string, args ...interface{}) pgx.Row {
			return mockRowP{scanFunc: func(dest ...interface{}) error {
				if len(dest) > 0 {
					if p, ok := dest[0].(*int); ok {
						*p = 0
					}
				}
				return nil
			}}
		},
	}
	logger := zap.NewNop()
	h := NewPacienteHandler(pool, logger)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/test/supabase", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.TestSupabaseConnection(ctx)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d body=%s", rec.Code, rec.Body.String())
	}
}

// InspectTables success path: returns at least one column
func TestInspectTables_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	pool := &mockPoolP{
		queryFunc: func(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
			rows := &mockRowsP{rowsCount: 1}
			rows.scanFunc = func(dest ...interface{}) error {
				if len(dest) >= 4 {
					if p, ok := dest[0].(*string); ok {
						*p = "id"
					}
					if p, ok := dest[1].(*string); ok {
						*p = "uuid"
					}
					if p, ok := dest[2].(*string); ok {
						*p = "NO"
					}
					if p, ok := dest[3].(**string); ok {
						*p = nil
					}
				}
				return nil
			}
			return rows, nil
		},
	}
	logger := zap.NewNop()
	h := NewPacienteHandler(pool, logger)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/inspect/tables?table=pacientes", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.InspectTables(ctx)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d body=%s", rec.Code, rec.Body.String())
	}
}

// ConnectAllTables success: at least one table has count>0 and columns/sample fetched
func TestConnectAllTables_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	pool := &mockPoolP{
		pingFunc: func(ctx context.Context) error { return nil },
		queryRowFunc: func(ctx context.Context, sql string, args ...interface{}) pgx.Row {
			// count queries
			if strings.HasPrefix(strings.ToUpper(strings.TrimSpace(sql)), "SELECT COUNT(*)") {
				return mockRowP{scanFunc: func(dest ...interface{}) error {
					if p, ok := dest[0].(*int); ok {
						if strings.Contains(strings.ToUpper(sql), "PACIENTES") {
							*p = 1
						} else {
							*p = 0
						}
					}
					return nil
				}}
			}
			return mockRowP{scanFunc: func(dest ...interface{}) error { return pgx.ErrNoRows }}
		},
		queryFunc: func(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
			up := strings.ToUpper(sql)
			if strings.Contains(up, "INFORMATION_SCHEMA.COLUMNS") {
				rows := &mockRowsP{rowsCount: 1}
				rows.scanFunc = func(dest ...interface{}) error {
					if len(dest) >= 4 {
						if p, ok := dest[0].(*string); ok {
							*p = "id"
						}
						if p, ok := dest[1].(*string); ok {
							*p = "uuid"
						}
						if p, ok := dest[2].(*string); ok {
							*p = "NO"
						}
						if p, ok := dest[3].(**string); ok {
							*p = nil
						}
					}
					return nil
				}
				return rows, nil
			}
			if strings.Contains(up, "LIMIT 3") {
				rows := &mockRowsP{rowsCount: 1}
				rows.scanFunc = func(dest ...interface{}) error { return nil }
				return rows, nil
			}
			return &mockRowsP{rowsCount: 0}, nil
		},
	}
	logger := zap.NewNop()
	h := NewPacienteHandler(pool, logger)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/connect/all-tables", nil)
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	h.ConnectAllTables(ctx)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d body=%s", rec.Code, rec.Body.String())
	}
}
