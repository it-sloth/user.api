package config

import "os"

type Env struct {
	Port  string
	DbDsn string
}

func GetEnv() *Env {
	return &Env{
		Port:  os.Getenv("PORT"),
		DbDsn: os.Getenv("DB_DSN"),
	}
}
