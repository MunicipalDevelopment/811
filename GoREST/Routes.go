package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"IndexPUT",
		"PUT",
		"/{id}",
		IndexPUT,
	},
	Route{
		"IndexGET",
		"GET",

		"/",
		IndexGET,
	},
	Route{
		"IndexPOST",
		"POST",
		"/",
		IndexPOST,
	},
	Route{
		"IndexDELETE",
		"DELETE",
		"/{id}",
		IndexDELETE,
	},
}
