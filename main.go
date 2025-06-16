package main

import (
	"employee-management-app/config"
	"employee-management-app/handler"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config.LoadEnv()

	db := config.DBInit()
	defer db.Close()

	empHandler := handler.EmployeeHandler{DB: db}

	router := httprouter.New()
	router.GET("/employees", empHandler.GetAllEmployees)
	router.GET("/employees/:id", empHandler.GetEmployeeByID)
	router.POST("/employees", empHandler.CreateEmployee)
	router.PUT("/employees/:id", empHandler.UpdateEmployee)
	router.DELETE("/employees/:id", empHandler.DeleteEmployee)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
