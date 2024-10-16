package api

import (
	"github.com/GazpachoGit/microservices/payment/internal/application/core/domain"
	"github.com/GazpachoGit/microservices/payment/internal/ports"
)

type Application struct {
	db ports.DBport
}

func NewApplication(db ports.DBport) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) Charge(payment domain.Payment) (domain.Payment, error) {
	err := a.db.Save(&payment)

	if err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}
