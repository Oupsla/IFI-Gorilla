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

func GetEnigma(res http.ResponseWriter, req *http.Request) {
    response, err := blackbox.ExecuteFunction("getFirstEnigme", req.Method, nil)

    if err != nil {
        myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
        utils.SendError(res, myError)
        return
    }

    utils.SendJSON(res, response)
}

type postParamsEnigma struct {
  Result int `json:"result"`
}

func PostEnigma(res http.ResponseWriter, req *http.Request) {
  params := &postParamsEnigma{}

  err := utils.GetJSONContent(params, req)

  if err != nil {
    myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
    utils.SendError(res, myError)
    return
  }

	response, err := blackbox.ExecuteFunction("submitFirstEnigme", req.Method, params.Result)

	if err != nil {
			myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
			utils.SendError(res, myError)
			return
	}

	utils.SendJSON(res, response)

}

func GetPlage(res http.ResponseWriter, req *http.Request) {
    response, err := blackbox.ExecuteFunction("getPlage", req.Method, nil)

    if err != nil {
        myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
        utils.SendError(res, myError)
        return
    }

    utils.SendJSON(res, response)
}

type postParamPlage struct {
  Result string `json:"result"`
}

func PostPlage(res http.ResponseWriter, req *http.Request) {
  params := &postParamPlage{}

  err := utils.GetJSONContent(params, req)

  if err != nil {
    myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
    utils.SendError(res, myError)
    return
  }

	response, err := blackbox.ExecuteFunction("addSomethingToTheBeach", req.Method, params.Result)

	if err != nil {
			myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
			utils.SendError(res, myError)
			return
	}

	utils.SendJSON(res, response)

}

func RemovePlage(res http.ResponseWriter, req *http.Request) {
  params := &postParamPlage{}

  err := utils.GetJSONContent(params, req)

  if err != nil {
    myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
    utils.SendError(res, myError)
    return
  }

	response, err := blackbox.ExecuteFunction("removeSomethingToTheBeach", req.Method, params.Result)

	if err != nil {
			myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
			utils.SendError(res, myError)
			return
	}

	utils.SendJSON(res, response)

}

func EditPlage(res http.ResponseWriter, req *http.Request) {
  params := &postParamPlage{}

  err := utils.GetJSONContent(params, req)

  if err != nil {
    myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
    utils.SendError(res, myError)
    return
  }

	response, err := blackbox.ExecuteFunction("editSomethingToTheBeach", req.Method, params.Result)

	if err != nil {
			myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
			utils.SendError(res, myError)
			return
	}

	utils.SendJSON(res, response)

}

func GetPresent(res http.ResponseWriter, req *http.Request) {
    response, err := blackbox.ExecuteFunction("getPresent", req.Method, nil)

    if err != nil {
        myError := utils.CustomError{StatusCode: 400, Message: err.Error()}
        utils.SendError(res, myError)
        return
    }

    utils.SendJSON(res, response)
}
