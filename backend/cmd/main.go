package main

import (
	"metroid_bookmarks/internal/app"
	"metroid_bookmarks/internal/models"
	"metroid_bookmarks/pkg/misc"
)

// @title METROID BOOKMARKS API
// @version 1.0
// @description API Server for metroid bookmarks
// @host localhost:3000
// @BasePath /api/v1
func main() {
	envConf, err := models.NewEnvConfig()
	if err != nil {
		panic(err.Error())
		return
	}

	logger := misc.GetLogger()
	logger.SetParams(envConf.LogLevel)

	appObj := app.NewApp(envConf)
	appObj.Init()
}
