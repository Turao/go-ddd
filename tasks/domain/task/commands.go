package task

type CreateTaskCommand struct {
	ProjectID   ProjectID
	Title       string
	Description string
}

type AssignToUserCommand struct {
	UserID UserID
}

type UnassignCommand struct{}

type UpdateTitleCommand struct {
	Title string
}

type UpdateDescriptionCommand struct {
	Description string
}

type UpdateStatusCommand struct {
	Status string
}
