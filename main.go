package main

import (
	projects "./apis/projects"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/projects", projects.ListProjectHandler)
		v1.POST("/projects", projects.CreateProjectHandler)
		v1.GET("/projects/:uuid", projects.GetProjectHandler)
	}
	r.Run("0.0.0.0:8080")
}
