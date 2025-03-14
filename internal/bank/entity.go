package bank

type BankModel struct {
	Name          string  `json:"nama" validate:"required"`
	CardID        string  `json:"nik" validate:"required"`
	PhoneNumber   string  `json:"no_hp" validate:"required"`
	Nominal       float64 `json:"nominal" validate:"required"`
	AccountNumber string  `json:"account_number" validate:"required"`
}

type RegisterAccountModel struct {
	Name        string `json:"nama" validate:"required"`
	CardID      string `json:"nik" validate:"required"`
	PhoneNumber string `json:"no_hp" validate:"required"`
}

type SavingModel struct {
	AccountNumber string  `json:"no_rekening" validate:"required"`
	Nominal       float64 `json:"nominal" validate:"required,numeric"`
}

type WithdrawalModel struct {
	AccountNumber string  `json:"no_rekening" validate:"required"`
	Nominal       float64 `json:"nominal" validate:"required,numeric"`
}
