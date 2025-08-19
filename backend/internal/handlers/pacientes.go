package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// PacienteHandler maneja las operaciones relacionadas con pacientes
type PacienteHandler struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

// NewPacienteHandler crea una nueva instancia del handler de pacientes
func NewPacienteHandler(pool *pgxpool.Pool, logger *zap.Logger) *PacienteHandler {
	return &PacienteHandler{
		pool:   pool,
		logger: logger,
	}
}

// Paciente representa la estructura de un paciente según Supabase
type Paciente struct {
	ID                 int       `json:"id" db:"id"`
	NroDocumento       string    `json:"nro_documento" db:"nro_documento"`
	TipoDocumento      string    `json:"tipo_documento" db:"tipo_documento"`
	Nombre             string    `json:"nombre" db:"nombre"`
	Apellido           string    `json:"apellido" db:"apellido"`
	FechaNacimiento    time.Time `json:"fecha_nacimiento" db:"fecha_nacimiento"`
	Email              string    `json:"email" db:"email"`
	Telefono           *string   `json:"telefono" db:"telefono"`
	Direccion          *string   `json:"direccion" db:"direccion"`
	EstadoCivil        *string   `json:"estado_civil" db:"estado_civil"`
	FechaCreacion      time.Time `json:"fecha_creacion" db:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion" db:"fecha_actualizacion"`
}

// GetPacientes godoc
// @Summary      Obtener lista de pacientes
// @Description  Obtiene todos los pacientes desde Supabase
// @Tags         pacientes
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/pacientes [get]
func (h *PacienteHandler) GetPacientes(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT 
			id, nro_documento, tipo_documento, nombre, apellido, 
			fecha_nacimiento, email, telefono, direccion, estado_civil,
			fecha_creacion, fecha_actualizacion
		FROM pacientes 
		ORDER BY fecha_creacion DESC
	`

	rows, err := h.pool.Query(ctx, query)
	if err != nil {
		h.logger.Error("Error al consultar pacientes", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error interno del servidor",
		})
		return
	}
	defer rows.Close()

	var pacientes []Paciente
	for rows.Next() {
		var p Paciente
		err := rows.Scan(
			&p.ID, &p.NroDocumento, &p.TipoDocumento, &p.Nombre, &p.Apellido,
			&p.FechaNacimiento, &p.Email, &p.Telefono, &p.Direccion, &p.EstadoCivil,
			&p.FechaCreacion, &p.FechaActualizacion,
		)
		if err != nil {
			h.logger.Error("Error al escanear paciente", zap.Error(err))
			continue
		}
		pacientes = append(pacientes, p)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "success",
		"pacientes": pacientes,
		"total":     len(pacientes),
	})
}

// GetPaciente godoc
// @Summary      Obtener paciente por ID
// @Description  Obtiene un paciente específico por su ID
// @Tags         pacientes
// @Produce      json
// @Param        id   path      int  true  "Patient ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/pacientes/{id} [get]
func (h *PacienteHandler) GetPaciente(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de paciente inválido",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT 
			id, nro_documento, tipo_documento, nombre, apellido, 
			fecha_nacimiento, email, telefono, direccion, estado_civil,
			fecha_creacion, fecha_actualizacion
		FROM pacientes 
		WHERE id = $1
	`

	var p Paciente
	err = h.pool.QueryRow(ctx, query, id).Scan(
		&p.ID, &p.NroDocumento, &p.TipoDocumento, &p.Nombre, &p.Apellido,
		&p.FechaNacimiento, &p.Email, &p.Telefono, &p.Direccion, &p.EstadoCivil,
		&p.FechaCreacion, &p.FechaActualizacion,
	)

	if err != nil {
		h.logger.Error("Error al consultar paciente", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Paciente no encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"paciente": p,
	})
}

