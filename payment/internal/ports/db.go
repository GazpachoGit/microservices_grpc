package ports

import "github.com/GazpachoGit/microservices/payment/internal/application/core/domain"

type DBport interface {
	Get(id string) (domain.Payment, error)
	//pointer to set ID coming from DB
	Save(payment *domain.Payment) error
}
