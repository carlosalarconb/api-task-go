package routers

import (
	controllers "api-tasks/controllers/tasks"

	"github.com/gin-gonic/gin"
)

// GetAll get all the tasks
func GetAll(c *gin.Context) {
	c.JSON(200, controllers.GetAll())
}

//GetOne get one task
func GetOne(c *gin.Context) {
	c.JSON(200, controllers.GetOne(c.Param("id")))
}
