package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zuhdannur/go-fiber-bank-api/config"
	"github.com/zuhdannur/go-fiber-bank-api/internal/bank"
	"github.com/zuhdannur/go-fiber-bank-api/internal/logger"
)

func main() {
	config.InitDB()
	logger.InitLogger("logs/app.txt")

	app := fiber.New()

	bankRepo := bank.NewBankRepository()
	bankService := bank.NewBankService(bankRepo)
	bank.RegisterBankRoutes(app, bankService)

	logger.Info("server", "Starting server on port 3000")
	app.Listen(":3000")
}
