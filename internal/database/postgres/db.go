package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"auth-service-backend/config"
	"auth-service-backend/internal/utils"
	"auth-service-backend/pkg/sqlcqueries"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DataBase struct {
	db      *pgxpool.Pool
	queries *sqlcqueries.Queries
}

func NewDataBase() *DataBase {
	return &DataBase{}
}

func (database *DataBase) InitializationDB() error {
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DB,
	)

	cfg, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return errors.New("unable to parse database URL")
	}

	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.HealthCheckPeriod = 5 * time.Minute

	database.db, err = pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return errors.New("unable to connect to database")
	}

	database.queries = sqlcqueries.New(database.db)
	log.Println("Connected to PostgreSQL!")

	err = utils.RunMigrations()
	if err != nil {
		return err
	}

	return nil
}

func (database *DataBase) GetDB() *pgxpool.Pool {
	return database.db
}

func (database *DataBase) GetQueries() *sqlcqueries.Queries {
	return database.queries
}

func (database *DataBase) PingDB() error {
	return database.db.Ping(context.Background())
}
