package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"GoNext/base/cmd/http/srvinit"
	"GoNext/base/ent"
	"GoNext/base/pkg/config"
	"GoNext/base/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

func main() {
	config := config.LoadConfig()

	var entClient *ent.Client
	if config.Db.Type == "sqlite" {
		log.Println("using sqlite")
		entClient = database.NewSQLiteEntClient(config)
	} else {
		entClient = database.NewEntClient(config)
	}
	defer entClient.Close()

	repos := srvinit.InitRepos(entClient)
	services := srvinit.InitServices(repos, config)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.Cors.AllowOrigins,
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		AllowCredentials: true,
	}))

	srvinit.InitHandlers(app, services, config)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-shutdown
		log.Println("Shutting down...")
		app.Shutdown()
		os.Exit(0)
	}()

	log.Fatal(app.Listen(":8080"))
}
