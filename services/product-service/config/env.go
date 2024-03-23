package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv        string `map.structure:"APP_ENV"`
	ServerAddress string `map.structure:"SERVER_ADDRESS"`
	DBHost        string `map.structure:"DB_HOST"`
	DBPort        string `map.structure:"DB_PORT"`
	DBUser        string `map.structure:"DB_USER"`
	DBPass        string `map.structure:"DB_PWD"`
	DBName        string `map.structure:"DB_NAME"`
}

func createEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		panic(err)
	}

	if env.AppEnv == "development" {
		log.Println("Service running in development environment")
	}

	return &env
}
