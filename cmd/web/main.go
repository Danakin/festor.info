package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/danakin/festor.info/cmd/config"
	"github.com/danakin/festor.info/cmd/database/psql"
	"github.com/danakin/festor.info/cmd/routes"
)

func main() {
	host := flag.String("host", "localhost", "The Host used to connect to the DB")
	port := flag.String("port", "5432", "The Port used to connect to the DB")
	user := flag.String("user", "user", "The User used to connect to the DB")
	password := flag.String("password", "password", "The Password used to connect to the DB")
	database := flag.String("database", "database", "The Database used to connect to the DB")
	sslmode := flag.String("sslmode", "disable", "The SSLMode used to connect to the DB")
	flag.Parse()

	dbCfg := &psql.Config{
		Host:     *host,
		Port:     *port,
		User:     *user,
		Password: *password,
		Database: *database,
		SSLMode:  *sslmode,
	}
	db, err := psql.Connect(*dbCfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app, err := config.NewApplication(db)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%v", app)

	routes := routes.MakeRoutes(app)

	http.ListenAndServe(":3000", routes)
}
