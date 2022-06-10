package wallet

import (
	"errors"

	"github.com/ShavqatKavrakov/wallet/pkg/types"
	"github.com/google/uuid"
)

var ErrPhoneRegistres = errors.New("phone allready registered")
var ErrMustBePositive = errors.New("amount must be greater than 0")
var ErrAccountNotFound = errors.New("account not found")
var ErrNotEnoughBalance = errors.New("not enougth balance in account")

type Service struct {
	nextAccountID int64
	payments      []*types.Payment
	accounts      []*types.Account
}

func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, accounts := range s.accounts {
		if accounts.Phone == phone {
			return nil, ErrPhoneRegistres
		}
	}
	s.nextAccountID++
	account := &types.Account{
		ID:      s.nextAccountID,
		Phone:   phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)
	return account, nil
}

func (s *Service) Deposit(accountId int64, amount types.Money) error {
	if amount <= 0 {
		return ErrMustBePositive
	}

	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountId {
			account = acc
			break
		}
	}
	if account == nil {
		return ErrAccountNotFound
	}
	account.Balance += amount
	return nil
}
func (s *Service) Pay(accountID int64, amount types.Money, category types.PaymentCategory) (*types.Payment, error) {
	if amount <= 0 {
		return nil, ErrMustBePositive
	}
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
		}
	}
	if account == nil {
		return nil, ErrAccountNotFound
	}
	if account.Balance < amount {
		return nil, ErrNotEnoughBalance
	}
	account.Balance -= amount
	paymentID := uuid.New().String()
	payment := &types.Payment{
		ID:        paymentID,
		AccountID: accountID,
		Amount:    amount,
		Category:  category,
		Status:    types.PaymentStatusInProgress,
	}
	s.payments = append(s.payments, payment)
	return payment, nil
}
func (s *Service) FindAccontById(accountId int64) (*types.Account, error) {
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountId {
			account = acc
			break
		}
	}
	if account == nil {
		return nil, ErrAccountNotFound
	}
	return account, nil
}
