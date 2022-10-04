package main

import (
	"fmt"
	"net/http"

	"github.com/danakin/festor.info/cmd/config"
	"github.com/danakin/festor.info/cmd/routes"
)

func main() {
	app, err := config.NewApplication("./ui/templates")
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%v", app)

	routes := routes.MakeRoutes(app)

	http.ListenAndServe(":3000", routes)
}
