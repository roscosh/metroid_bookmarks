package main

import (
	"flag"
	"metroid_bookmarks/internal/app"
	"metroid_bookmarks/internal/models"
	"metroid_bookmarks/pkg/misc/log"
)

// @title METROID BOOKMARKS API
// @version 1.0
// @description API Server for metroid bookmarks
// @host localhost:3000
// @BasePath /api/v1
func main() {
	flag.Parse()

	envConf, err := models.NewEnvConfig()
	if err != nil {
		panic(err.Error())
	}

	logger := log.GetLogger()
	logger.SetParams(envConf.LogLevel)

	appObj := app.NewApp(envConf)
	appObj.Init()
}
