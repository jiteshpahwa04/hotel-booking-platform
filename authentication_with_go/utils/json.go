package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = NewValidator()
}

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled()) // struct level validation -> 'required', 'email', 'min' etc. in dto
}

func WriteJSONResponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func WriteJSONSuccessResponse(w http.ResponseWriter, status int, message string, data any) error {
	response := map[string]any{}
	response["message"] = message
	response["data"] = data
	response["status"] = "success"

	return WriteJSONResponse(w, status, response)
}

func WriteJSONErrorResponse(w http.ResponseWriter, status int, message string, err error) error {
	response := map[string]any{}
	response["message"] = message
	response["data"] = nil
	response["status"] = "error"
	response["error"] = err.Error()

	return WriteJSONResponse(w, status, response)
}

func ReadJSONRequest(r *http.Request, result any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Optional: disallow unknown fields
	return decoder.Decode(result)
}