// TestSupabaseConnection godoc
// @Summary      Probar conexión con Supabase
// @Description  Prueba la conexión con las tablas existentes en Supabase
// @Tags         test
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/test/supabase [get]
func (h *PacienteHandler) TestSupabaseConnection(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Probar conectividad básica
	if err := h.pool.Ping(ctx); err != nil {
		h.logger.Error("Error de ping a Supabase", zap.Error(err))
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "No se puede conectar a Supabase",
		})
		return
	}

	// Contar registros en todas las tablas de Supabase (nombres exactos)
	tables := []string{
		"auditorias",
		"consultorios",
		"datos_personales",
		"historia_clinica_version",  // Corregido: singular
		"historias_clinicas",
		"pacientes",
		"permisos",
		"recetas_medicas",
		"rol_permiso",
		"roles",
		"turnos",
		"usuarios",
	}
	counts := make(map[string]int)
	tableDetails := make(map[string]map[string]interface{})

	for _, table := range tables {
		var count int
		query := "SELECT COUNT(*) FROM " + table
		err := h.pool.QueryRow(ctx, query).Scan(&count)
		if err != nil {
			h.logger.Warn("Error al consultar tabla", zap.String("table", table), zap.Error(err))
			counts[table] = -1
			tableDetails[table] = map[string]interface{}{
				"count": -1,
				"error": err.Error(),
			}
		} else {
			counts[table] = count
			tableDetails[table] = map[string]interface{}{
				"count":  count,
				"status": "accessible",
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":           "success",
		"database":         "connected",
		"supabase_project": "mediapp-db",
		"tables_count":     counts,
		"table_details":    tableDetails,
		"total_tables":     len(tables),
		"connection_pool":  h.pool.Stat(),
		"timestamp":        time.Now(),
	})
}

// InspectTables godoc
// @Summary      Inspeccionar estructura de tablas
// @Description  Obtiene la estructura de las tablas para verificar columnas
// @Tags         test
// @Produce      json
// @Param        table   query     string  false  "Nombre de la tabla a inspeccionar (default: pacientes)"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/inspect/tables [get]
func (h *PacienteHandler) InspectTables(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Obtener tabla a inspeccionar desde query parameter
	tableName := c.DefaultQuery("table", "pacientes")
	
	// Lista de tablas permitidas por seguridad (nombres exactos de Supabase)
	allowedTables := map[string]bool{
		"auditorias":                  true,
		"consultorios":                true,
		"datos_personales":            true,
		"historia_clinica_version":    true,  // Corregido: singular
		"historias_clinicas":          true,
		"pacientes":                   true,
		"permisos":                    true,
		"recetas_medicas":             true,
		"rol_permiso":                 true,
		"roles":                       true,
		"turnos":                      true,
		"usuarios":                    true,
	}

	if !allowedTables[tableName] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tabla no permitida",
			"allowed_tables": []string{
				"auditorias", "consultorios", "datos_personales", 
				"historia_clinica_version", "historias_clinicas", "pacientes",
				"permisos", "recetas_medicas", "rol_permiso", "roles", "turnos", "usuarios",
			},
		})
		return
	}

	// Obtener columnas de la tabla especificada
	query := `
		SELECT column_name, data_type, is_nullable, column_default
		FROM information_schema.columns 
		WHERE table_name = $1
		ORDER BY ordinal_position
	`

	rows, err := h.pool.Query(ctx, query, tableName)
	if err != nil {
		h.logger.Error("Error al consultar estructura de tabla", 
			zap.String("table", tableName), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error interno del servidor",
			"table": tableName,
		})
		return
	}
	defer rows.Close()

	var columns []map[string]interface{}
	for rows.Next() {
		var columnName, dataType, isNullable string
		var columnDefault *string

		err := rows.Scan(&columnName, &dataType, &isNullable, &columnDefault)
		if err != nil {
			h.logger.Error("Error al escanear columna", zap.Error(err))
			continue
		}

		column := map[string]interface{}{
			"name":     columnName,
			"type":     dataType,
			"nullable": isNullable == "YES",
		}
		if columnDefault != nil {
			column["default"] = *columnDefault
		}
		columns = append(columns, column)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"table":   tableName,
		"columns": columns,
		"count":   len(columns),
	})
}

