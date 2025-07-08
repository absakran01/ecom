package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct{
	PublicHost string
	Port string
	DBUser string
	DBPassword string
	DBAddress string
	DBName string
	JWTExpiration int64 // in seconds
	JWTSecret []byte // JWT secret key
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "localhost"),
		Port: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "123123123"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.1"), getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "ecom"),
		JWTExpiration: getEnvAsInt("JWT_EXPIRATION", 3600*24*7), // 7 days in seconds
		JWTSecret: []byte(getEnv("JWT_SECRET", "mysecret")),
	}
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return fallback
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}