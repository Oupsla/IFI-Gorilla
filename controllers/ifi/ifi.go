package ifi

import (
	"github.com/oupsla/IFI-Gorilla/blackbox"
	"github.com/oupsla/IFI-Gorilla/utils"

	"net/http"
)

func Welcome(res http.ResponseWriter, req *http.Request) {
	welcomeResponse, err := blackbox.ExecuteFunction("welcome", req.Method, nil)

	if err != nil {
		myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
		utils.SendError(res, myError)
		return
	}

	utils.SendJSON(res, welcomeResponse)
}
