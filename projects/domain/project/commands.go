package project

type CreateProjectCommand struct {
	Name      string
	CreatedBy UserID
}

type UpdateProjectCommand struct {
	Name string
}

type DeleteProjectCommand struct{}
