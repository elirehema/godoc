package main

import (
	"godocker/api"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	gin.SetMode(gin.DebugMode)
	api.CreateUrlMappings()
	//LIsten to server at port 8080
	//mappings.Router.Run(":8080")
}
