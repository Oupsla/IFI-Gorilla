package utils

import (
	"github.com/gorilla/mux"

	"encoding/json"
	"net/http"
	"os/exec"
)

type CustomError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func SendJSON(res http.ResponseWriter, data interface{}) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(200)

	err := json.NewEncoder(res).Encode(data)
	if err != nil {
		panic(err)
	}
}

func SendError(res http.ResponseWriter, errorInfos CustomError) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(errorInfos.StatusCode)

	err := json.NewEncoder(res).Encode(errorInfos)
	if err != nil {
		panic(err)
	}
}

func MakeCommand(name string, args []string) (string, error) {
	customCmd := exec.Command(name, args...)
	response, err := customCmd.Output()

	return string(response), err
}

// GetJSONContent returns the JSON content of a request
func GetJSONContent(v interface{}, r *http.Request) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

// ParamAsString returns an URL parameter /{name} as a string
func ParamAsString(name string, r *http.Request) string {
	vars := mux.Vars(r)
	value := vars[name]
	return value
}
