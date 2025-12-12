package main

import (
	"golang-file-processing-api/internal/config"
	"log"
)

// @title File Processing API
// @version 1.0
// @description API untuk upload file dan insert ke Database
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
func main() {
	cfg := config.Load()

	log.Println("Server starting on port:", cfg.Port)

}
