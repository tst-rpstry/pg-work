package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type dbConfig struct {
	username string `json:"pg_user"`
	password string `json:"pg_password"`
	dbName   string `json:"pg_dbname"`
	host     string `json:"pg_host"`
	port     string `json:"pg_port"`
}

// NewConfig is empty.
func NewConfig() *dbConfig {
	return &dbConfig{}
}

func (cfg *dbConfig) FromFileOrDefault(path string) (*dbConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(f).Decode(cfg); err != nil {
		return nil, err
	}

	if cfg.dbName == "" {
		cfg.dbName = "postgres"
	}
	if cfg.username == "" {
		cfg.username = "postgres"
	}

	if cfg.host == "" {
		cfg.host = "127.0.0.1"
	}
	if cfg.port == "" {
		cfg.port = "5432"
	}
	return cfg, nil
}

func (cfg *dbConfig) ConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.username, cfg.password, cfg.host, cfg.port, cfg.dbName)
}
