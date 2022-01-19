# go-ddd

Kind of trying to create a Project Manager (e.g. Jira) with DDD, CQRS, and Event Sourcing in mind...

- User
  - Events: UserRegistered
- Project
  - Events: ProjectCreated, ProjectUpdated, ProjectDeleted
- Task
  - Events: TaskCreated, TaskAssigned, TaskUnassigned, DescriptionUpdated
