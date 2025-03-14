package bank

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/zuhdannur/go-fiber-bank-api/internal/logger"
	"github.com/zuhdannur/go-fiber-bank-api/prisma/db"
	"github.com/zuhdannur/go-fiber-bank-api/utils"
)

type BankService struct {
	repo BankRepository // Use the interface, not the struct
}

func NewBankService(repo BankRepository) *BankService {
	return &BankService{repo: repo}
}

var serviceTag = "BankService"

func (s *BankService) CreateBank(ctx context.Context, input RegisterAccountModel) (*db.BankModel, error) {
	logger.Info(serviceTag, "running CreateBank Func")

	account := BankModel{
		Name:          input.Name,
		CardID:        input.CardID,
		PhoneNumber:   input.PhoneNumber,
		Nominal:       0,
		AccountNumber: utils.GenerateAccountNumber(12),
	}

	newAccount, err := s.repo.CreateAccount(
		ctx,
		account,
	)
	if err != nil {
		if strings.Contains(err.Error(), "Unique constraint failed") {
			if strings.Contains(err.Error(), "cardId") {
				return nil, fiber.NewError(fiber.StatusBadRequest, "NIK sudah terdaftar")
			}
			if strings.Contains(err.Error(), "phoneNumber") {
				return nil, fiber.NewError(fiber.StatusBadRequest, "Nomor telepon sudah digunakan")
			}
			return nil, fiber.NewError(fiber.StatusBadRequest, "Data sudah terdaftar")
		}
		return nil, err
	}
	return newAccount, err
}

func (s *BankService) Saving(ctx context.Context, input SavingModel) (*db.BankModel, error) {
	logger.Info(serviceTag, "running FindAccount Func")

	account, err := s.repo.GetBankByAccountNumber(
		ctx,
		input.AccountNumber,
	)
	if errors.Is(err, db.ErrNotFound) {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Nomor Rekening Tidak Ditemukan")
	}
	newNominal := account.Nominal + input.Nominal

	payload := db.InnerBank{
		Name:          account.Name,
		CardID:        account.CardID,
		PhoneNumber:   account.PhoneNumber,
		Nominal:       newNominal,
		AccountNumber: account.AccountNumber,
	}

	updatedAccount, err := s.repo.UpdateAccount(ctx, account.ID, payload)
	if err != nil {
		return nil, err
	}
	return updatedAccount, nil
}

func (s *BankService) Withdrawal(ctx context.Context, input WithdrawalModel) (*db.BankModel, error) {
	logger.Info(serviceTag, "running Withdrawal Func")

	account, err := s.repo.GetBankByAccountNumber(
		ctx,
		input.AccountNumber,
	)
	if errors.Is(err, db.ErrNotFound) {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Nomor Rekening Tidak Ditemukan")
	}

	if account.Nominal < input.Nominal {
		logger.Warning(serviceTag, "running Withdrawal Func saldo tidak cukup")
		return nil, fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Saldo tidak cukup. Saldo saat ini: %.0f", account.Nominal))
	}

	newNominal := account.Nominal - input.Nominal

	payload := db.InnerBank{
		Name:          account.Name,
		CardID:        account.CardID,
		PhoneNumber:   account.PhoneNumber,
		Nominal:       newNominal,
		AccountNumber: account.AccountNumber,
	}

	updatedAccount, err := s.repo.UpdateAccount(ctx, account.ID, payload)
	if err != nil {
		return nil, err
	}
	return updatedAccount, nil
}

func (s *BankService) GetBalance(ctx context.Context, accountNumber string) (*db.BankModel, error) {
	logger.Info(serviceTag, "running GetBalance Func")
	account, err := s.repo.GetBankByAccountNumber(
		ctx,
		accountNumber,
	)
	if errors.Is(err, db.ErrNotFound) {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Nomor Rekening Tidak Ditemukan")
	}

	return account, nil
}
