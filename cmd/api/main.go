package main

import (
	"gopher-finance-engine/configs"
	"gopher-finance-engine/internal/application"
)

func main() {
	configs.LoadConfigs()
	app := application.NewApplication()

	defer app.Logger.Sync()

	app.Routes.StartServer(":8080")
	app.Logger.Info("INICIALIZA LOGO!")
}
