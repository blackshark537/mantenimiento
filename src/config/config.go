package config

import (
	"os"
	"runtime"
)

const VERSION string = "0.0.1"

const AppName string = "Dataprod - Maintenance"
const AppDesc string = "A simple CLI system to manage preventive maintenance"
const AppArch string = runtime.GOARCH

type env struct {
	GRPC_HOST string
	GRPC_PORT string
}

func GetEnv() *env {
	e := new(env)

	if os.Getenv("GRPC_HOST") != "" {
		e.GRPC_HOST = os.Getenv("GRPC_HOST")
	} else {
		e.GRPC_HOST = "localhost"
	}

	if os.Getenv("GRPC_PORT") != "" {
		e.GRPC_PORT = os.Getenv("GRPC_HOST")
	} else {
		e.GRPC_PORT = "3000"
	}
	return e
}
