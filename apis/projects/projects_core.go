package projects

import (
	"errors"
	"fmt"
)

// ListProjectsCore : list all projects
func ListProjectsCore() (listProjects ListProjects, err error) {
	listProjects = initialDB
	return
}

// CreateProjectCore : Create a new project
func CreateProjectCore(request *CreateProjectRequest) (project Project, err error) {
	project = Project{
		ID:          GenerateUUID(),
		Name:        request.Name,
		Description: request.Description,
	}
	initialDB.Projects = append(initialDB.Projects, project)
	return
}

// GetProjectCore : Get a project by ID
func GetProjectCore(uuid string) (project Project, err error) {
	for _, e := range initialDB.Projects {
		if uuid == e.ID {
			project = e
			break
		}
	}
	if (project == Project{}) {
		message := fmt.Sprintf("Unable to find %s project.", uuid)
		err = errors.New(message)
	}
	return
}
