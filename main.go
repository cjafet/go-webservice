// package main specifies an entry point for your application
package main

import (
	"github.com/cjafet/go-webservice/controllers"
	"net/http"
)

func main() {
	/*
		in the same way that we need to prefix Println with the fmt package so that we can use it we also need to
		prefix User with its package name, that in this case is the models.
		Note the import is using the fully qualified package name: github.com/cjafet/go-webservice/models

	*/

	//u := models.User{
	//	ID:        2,
	//	FirstName: "Carlos",
	//	LastName:  "Jafet",
	//}
	//fmt.Println(u)

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
