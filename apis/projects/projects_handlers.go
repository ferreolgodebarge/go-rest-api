package projects

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ListProjectHandler : Lists all projects
func ListProjectHandler(c *gin.Context) {
	res, err := ListProjectsCore()
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, res)
	}
}

// CreateProjectHandler : Create a new project
func CreateProjectHandler(c *gin.Context) {
	projectRequest := new(CreateProjectRequest)
	err := c.Bind(projectRequest)
	if err != nil {
		fmt.Println(err)
	}
	res, err := CreateProjectCore(projectRequest)
	c.JSON(201, res)
}

// GetProjectHandler : Get a project given its ID
func GetProjectHandler(c *gin.Context) {
	uuid := c.Param("uuid")
	res, err := GetProjectCore(uuid)
	if err != nil {
		c.JSON(404, ErrorResponse{ErrorCode: 404, ErrorMessage: err.Error()})
	} else {
		c.JSON(200, res)
	}
}

// UpdateProjectHandler : Update a project given
// its ID, Name, and Description
func UpdateProjectHandler(c *gin.Context) {
	// uuid := c.Param("uuid")
}
