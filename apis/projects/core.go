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
	for i := range initialDB.Projects {
		if uuid == initialDB.Projects[i].ID {
			project = initialDB.Projects[i]
			break
		}
	}
	if (project == Project{}) {
		message := fmt.Sprintf("Unable to find %s project.", uuid)
		err = errors.New(message)
	}
	return
}

// UpdateProjectCore : Update a project given its
// Name, and Description
func UpdateProjectCore(uuid string, projectRequest *CreateProjectRequest) (project Project, err error) {
	for i := range initialDB.Projects {
		if uuid == initialDB.Projects[i].ID {
			initialDB.Projects[i].Name = projectRequest.Name
			initialDB.Projects[i].Description = projectRequest.Description
			project = initialDB.Projects[i]
			break
		}
	}
	if (project == Project{}) {
		message := fmt.Sprintf("Unable to find %s project.", uuid)
		err = errors.New(message)
	}
	return
}

// DeleteProjectCore : Delete a project given its ID
func DeleteProjectCore(uuid string) (err error) {
	for i := range initialDB.Projects {
		if uuid == initialDB.Projects[i].ID {
			initialDB.Projects = append(initialDB.Projects[:i], initialDB.Projects[i+1:]...)
			break
		}
	}
	return
}
