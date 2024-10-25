package main

import (
	"encoding/json"
	"net/http"
)

func (a *applicationDependencies) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "avaliable",
		"system_info": map[string]string{
			"enviroment": a.config.enviroment,
			"version":    appVersion,
		},
	}
	err := a.writeJSON(w, http.StatusOK, data, nil)
	jsResponse, err := json.Marshal(data)
	if err != nil {
		//a.logger.Error(err.Error())
		//http.Error(w, "The server encountered a problem and could not prcess your request", http.StatusInternalServerError)
		a.serverErrorResponse(w, r, err)
		return
	}

	jsResponse = append(jsResponse, '\n')
	w.Header().Set("Content-Type", "application-json")
	w.Write(jsResponse)
}
