package bank

import (
	"context"

	"github.com/zuhdannur/go-fiber-bank-api/internal/logger"
	"github.com/zuhdannur/go-fiber-bank-api/utils"

	"github.com/gofiber/fiber/v2"
)

type BankResponse struct {
	Name        string  `json:"name"`
	CardID      string  `json:"card_id"`
	PhoneNumber string  `json:"phone_number"`
	Nominals    float64 `json:"nominals"`
	CheckNumber string  `json:"check_number"`
}

var routerTag = "BankHandler"

func RegisterBankRoutes(app *fiber.App, service *BankService) {

	app.Post("/daftar", func(c *fiber.Ctx) error {

		logger.Info(routerTag, "running handler /daftar")

		var input RegisterAccountModel
		if err := c.BodyParser(&input); err != nil {
			logger.Error(routerTag, "Invalid request /daftar", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		account, err := service.CreateBank(context.Background(), input)
		if err != nil {
			logger.Error(routerTag, "failed to /daftar", err)
			return utils.ErrorResponse(c, fiber.ErrBadRequest.Code, err.Error())
		}

		return utils.SuccessResponse(c, fiber.StatusOK, "Pendaftaran akun berhasil", fiber.Map{
			"no_rekening": account.AccountNumber,
		})
	})

	app.Post("/tabung", func(c *fiber.Ctx) error {
		logger.Info(routerTag, "running Handler /tabung")
		var input SavingModel
		if err := c.BodyParser(&input); err != nil {
			logger.Error(routerTag, "Invalid request /tabung", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		account, err := service.Saving(context.Background(), input)
		if err != nil {
			logger.Error(routerTag, "Failed to tabung", err)
			return utils.ErrorResponse(c, fiber.ErrBadRequest.Code, err.Error())
		}

		return utils.SuccessResponse(c, fiber.StatusOK, "Saldo telah ditambahkan.", fiber.Map{
			"saldo": account.Nominal,
		})
	})

	app.Post("/tarik", func(c *fiber.Ctx) error {
		logger.Info(routerTag, "running handler /tarik")
		var input WithdrawalModel
		if err := c.BodyParser(&input); err != nil {
			logger.Error(routerTag, "Invalid request /tarik", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}

		account, err := service.Withdrawal(context.Background(), input)
		if err != nil {
			logger.Error(routerTag, "Failed to tarik", err)
			return utils.ErrorResponse(c, fiber.ErrBadRequest.Code, err.Error())
		}

		return utils.SuccessResponse(c, fiber.StatusOK, "Dana telah ditarik.", fiber.Map{
			"saldo": account.Nominal,
		})
	})

	app.Get("/saldo/:no_rekening", func(c *fiber.Ctx) error {
		logger.Info(routerTag, "running Handler /tarik")

		accountNumber := c.Params("no_rekening")

		account, err := service.GetBalance(context.Background(), accountNumber)
		if err != nil {
			logger.Error(routerTag, "Failed to check saldo", err)
			return utils.ErrorResponse(c, fiber.ErrBadRequest.Code, err.Error())
		}

		return utils.SuccessResponse(c, fiber.StatusOK, "Berhasil mendapatkan informasi saldo.", fiber.Map{
			"saldo": account.Nominal,
		})
	})

}
