package projects

// Project : definition of a project
// structure
type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ListProjects : definition of the list
// of projects structure
type ListProjects struct {
	Projects []Project `json:"projects"`
}

// CreateProjectRequest : definition of a
// project creation request structure
type CreateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateProjectRequest : definition of a
// project update request structure
type UpdateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ErrorResponse : definition of a not found
// default response
type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
