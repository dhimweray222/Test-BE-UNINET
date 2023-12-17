package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

func NewPostgresPool() *pgxpool.Pool {

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	poolConfig, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		log.Fatal(err)
	}

	minConnsInt, err := strconv.Atoi(os.Getenv("DB_POOL_MIN"))

	maxConnsInt, err := strconv.Atoi(os.Getenv("DB_POOL_MAX"))

	poolConfig.MinConns = int32(minConnsInt)
	poolConfig.MaxConns = int32(maxConnsInt)

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	log.Println(pool)
	if err != nil {
		log.Fatal(err)
	}

	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := pool.Ping(c); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected", dsn)
	return pool
}
