package main

import (
	"ims-be/router"
	"log"
	"net/http"
	"os"
)

func init() {
	ErrLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	// loadEnvParameters()
}

var ErrLog *log.Logger

// @title Swagger scheduling-api project API
// @version 2.0
// @description This is list api for scheduling-api project

// @host localhost:3000
// @BasePath /api/tda/student-management

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	if err := http.ListenAndServe(":3000", router.Router()); err != nil {
		ErrLog.Printf("Startup Error! %v", err)
	}
}
