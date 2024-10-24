//routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *applicationDependencies) routes() http.Handler {

	//setup a new router
	router := httprouter.New()

	//handle 404
	//router.NotFound = http.HandlerFunc(a.notFoundResponse)

	//handle 405
	//router.MethodNotAllowed = http.HandlerFunc(a.MethodNotAllowed)

	//routes
	router.HandlerFunc(http.MethodGet, "/", a.Index)

	return router
}
