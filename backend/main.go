package main

import (
	"fmt"
	"saifutdinov/believe-or-not/backend/api"
	"saifutdinov/believe-or-not/backend/packages/dotenv"
)

func main() {
	config, err := dotenv.LoadEnvFile()
	if err != nil {
		panic(fmt.Sprintf("can't load .env file: %v.", err.Error()))
	}

	api.StartListen(config)
}
