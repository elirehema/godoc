package api

import (
	"godocker/controllers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//var Router *gin.Engine

func CreateUrlMappings() {
	router := gin.Default()
	//router.Use(cors.Default())

	//v1 of the API
	v1 := router.Group("/api/v1")
	{
		v1.POST("/agents", controllers.CreateAgent)
		v1.GET("/agents", controllers.GetAgents)

	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
