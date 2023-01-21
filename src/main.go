package main

import (
	"github.com/beschlz/dartsly-match-service/src/database"
	"github.com/beschlz/dartsly-match-service/src/status"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.MigrateDB()

	status.RegisterStatusEndpoints(r)

	r.Run()
}
