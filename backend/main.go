package main

import (
	"dagger/backend/gin/databases"
	router "dagger/backend/gin/routers"
	"dagger/backend/gin/runtime"
	"dagger/backend/gin/utils"
	"fmt"
	"os"

	"go.uber.org/zap"
)

// @title dagger backend api
// @version 1.0.0
// @description this is dagger backend api server
// @BasePath /
func main() {
	if runtime.LokiServer == "" {
		runtime.LokiServer = os.Getenv("LOKI_SERVER")
		if runtime.LokiServer == "" {
			utils.Log4Zap(zap.ErrorLevel).Error(fmt.Sprintf("start server error: missing loki-server param or LOKI_SERVER env"))
			return
		}
	}

	if runtime.Migrate {
		databases.MigrateDB(databases.DB)
	}

	db, _ := databases.DB.DB()
	defer db.Close()

	router.InitRouter()
}
