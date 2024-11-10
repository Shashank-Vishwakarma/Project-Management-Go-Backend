package constants

const (
	USER_ADMIN_ROLE = "admin"

	BASE_ENDPOINT          = "/api/v1"
	AUTH_BASE_ENDPOINT     = BASE_ENDPOINT + "/auth"
	USERS_BASE_ENDPOINT    = BASE_ENDPOINT + "/users"
	PROJECTS_BASE_ENDPOINT = BASE_ENDPOINT + "/projects"
	TASKS_BASE_ENDPOINT    = BASE_ENDPOINT + "/projects/{project_id}/tasks"
)