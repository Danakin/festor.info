package psql

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func connectionString(config Config) string {
	fmt.Println(config)
	if config.SSLMode == "" {
		config.SSLMode = "disable"
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
		config.SSLMode,
	)
}

func Connect(config Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", connectionString(config))
	if err != nil {
		return nil, err
	}
	return db, nil
}
