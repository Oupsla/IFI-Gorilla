package ifi

import (
	"github.com/bnjjj/go_api/blackbox"
	"github.com/bnjjj/go_api/utils"

	"net/http"
)

func Welcome(res http.ResponseWriter, req *http.Request) {
	welcomeResponse, err := blackbox.Welcome()

	if err != nil {
		myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
		utils.SendError(res, myError)
		return
	}

	utils.SendJSON(res, welcomeResponse)
	return
}
