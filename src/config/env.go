package config

import "os"

type env struct {
	MONGO_URI    string
	APP_HOST     string
	APP_PORT     string
	APP_BASE_URL string
}

var Env env

func getIndividualEnvVariable(name string) string {
	envVar := os.Getenv(name)

	if envVar == "" {
		panic("You must set" + " '" + name + "' " + "environmental variable.")
	}

	return envVar
}

func RegisterEnv() {
	APP_HOST := getIndividualEnvVariable("APP_HOST")
	APP_PORT := getIndividualEnvVariable("APP_PORT")
	MONGO_URI := getIndividualEnvVariable("MONGO_URI")

	Env = env{
		MONGO_URI:    MONGO_URI,
		APP_HOST:     APP_HOST,
		APP_PORT:     APP_PORT,
		APP_BASE_URL: APP_HOST + ":" + APP_PORT,
	}
}
