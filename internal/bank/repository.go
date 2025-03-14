package bank

import (
	"context"
	"fmt"
	"log"

	"github.com/zuhdannur/go-fiber-bank-api/config"
	"github.com/zuhdannur/go-fiber-bank-api/internal/logger"
	"github.com/zuhdannur/go-fiber-bank-api/prisma/db"
)

type BankRepository interface {
	CreateAccount(ctx context.Context, account BankModel) (*db.BankModel, error)
	GetBankByAccountNumber(ctx context.Context, accountNumber string) (*db.BankModel, error)
	UpdateAccount(ctx context.Context, id string, bank db.InnerBank) (*db.BankModel, error)
}

type bankRepository struct{}

func NewBankRepository() BankRepository {
	return &bankRepository{}
}

var repositoryTag string = "BANK REPOSITORY"

func (r *bankRepository) CreateAccount(ctx context.Context, account BankModel) (*db.BankModel, error) {
	logger.Info(repositoryTag, "running CreateAccount Func")
	createdBank, err := config.DB.Bank.CreateOne(
		db.Bank.Name.Set(account.Name),
		db.Bank.CardID.Set(account.CardID),
		db.Bank.PhoneNumber.Set(account.PhoneNumber),
		db.Bank.Nominal.Set(account.Nominal),
		db.Bank.AccountNumber.Set(account.AccountNumber),
	).Exec(ctx)
	if err != nil {
		logger.Error(repositoryTag, "Failed to CreateAccount", err)
		return nil, err
	}
	logger.Info(repositoryTag, "Account record created successfully")
	return createdBank, nil
}

func (r *bankRepository) GetBankByAccountNumber(ctx context.Context, accountNumber string) (*db.BankModel, error) {
	bank, err := config.DB.Bank.FindFirst(
		db.Bank.AccountNumber.Equals(accountNumber),
	).Exec(ctx)
	if err != nil {
		logger.Warning(fmt.Sprintf("WARNING: Account record not found: %v", err))
		return nil, err
	}
	log.Println("INFO: Retrieved account record successfully")
	return bank, nil
}

func (r *bankRepository) UpdateAccount(ctx context.Context, id string, bank db.InnerBank) (*db.BankModel, error) {
	updatedBank, err := config.DB.Bank.FindUnique(
		db.Bank.ID.Equals(id),
	).Update(
		db.Bank.Name.Set(bank.Name),
		db.Bank.CardID.Set(bank.CardID),
		db.Bank.PhoneNumber.Set(bank.PhoneNumber),
		db.Bank.Nominal.Set(bank.Nominal),
		db.Bank.AccountNumber.Set(bank.AccountNumber),
	).Exec(ctx)
	if err != nil {
		logger.Error(repositoryTag, "Failed to UpdateAccount", err)
		return nil, err
	}
	logger.Info(repositoryTag, "Account record updated successfully")
	return updatedBank, nil
}
