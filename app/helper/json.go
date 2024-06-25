package helper

import (
	"encoding/json"
	"net/http"
)

// write json response using map
// helper.JSON(w, http.StatusOK, map[string]string{"message": "pong"})
func JSON(w http.ResponseWriter, code int, message map[string]interface{}) {
	json.NewEncoder(w).Encode(message)
}
