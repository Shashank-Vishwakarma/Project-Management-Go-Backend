package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/config"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/constants"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/database"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/lib"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID, err := uuid.Parse(vars["project_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, "Could not parse the project id", nil)
		return
	}

	var project models.Project
	result := database.DBClient.Model(&models.Project{}).Where("id = ?", projectID).First(&project)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		lib.HandleResponse(w, http.StatusNotFound, "Project not found", nil)
		return
	}

	claims := r.Context().Value(constants.USER_CONTEXT_KEY).(jwt.MapClaims)
	userID, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		lib.HandleResponse(w, http.StatusInternalServerError, "Something went wrong...", nil)
		return
	}

	isPartOfTheProject := false
	if userID == project.OwnerID {
		isPartOfTheProject = true
	}
	for _, member := range project.Members {
		if member.ID == userID {
			isPartOfTheProject = true
			break
		}
	}

	if !isPartOfTheProject {
		w.WriteHeader(http.StatusUnauthorized)
		lib.HandleResponse(w, http.StatusUnauthorized, "Not authorized to create a task in this project", nil)
		return
	}

	var body config.TaskRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	result = database.DBClient.Model(&models.Task{}).Create(&models.Task{
		Title: body.Title,
		Description: body.Description,
		Status: body.Status,
		DueDate: body.DueDate,
		CreatedByID: userID,
		AssignedToID: body.AssignedToID,
		ProjectID: projectID,
	})
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.HandleResponse(w, http.StatusInternalServerError, "Error creating task", nil)
		return
	}

	lib.HandleResponse(w, http.StatusCreated, "Task created successfully", nil)
}

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID, err := uuid.Parse(vars["project_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, "Could not parse the project id", nil)
		return
	}

	var project models.Project
	result := database.DBClient.Model(&models.Project{}).Where("id = ?", projectID).First(&project)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		lib.HandleResponse(w, http.StatusNotFound, "Project not found", nil)
		return
	}

	claims := r.Context().Value(constants.USER_CONTEXT_KEY).(jwt.MapClaims)
	userID, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		lib.HandleResponse(w, http.StatusInternalServerError, "Something went wrong...", nil)
		return
	}

	isPartOfTheProject := false
	if userID == project.OwnerID {
		isPartOfTheProject = true
	}
	for _, member := range project.Members {
		if member.ID == userID {
			isPartOfTheProject = true
			break
		}
	}

	if !isPartOfTheProject {
		w.WriteHeader(http.StatusUnauthorized)
		lib.HandleResponse(w, http.StatusUnauthorized, "Not authorized to see the tasks of this project", nil)
		return
	}

	var tasks []models.Task
	result = database.DBClient.Model(&models.Task{}).Where("project_id = ?", projectID).Preload("CreatedByUser").Preload("AssignedTo").Preload("Project").Find(&tasks)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.HandleResponse(w, http.StatusInternalServerError, "Error getting tasks", nil)
		return
	}

	lib.HandleResponse(w, http.StatusOK, "Tasks fetched successfully", tasks)
}

func GetTaskDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID, err := uuid.Parse(vars["project_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, "Could not parse the project id", nil)
		return
	}
	taskID, err := uuid.Parse(vars["task_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, "Could not parse the project id", nil)
		return
	}

	var project models.Project
	result := database.DBClient.Model(&models.Project{}).Where("id = ?", projectID).First(&project)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		lib.HandleResponse(w, http.StatusNotFound, "Project not found", nil)
		return
	}

	var task models.Task
	result = database.DBClient.Model(&models.Task{}).Where("id = ?", taskID).Preload("CreatedByUser").Preload("AssignedTo").Preload("Project").First(&task)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		lib.HandleResponse(w, http.StatusNotFound, "Task not found", nil)
		return
	}

	claims := r.Context().Value(constants.USER_CONTEXT_KEY).(jwt.MapClaims)
	userID, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		lib.HandleResponse(w, http.StatusInternalServerError, "Something went wrong...", nil)
		return
	}

	isPartOfTheProject := false
	if userID == project.OwnerID {
		isPartOfTheProject = true
	}
	for _, member := range project.Members {
		if member.ID == userID {
			isPartOfTheProject = true
			break
		}
	}

	if !isPartOfTheProject {
		w.WriteHeader(http.StatusUnauthorized)
		lib.HandleResponse(w, http.StatusUnauthorized, "Not authorized to see the task details", nil)
		return
	}

	lib.HandleResponse(w, http.StatusOK, "Task details fetched successfully", task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID, err := uuid.Parse(vars["project_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, "Could not parse the project id", nil)
		return
	}
	taskID, err := uuid.Parse(vars["task_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, "Could not parse the project id", nil)
		return
	}

	var body config.TaskRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if body.Title == "" && body.Description == "" && body.Status == "" && body.AssignedToID == uuid.Nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, "No changes found", nil)
		return
	}

	var project models.Project
	result := database.DBClient.Model(&models.Project{}).Where("id = ?", projectID).First(&project)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		lib.HandleResponse(w, http.StatusNotFound, "Project not found", nil)
		return
	}

	var task models.Task
	result = database.DBClient.Model(&models.Task{}).Where("id = ?", taskID).First(&task)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.HandleResponse(w, http.StatusInternalServerError, "Task not found", nil)
		return
	}

	claims := r.Context().Value(constants.USER_CONTEXT_KEY).(jwt.MapClaims)
	userID, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		lib.HandleResponse(w, http.StatusInternalServerError, "Something went wrong...", nil)
		return
	}

	isPartOfTheProject := false
	if userID == project.OwnerID {
		isPartOfTheProject = true
	}
	for _, member := range project.Members {
		if member.ID == userID {
			isPartOfTheProject = true
			break
		}
	}

	if !isPartOfTheProject {
		w.WriteHeader(http.StatusUnauthorized)
		lib.HandleResponse(w, http.StatusUnauthorized, "Not authorized to update the task", nil)
		return
	}

	if body.Title == "" {
		body.Title = task.Title
	}

	if body.Description == "" {
		body.Description = task.Description
	}

	if body.Status == "" {
		body.Status = task.Status
	}

	if body.AssignedToID == uuid.Nil {
		body.AssignedToID = task.AssignedToID
	}

	result = database.DBClient.Model(&models.Task{}).Where("id = ?", taskID).Updates(&models.Task{
		Title:        body.Title,
		Description:  body.Description,
		Status:       body.Status,
		// DueDate:      body.DueDate,
		AssignedToID: body.AssignedToID,
	})
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.HandleResponse(w, http.StatusInternalServerError, "Error updating the task", nil)
		return
	}

	lib.HandleResponse(w, http.StatusOK, "Task updated successfully", nil)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID, err := uuid.Parse(vars["project_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, "Could not parse the project id", nil)
		return
	}
	taskID, err := uuid.Parse(vars["task_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, "Could not parse the project id", nil)
		return
	}

	var project models.Project
	result := database.DBClient.Model(&models.Project{}).Where("id = ?", projectID).First(&project)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		lib.HandleResponse(w, http.StatusNotFound, "Project not found", nil)
		return
	}

	claims := r.Context().Value(constants.USER_CONTEXT_KEY).(jwt.MapClaims)
	userID, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		lib.HandleResponse(w, http.StatusInternalServerError, "Something went wrong...", nil)
		return
	}

	isPartOfTheProject := false
	if userID == project.OwnerID {
		isPartOfTheProject = true
	}
	for _, member := range project.Members {
		if member.ID == userID {
			isPartOfTheProject = true
			break
		}
	}

	if !isPartOfTheProject {
		w.WriteHeader(http.StatusUnauthorized)
		lib.HandleResponse(w, http.StatusUnauthorized, "Not authorized to delete the task", nil)
		return
	}

	result = database.DBClient.Model(&models.Task{}).Where("id = ?", taskID).Delete(&models.Task{})
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.HandleResponse(w, http.StatusInternalServerError, "Error deleting the task or task does not exist", nil)
		return
	}

	lib.HandleResponse(w, http.StatusOK, "Task deleted successfully", nil)
}

func ChangeTaskAssigneeHandler(w http.ResponseWriter, r *http.Request) {}
