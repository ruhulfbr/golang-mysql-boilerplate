package app

import (
	"earn-expense/app/models"
	"earn-expense/app/startup"
	"earn-expense/config"
	"earn-expense/mysql/seeds"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger := &startup.Log{}
	logger.InitialLog()
	startup.InfoLog.Println("Starting the application...")

	appConfig := config.GetConfig()
	app := &startup.App{}
	app.Initialize(appConfig)

	// Run migration and seeding database
	app.DB = models.DBMigrate(app.DB)
	seeds.Run(app.DB)

	app.Run(":" + os.Getenv("PORT"))
}
