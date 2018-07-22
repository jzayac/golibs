package controllerHelper

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "github.com/dgrijalva/jwt-go"
)

type libVersion struct {
	LibVersion string `json:"libVersion"`
}

type ErrorHttp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	// Data    map[string]interface{} `json:"data"`
}

var (
	// ErrorHttp internal server error
	ErrInternalServerErr = &ErrorHttp{
		Code:    500,
		Message: "An error occurred on the server. Please try again later.",
	}

	ErrUnauthorized = &ErrorHttp{
		Code:    401,
		Message: "Unauthorized",
	}

	ErrForbidden = &ErrorHttp{
		Code:    403,
		Message: "Forbidden",
	}

	ErrNotFound = &ErrorHttp{
		Code:    404,
		Message: "not found",
	}
)

type StatusOk struct {
	Status string `json:"status"`
}

func JsonOkResponse(w http.ResponseWriter) {
	ok := StatusOk{
		Status: "OK",
	}
	json, err := json.Marshal(ok)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// w.Write(json)
	fmt.Fprintf(w, "%s", json)
}

func JsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// w.Write(json)
	fmt.Fprintf(w, "%s", json)
}

func ErrorResponse(errResp *ErrorHttp, w http.ResponseWriter) {
	// json, err := json.Marshal(errResp.Message)
	json, err := json.Marshal(errResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errResp.Code)
	fmt.Fprintf(w, "%s", json)
}

// type RestApiConfig = map[string]string

func LibraryVersion(w http.ResponseWriter, r *http.Request) {
	idx := libVersion{
		LibVersion: "0.0.1",
	}

	resp, _ := json.Marshal(idx)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", resp)

	// jsonResponse(&config, w) }
}
