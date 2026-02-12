package postgres

import (
	"context"
	"fmt"
	"gopher-finance-engine/configs"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func NewPostgresConn(ctx context.Context, conn *configs.PostgresConfigs) {
	if conn == nil {
		log.Fatal("string connection is nil")
	}
	databaseURL := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		conn.Host,
		conn.Port,
		conn.User,
		conn.Password,
		conn.DbName,
	)

	if conn.Host == "" || conn.Port == "" || conn.User == "" || conn.Password == "" || conn.DbName == "" {
		log.Fatalf("Missing database configuration values: HOST=%s PORT=%s USER=%s PASSWORD=%s DBNAME=%s",
			conn.Host, conn.Port, conn.User, conn.Password, conn.DbName)
	}

	config, err := pgxpool.ParseConfig(databaseURL)

	if err != nil {
		log.Fatalf("unable to parse database URL: %v", err)
	}
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	config.ConnConfig.StatementCacheCapacity = 0

	db, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("unable to create connection pool: %v", err)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatalf("unable to ping database: %v", err)
	}

	log.Println("Database connected ✅")
}
