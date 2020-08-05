package usecase

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/accounts"
	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/accounts/entities"
)

type Accounts struct {
	log        *logrus.Logger
	repository accounts.Repository
}

func NewAccountUseCase(log *logrus.Logger, repository accounts.Repository) *Accounts {
	return &Accounts{
		log:        log,
		repository: repository,
	}
}

func (a Accounts) CreateAccount(input accounts.AccountInput) error {
	if input.Type == "" {
		return errors.New("missing input type")
	}
	if input.Owner == "" {
		return errors.New("missing input owner")
	}
	if input.Name == "" {
		return errors.New("missing input name")
	}

	accountType := entities.AccountType(input.Type)
	if accountType != entities.Asset && accountType != entities.Liability {
		return fmt.Errorf("unknown account type '%s'", input.Type)
	}

	account := entities.Account{
		Type:     accountType,
		Owner:    input.Owner,
		Name:     input.Name,
		OwnerID:  input.OwnerID,
		Metadata: input.Metadata,
		Balance:  0,
	}

	if err := a.repository.Create(&account); err != nil {
		return fmt.Errorf("can't create account: %s", err.Error())
	}

	return nil
}