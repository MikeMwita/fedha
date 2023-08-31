package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Printf("Error connecting to database:%v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
}
