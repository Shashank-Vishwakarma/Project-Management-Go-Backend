package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"time"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/config"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/constants"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/database"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/lib"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	var body config.ProjectRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if body.Name=="" || body.Description=="" {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, "All fields are required", nil)
		return
	}

	// check if the logged in user is creating the project
	claims := r.Context().Value(constants.USER_CONTEXT_KEY).(jwt.MapClaims)

	ownerID, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		lib.HandleResponse(w, http.StatusInternalServerError, "Something went wrong...", nil)
		return
	}

	if ownerID != body.OwnerID {
		w.WriteHeader(http.StatusUnauthorized)
		lib.HandleResponse(w, http.StatusUnauthorized, "Unauthorized - Please login to create a project", nil)
		return
	}

	// check if user exists or not
	user := database.DBClient.Model(&models.User{}).Where("id = ?", body.OwnerID).First(&models.User{})
	if user.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		lib.HandleResponse(w, http.StatusNotFound, "User does not exist", nil)
		return
	}

	// create project
	project := &models.Project{
		Name:        body.Name,
		Description: body.Description,
		OwnerID:     body.OwnerID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := database.DBClient.Model(&models.Project{}).Create(&project)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.HandleResponse(w, http.StatusInternalServerError, result.Error.Error(), nil)
		return
	}

	lib.HandleResponse(w, http.StatusCreated, "Project created successfully", project)
}

func GetAllProjectsHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(constants.USER_CONTEXT_KEY).(jwt.MapClaims)
	userID, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		lib.HandleResponse(w, http.StatusInternalServerError, "Something went wrong...", nil)
		return
	}

	var projects []models.Project
	result := database.DBClient.Model(&models.Project{}).Preload("Owner").Find(&projects)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.HandleResponse(w, http.StatusInternalServerError, result.Error.Error(), nil)
		return
	}

	var response []models.Project
	for _, project := range projects {
		if project.OwnerID == userID || slices.Contains(project.MemeberIDs, userID) {
			response = append(response, project)
		}
	}

	lib.HandleResponse(w, http.StatusOK, "", response)
}

func GetProjectDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID, err := uuid.Parse(vars["project_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, "Could not parse the project id", nil)
		return
	}

	var project models.Project
	result := database.DBClient.Model(&models.Project{}).Where("id = ?", projectID).Preload("Owner").Find(&project)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		lib.HandleResponse(w, http.StatusNotFound, "Project not found", nil)
		return
	}

	lib.HandleResponse(w, http.StatusOK, "Project details obtained...", project)
}

func UpdateProjectHandler(w http.ResponseWriter, r *http.Request) {}

func DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {}

func AddMemberToProject(w http.ResponseWriter, r *http.Request) {}

func RemoveMemberFromProject(w http.ResponseWriter, r *http.Request) {}
