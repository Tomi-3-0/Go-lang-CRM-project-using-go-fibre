package main

import (
	"github.com/Tomi-3-0/go-basic-crm-fiber/database"
	"github.com/Tomi-3-0/go-basic-crm-fiber/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initdatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect to database")
	} 
	fmt.Println("Connection to database established")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrate")
}
func main () {
	app:= fiber.New()
	initdatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}