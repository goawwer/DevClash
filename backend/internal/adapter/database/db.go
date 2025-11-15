package database

import (
	"context"
	"fmt"

	"github.com/goawwer/devclash/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type ENV string

const (
	EnvTest ENV = "test"
	EnvDev  ENV = "dev"
)

type Config struct {
	Name     string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`

	// test database for unit tests
	TestName string `env:"TEST_DB_NAME"`
	TestPort int    `env:"TEST_DB_PORT"`

	ENV ENV `env:"ENV" envDefault:"dev"`
}

func (c *Config) DSN() string {
	name := c.Name
	port := c.Port

	if c.ENV != EnvDev {
		name = c.TestName
		port = c.TestPort
	}

	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		port,
		name,
	)
}

type Db struct {
	*sqlx.DB
}

var database *Db

func Init(ctx context.Context, cfg *Config) error {
	db, err := sqlx.Open("postgres", cfg.DSN())
	// Open may just validate its arguments without creating a connection to the database.
	// To verify that the data source name is valid, call [DB.Ping].
	if err != nil {
		return fmt.Errorf("failed to validate arguments for database connection: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to create a connection to database: %w", err)
	}

	logger.WithFields(logrus.Fields{
		"component": "database",
		"address":   cfg.DSN(),
	}).Info("started successfully")

	database = &Db{db}

	return nil
}

func Get() *sqlx.DB {
	return database.DB
}

func Close() error {
	return database.DB.Close()
}
