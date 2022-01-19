# go-ddd

Kind of trying to create a Project Manager (e.g. Jira)...

## What I want to practice:
- Clean Architecture (e.g. domain, application, infrastructure layers) 
- DDD (e.g. Aggregates, Domain Events, ...)
- CQRS (e.g. Read / Write models)
- Event Sourcing

## Contexts

- User
  - Events: UserRegistered
- Project
  - Events: ProjectCreated, ProjectUpdated, ProjectDeleted
- Task*
  - Events: TaskCreated, TaskAssigned, TaskUnassigned, DescriptionUpdated

\* Tasks should likely be part of Project's context (Projects often don't have that many tasks)... but (1) I don't care about strong consistency between Project and Task really that much (2) let's pretend users will create a lot of tasks for each project...

