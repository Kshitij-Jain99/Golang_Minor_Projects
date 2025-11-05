package main

import (
	"fmt"
	"go-fiber-CRM/database"
	"go-fiber-CRM/lead"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// 👇 add this line for pure Go SQLite
	_ "modernc.org/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	// var err error
	// database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// fmt.Println("Connection opened to database")
	// database.DBConn.AutoMigrate(&lead.Lead{})
	// fmt.Println("Database Migrated")

	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("❌ Failed to connect database: %v\n", err)
		panic(err)
	}
	fmt.Println("✅ Connection opened to database")

	err = database.DBConn.AutoMigrate(&lead.Lead{})
	if err != nil {
		panic(fmt.Sprintf("❌ AutoMigrate failed: %v", err))
	}
	fmt.Println("✅ Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	sqlDB, err := database.DBConn.DB()
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	defer sqlDB.Close()
	app.Listen(":3000")

}
