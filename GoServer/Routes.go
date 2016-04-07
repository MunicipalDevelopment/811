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
		"POSTclose811",
		"POST",
		"/",
		POSTclose811,
	},
	Route{
		"IndexGET",
		"GET",
		"/",
		IndexGET,
	},
}
