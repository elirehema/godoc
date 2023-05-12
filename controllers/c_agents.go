package controllers

import (
	"godocker/datasource"
	"godocker/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetAgents(c *gin.Context) {
	var agents []models.Agent
	if result := datasource.Instance.Find(&agents); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error})
	} else {
		c.JSON(http.StatusOK, agents)
	}
}
func GetAgentById(c *gin.Context) {
	agentId := c.Params.ByName("id")
	var agent = models.Agent{Id: agentId}
	if result := datasource.Instance.Find(&agent); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error})
	} else {
		c.JSON(http.StatusOK, agent)
	}

}
