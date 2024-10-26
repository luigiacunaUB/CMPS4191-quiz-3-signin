package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/luigiacunaUB/CMPS4191-quiz-3-signin/internal/data"
	"github.com/luigiacunaUB/CMPS4191-quiz-3-signin/internal/validator"
)

func (a *applicationDependencies) createSignInHandler(w http.ResponseWriter, r *http.Request) {
	a.logger.Info("starting to create a signin")
	//creating struct to hold signin info
	var incomingData struct {
		Email    string `json:"email"`
		FullName string `json:"fullname"`
	}

	//decoding done here
	err := a.readJSON(w, r, &incomingData)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}
	signin := &data.SignIN{
		Email:    incomingData.Email,
		FullName: incomingData.FullName,
	}

	v := validator.New()

	data.ValidateSignin(v, signin)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = a.SignINModel.Insert(signin)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	//display result
	//fmt.Fprintf(w, "%+v\n", incomingData)
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/view-sign-in/%d", signin.ID))

	data := envelope{
		"signin": signin,
	}
	err = a.writeJSON(w, http.StatusCreated, data, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}

func (a *applicationDependencies) displaySignInHandler(w http.ResponseWriter, r *http.Request) {
	a.logger.Info("Inside displaySignInHandler")
	id, err := a.ReadIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}
	a.logger.Info("displaySignInHandler", "id", id)
	signin, err := a.SignINModel.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}
	data := envelope{
		"signin": signin,
	}
	err = a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}
