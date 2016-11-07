package controllers

import (
	//"github.com/bnjjj/go_api/models"
	"github.com/bnjjj/go_api/utils"

	"net/http"
	"strings"
)

func GetAll(res http.ResponseWriter, req *http.Request) {
	//test := models.User{ID: "0D9SNJSD", Firstname: "Benjamin", Lastname: "Coenen"}

	// myError := utils.CustomError{StatusCode: 404, Message: "user '" + test.Firstname + "' not found"}
	// utils.SendError(res, myError)
	info, err := utils.MakeCommand("ls", []string{"-l", "-a"})

	if err != nil {
		myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
		utils.SendError(res, myError)
		return
	}

	utils.SendJSON(res, strings.Split(info, "\n"))
}
