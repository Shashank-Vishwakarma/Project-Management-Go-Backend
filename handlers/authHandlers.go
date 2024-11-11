package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/config"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/constants"
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

	responseData := config.UserData{
		Name: body.Name,
		Email: body.Email,
		Role: body.Role,
	}
	lib.HandleResponse(w, http.StatusCreated, "User created successfully", responseData)
}

func UserLoginHandler(w http.ResponseWriter, r * http.Request) {
	var body config.LoginRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if body.Email == "" || body.Password == "" || body.Role == "" {
		w.WriteHeader(http.StatusBadRequest)
		lib.HandleResponse(w, http.StatusBadRequest, "All fields are required", nil)
		return
	}

	// check if user exists
	user := &models.User{}
	result := database.DBClient.Where("email = ?", body.Email).First(&user)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		lib.HandleResponse(w, http.StatusNotFound, "User not found", nil)
		return
	}

	// check if password is correct
	if isPasswordMatch := utils.ComparePasswordHash(body.Password, user.Password); !isPasswordMatch {
		w.WriteHeader(http.StatusUnauthorized)
		lib.HandleResponse(w, http.StatusUnauthorized, "Invalid password", nil)
		return
	}

	if body.Role != user.Role {
		w.WriteHeader(http.StatusUnauthorized)
		lib.HandleResponse(w, http.StatusUnauthorized, "Invalid role", nil)
		return
	}

	// generate jwt token
	token, err := lib.GenerateJWT(user.ID, user.Name, user.Email, user.Role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.HandleResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// set cookie
	cookie := http.Cookie{
		Name: "token",
		Value: token,
		MaxAge: 60 * 60 * 24,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Path: constants.BASE_ENDPOINT,
	}
	http.SetCookie(w, &cookie)

	responseData := config.UserData{
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
		Token: token,
	}

	lib.HandleResponse(w, http.StatusOK, "Login successful...", responseData)
}

func UserLogoutHandler(w http.ResponseWriter, r * http.Request) {
	cookie := http.Cookie{
		Name: "token",
		Value: "",
		MaxAge: 0,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Path: constants.BASE_ENDPOINT,
	}
	http.SetCookie(w, &cookie)

	lib.HandleResponse(w, http.StatusOK, "Logout successful...", nil)
}
