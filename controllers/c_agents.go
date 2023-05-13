package controllers

import (
	"encoding/json"
	"fmt"
	"godocker/datasource"
	"godocker/models"
	"godocker/redis"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateAgent(c *gin.Context) {
	var agent models.Agent
	if err := c.ShouldBindJSON(&agent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := datasource.Instance.Create(&agent); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
	} else {

		c.JSON(http.StatusOK, agent)

	}
}
func GetAgents(c *gin.Context) {
	var agents []models.Agent
	objStr, err := redis.RedisClient.Get("agents").Result()
	if err != nil {
		fmt.Printf("ERROR: %v", err.Error())
	}
	if objStr != "" {
		var slice []models.Agent
		err = json.Unmarshal([]byte(objStr), &slice)
		if err != nil {
			fmt.Println(err)
			return
		}
		print("From Redis")
		c.JSON(http.StatusOK, slice)
	} else {
		if result := datasource.Instance.Find(&agents); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error})
		} else {
			print("From Query")
			json, _ := json.Marshal(agents)
			redis.RedisClient.Set("agents", string(json), 0)
			c.JSON(http.StatusOK, agents)
		}
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
