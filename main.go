package main

import (
	//"fmt"
	"net/http"

	"github.com/pluralsight/webservice/controllers"
	//"github.com/pluralsight/webservice/models"
)

func main() {
	//fmt.Println("Hello from a modules")
	/*
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}

	fmt.Println(u)*/
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
