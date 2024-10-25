package main

import (
	"encoding/json"
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
	err := json.NewDecoder(r.Body).Decode(&incomingData)
	if err != nil {
		a.errorResponseJSON(w, r, http.StatusBadRequest, err.Error())
		return
	}

	//display result
	fmt.Fprintf(w, "%+v\n", incomingData)
}
