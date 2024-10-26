package main

import (
	"fmt"
	"net/http"

	"github.com/luigiacunaUB/CMPS4191-quiz-3-signin/internal/validator"

	"github.com/luigiacunaUB/CMPS4191-quiz-3-signin/internal/data"
)

func (a *applicationDependencies) createSignInHandler(w http.ResponseWriter, r *http.Request) {
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
	headers.Set("Location", fmt.Sprintf("/signin/%d", signin.ID))

	data := envelope{
		"signin": signin,
	}
	err = a.writeJSON(w, http.StatusCreated, data, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}
