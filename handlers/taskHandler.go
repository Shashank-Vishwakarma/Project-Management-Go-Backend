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

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {}

func GetTaskDetailsHandler(w http.ResponseWriter, r *http.Request) {}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {}

func ChangeTaskAssigneeHandler(w http.ResponseWriter, r *http.Request) {}
