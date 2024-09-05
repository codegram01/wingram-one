package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	Mode                                          string
	Port                                          string
	DbHost, DbPort, DbUser, DbName, DbPass, DbSSL string
}

func Init() *Config {
	cfg := &Config{
		Mode: GetEnv("WINGRAM_MODE", "LOCAL"),
		Port: GetEnv("WINGRAM_PORT", "8080"),

		DbHost: GetEnv("WINGRAM_DB_HOST", "localhost"),
		DbPort: GetEnv("WINGRAM_DB_PORT", "5432"),
		DbUser: GetEnv("WINGRAM_DB_USER", "testuser"),
		DbName: GetEnv("WINGRAM_DB_NAME", "testdb"),
		DbPass: GetEnv("WINGRAM_DB_PASS", "testpassword"),
		DbSSL:  GetEnv("WINGRAM_DB_SSL", "disable"),
	}
	return cfg
}

// this function load env value
// if not have value it will assign defaultVal
func GetEnv(key string, defaultVal string) string {
	val := os.Getenv(key)

	if val == "" {
		val = defaultVal
	}

	return val
}

// Stop running if not have some env
func MustGetEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.", k)
	}
	return v
}

func (cfg *Config) GetDbConStr() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbName,
		cfg.DbPass,
		cfg.DbSSL,
	)
}
