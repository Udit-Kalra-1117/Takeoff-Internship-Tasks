package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/uditkalra/swaggerRestApi/csv"
	"github.com/uditkalra/swaggerRestApi/functions"

	_ "github.com/uditkalra/swaggerRestApi/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

// @title Implementing and Documenting Employee Management System API in Go using Swagger
// @version 1
// @Description This is the implementation and documentation of the Employee Management System API in Go using Swagger

// @contact.name Udit Kalra
// @contact.url https://github.com/Udit-Kalra-1117
// @contact.email kalra.udit15@gmail.com

// @securityDefinitions.apikey bearerToken
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /api/v1
func main() {
	fmt.Println("\nWelcome to Employee Management System API Implementation and Documentation in GO using Swagger!\n")

	if _, err := os.Stat("employees.csv"); err == nil {
		csv.LoadFromCSV()
	} else {
		file, err := os.Create("employees.csv")
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	api := chi.NewRouter()
	api.Route("/api/v1", func(r chi.Router) {
		r.Mount("/employees", employeesRouter())
	})

	r.Mount("/", api)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
	))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

func employeesRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", functions.GetEmployees)
	r.Get("/{id}", functions.GetEmployeeByID)
	r.Post("/", functions.CreateEmployee)
	r.Put("/{id}", functions.UpdateEmployee)
	r.Delete("/{id}", functions.DeleteEmployee)

	return r
}