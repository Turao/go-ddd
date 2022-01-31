# go-ddd

Kind of trying to create a Project Manager (e.g. Jira)...

## What I want to practice:
- Clean Architecture (e.g. domain, application, infrastructure layers) 
- DDD (e.g. Aggregates, Domain Events, ...)
- CQRS (e.g. Read / Write models)
  - Have some projections to list tasks by project, tasks by user, and do some full-text searching based on the task description
- Event Sourcing
  - Billing context, to track how many tasks have been assigned and completed by each user so we can pay them later... IDK

## Contexts

#### User
- Commands: 
  - `RegisterUsed`
- Queries: 
  - `ListUsers`
- Events: 
  - `UserRegistered`

#### Project
- Commands: 
  - `CreateProject`
  - `UpdateProject`
  - `DeleteProject`
- Queries: 
  - `FindProject` (by ID)
  - `ListProjects`
- Events: 
  - `ProjectCreated`
  - `ProjectUpdated`
  - `ProjectDeleted`

#### Task
- Commands: 
  - `CreateTask`
  - `AssignToUser`
  - `UnassignUser`
  - `UpdateTitle`
  - `UpdateDescription`
  - `UpdateStatus`
- Queries:
  - `TasksByProject`
  - `TasksByAssignedUser`
- Events: 
  - `TaskCreated`
  - `TaskAssigned`
  - `TaskUnassigned`
  - `TitleUpdated`
  - `DescriptionUpdated`
  - `StatusUpdated`


> Tasks should likely be part of Project's context (Projects often don't have that many tasks)... but (1) I don't care about strong consistency between Project and Task really that much (2) let's pretend users will create a lot of tasks for each project, so we need to move the collection to its own bounded context...

#### Billing
- Commands:
  - `CreateAccount`
  - `AddTaskToUser`
  - `RemoveTaskFromUser`
- Queries:
  - AccountDetails (TODO)
    - get information like invoice id + number of tasks concluded + what the user needs to pay
- Events:
  - `AccountCreated`
  - `TaskAdded`
    - whenever an user is assigned to a task, we add it into the User's Account Invoice
  - `TaskRemoved`
    - whenever an user is unassigned from a task, we remove it from the User's Account Invoice

Will listen to:
- `UserRegistered`
- `TaskAssigned`
- `TaskUnassigned`
- `StatusUpdated`