package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		fmt.Println("Falta la variable DATABASE_URL")
		return
	}
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		fmt.Println("Error al crear el Pool:", err)
		return
	}
	defer pool.Close()
	err = pool.Ping(context.Background())

	if err != nil {
		fmt.Println("Error al hacer el Ping a la base de datos:", err)
	} else {
		fmt.Println("Conexion exitosa a PostgresSQL")
	}
}
