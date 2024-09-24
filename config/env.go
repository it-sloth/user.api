package config

import "os"

type Env struct {
	Port string
}

func GetEnv() *Env {
	return &Env{
		Port: os.Getenv("PORT"),
	}
}
