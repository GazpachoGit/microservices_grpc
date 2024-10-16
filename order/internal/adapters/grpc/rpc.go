package grpc

import (
	"context"

	"github.com/GazpachoGit/microservices/order/internal/application/core/domain"
	"github.com/GazpachoGit/microservices_proto/golang/order"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var validationErrors []*errdetails.BadRequest_FieldViolation
	if request.UserId < 1 {
		validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field:       "user_id",
			Description: "user id cannot be less than 1",
		})
	}
	// if request.ProductId < 0 {
	// 	validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
	// 		Field:       "product_id",
	// 		Description: "product id cannot be negative",
	// 	})
	// }
	if len(validationErrors) > 0 {
		stat := status.New(400, "invalid order request")
		badRequest := &errdetails.BadRequest{}
		badRequest.FieldViolations = validationErrors
		s, _ := stat.WithDetails(badRequest)
		return nil, s.Err()
	}
	var orderItems []domain.OrderItem
	for _, orderItem := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}
