package lib

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/config"
)

func HandleResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(config.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
	if err != nil {
		log.Printf("Error encoding response: %s", err)
	}
}