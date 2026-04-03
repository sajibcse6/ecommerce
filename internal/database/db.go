package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(dbUrl string) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	if err := db.Ping(ctx); err !=  nil {
		log.Fatal("Database not reachable: ", err)
	}

	log.Println("Database connected")

	return db
}