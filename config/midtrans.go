package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var c coreapi.Client

func globalConfig(){
	if err := godotenv.Load() ; err != nil {
		panic(err.Error())
	}

	var environment midtrans.EnvironmentType

	env := os.Getenv("APP_ENV")
	SERVERKEY := os.Getenv("SERVERKEY")
	CLIENTKEY := os.Getenv("CLIENTKEY")

	switch env {
	case "development":
		environment = midtrans.Sandbox
		break
	case "production":
		environment =  midtrans.Production
		break
	default:
		environment = midtrans.Sandbox
	}

	c.ServerKey = SERVERKEY
	c.ClientKey = CLIENTKEY
	midtrans.Environment = environment
}

func RunMidtransConfig(){
	globalConfig()
}