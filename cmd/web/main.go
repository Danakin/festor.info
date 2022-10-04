package main

import (
	"net/http"

	"github.com/danakin/festor.info/cmd/routes"
)

func main() {
	routes := routes.MakeRoutes()

	http.ListenAndServe(":3000", routes)
}
