package main

import (
	"fmt"
	"net/http"
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

	//display result
	fmt.Fprintf(w, "%+v\n", incomingData)
}