// ConnectAllTables godoc
// @Summary      Conectar y verificar todas las tablas
// @Description  Conecta con todas las tablas de Supabase y obtiene un resumen completo
// @Tags         test
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /api/v1/connect/all-tables [get]
func (h *PacienteHandler) ConnectAllTables(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Lista completa de tablas en Supabase (nombres exactos)
	tables := []string{
		"auditorias",
		"consultorios", 
		"datos_personales",
		"historia_clinica_version",  // Corregido: singular, no plural
		"historias_clinicas",
		"pacientes",
		"permisos",
		"recetas_medicas",
		"rol_permiso",
		"roles",
		"turnos",
		"usuarios",
	}

	allTablesData := make(map[string]interface{})
	successfulConnections := 0
	failedConnections := 0

	for _, tableName := range tables {
		tableData := make(map[string]interface{})
		
		// 1. Contar registros
		var count int
		countQuery := "SELECT COUNT(*) FROM " + tableName
		err := h.pool.QueryRow(ctx, countQuery).Scan(&count)
		if err != nil {
			h.logger.Warn("Error al consultar tabla", 
				zap.String("table", tableName), zap.Error(err))
			tableData["status"] = "error"
			tableData["error"] = err.Error()
			tableData["count"] = -1
			failedConnections++
		} else {
			tableData["status"] = "connected"
			tableData["count"] = count
			successfulConnections++

			// 2. Obtener estructura de tabla (si la conexión fue exitosa)
			columnsQuery := `
				SELECT column_name, data_type, is_nullable, column_default
				FROM information_schema.columns 
				WHERE table_name = $1
				ORDER BY ordinal_position
			`
			
			rows, err := h.pool.Query(ctx, columnsQuery, tableName)
			if err != nil {
				h.logger.Warn("Error al consultar estructura", 
					zap.String("table", tableName), zap.Error(err))
				tableData["columns_error"] = err.Error()
			} else {
				var columns []map[string]interface{}
				for rows.Next() {
					var columnName, dataType, isNullable string
					var columnDefault *string

					err := rows.Scan(&columnName, &dataType, &isNullable, &columnDefault)
					if err != nil {
						continue
					}

					column := map[string]interface{}{
						"name":     columnName,
						"type":     dataType,
						"nullable": isNullable == "YES",
					}
					if columnDefault != nil {
						column["default"] = *columnDefault
					}
					columns = append(columns, column)
				}
				rows.Close()
				tableData["columns"] = columns
				tableData["columns_count"] = len(columns)
			}

			// 3. Obtener muestra de datos (primeros 3 registros si existen)
			if count > 0 {
				sampleQuery := "SELECT * FROM " + tableName + " LIMIT 3"
				sampleRows, err := h.pool.Query(ctx, sampleQuery)
				if err != nil {
					tableData["sample_error"] = err.Error()
				} else {
					// Convertir a slice de maps genérico
					var samples []map[string]interface{}
					for sampleRows.Next() {
						// Obtener nombres de columnas
						fieldDescriptions := sampleRows.FieldDescriptions()
						columnNames := make([]string, len(fieldDescriptions))
						for i, desc := range fieldDescriptions {
							columnNames[i] = string(desc.Name)
						}

						// Crear slice de interfaces para escanear
						values := make([]interface{}, len(columnNames))
						valuePtrs := make([]interface{}, len(columnNames))
						for i := range values {
							valuePtrs[i] = &values[i]
						}

						if err := sampleRows.Scan(valuePtrs...); err == nil {
							row := make(map[string]interface{})
							for i, colName := range columnNames {
								row[colName] = values[i]
							}
							samples = append(samples, row)
						}
					}
					sampleRows.Close()
					tableData["sample_data"] = samples
					tableData["sample_count"] = len(samples)
				}
			}
		}

		allTablesData[tableName] = tableData
	}

	// Resumen final
	summary := map[string]interface{}{
		"status":                "success",
		"total_tables":          len(tables),
		"successful_connections": successfulConnections,
		"failed_connections":    failedConnections,
		"connection_rate":       float64(successfulConnections) / float64(len(tables)) * 100,
		"database":              "Supabase PostgreSQL",
		"project":               "mediapp-db",
		"timestamp":             time.Now(),
		"pool_stats":            h.pool.Stat(),
	}

	c.JSON(http.StatusOK, gin.H{
		"summary": summary,
		"tables":  allTablesData,
	})
}
