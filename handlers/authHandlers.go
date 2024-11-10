package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/config"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/database"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/lib"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/models"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/utils"
)

func UserRegistrationHandler(w http.ResponseWriter, r * http.Request) {
	// get the request body
	var body config.RegisterRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		lib.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// validate the request body
	if body.Name == "" || body.Email == "" || body.Password == "" || body.Role == "" {
		lib.HandleResponse(w, http.StatusBadRequest, "All fields are required", nil)
		return
	}

	// validate email
	if !utils.ValidateEmail(body.Email) {
		lib.HandleResponse(w, http.StatusBadRequest, "Invalid email address", nil)
		return
	}

	// check if user already exists
	user := database.DBClient.Where("email = ?", body.Email).First(&models.User{})
	if user.Error == nil {
		lib.HandleResponse(w, http.StatusFound, "User already exists", nil)
		return
	}

	// hash the password
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		lib.HandleResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// create the user
	user = database.DBClient.Model(&models.User{}).Create(&models.User{
		Name: body.Name,
		Email: body.Email,
		Password: hashedPassword,
		Role: body.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if user.Error != nil {
		lib.HandleResponse(w, http.StatusInternalServerError, user.Error.Error(), nil)
		return
	}

	responseData := config.RegisterRequestBody{
		Name: body.Name,
		Email: body.Email,
		Role: body.Role,
	}
	lib.HandleResponse(w, http.StatusCreated, "User created successfully", responseData)
}

func UserLoginHandler(w http.ResponseWriter, r * http.Request) {}

func UserLogoutHandler(w http.ResponseWriter, r * http.Request) {}
