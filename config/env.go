package config

import "os"

type environmentVariables struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDbName   string
}

var Env *environmentVariables

func LoadEnv() {
	env := &environmentVariables{}

	env.PostgresDbName = os.Getenv("DB_NAME")
	env.PostgresPassword = os.Getenv("DB_PASSWORD")
	env.PostgresPort = os.Getenv("DB_PORT")
	env.PostgresHost = os.Getenv("DB_HOST")
	env.PostgresUser = os.Getenv("DB_USER")

	Env = env
}
