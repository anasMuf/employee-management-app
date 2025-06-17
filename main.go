package main

import (
	"employee-management-app/config"
	"employee-management-app/handler"
	"employee-management-app/repositories"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config.LoadEnv()

	db := config.DBInit()
	defer db.Close()

	empRepo := repositories.NewEmployeeRepository(db)
	empHandler := handler.NewEmployeeHandler(empRepo)

	router := httprouter.New()
	router.GET("/employees", empHandler.GetAllEmployees)
	router.GET("/employees/:id", empHandler.GetEmployeeByID)
	router.POST("/employees", empHandler.CreateEmployee)
	router.PUT("/employees/:id", empHandler.UpdateEmployee)
	router.DELETE("/employees/:id", empHandler.DeleteEmployee)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }
	// err := http.ListenAndServe(":"+port, router)
	// if err != nil {
	// 	log.Fatal("server gagal dijalankan", err)
	// }
}
