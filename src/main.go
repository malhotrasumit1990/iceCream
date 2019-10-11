package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iceCream/src/dataAccess"
	"github.com/iceCream/src/web"
)

func init() {

	err := dataAccess.LoadDBConfig("DBConnectionStrings.json")
	if err != nil {
		log.Println(err)
	}

}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	router := web.SetupRoutes()

	err := http.ListenAndServe(":8888", router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

}
