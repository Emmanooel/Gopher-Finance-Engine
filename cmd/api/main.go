package main

import (
	"gopher-finance-engine/configs"
	"gopher-finance-engine/internal/application"
)

func main() {
	configs.LoadConfigs()
	route := application.NewApplication()

	route.StartServer(":8080")
}
