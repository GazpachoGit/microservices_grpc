package ports

import "github.com/GazpachoGit/microservices/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	//pointer to set ID coming from DB
	Save(*domain.Order) error
}
