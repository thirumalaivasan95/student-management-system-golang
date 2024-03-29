package main

import (
	"net/http"
	"time"

	"github.com/thirumalaivasan95/student-management-system-golang/models"
	"github.com/thirumalaivasan95/student-management-system-golang/routes"
)

func main() {
	// connect to MYSQL
	models.InitDB()

	// Listen and Server in 0.0.0.0:8080
	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	routes.SetupRouter(s)
}
