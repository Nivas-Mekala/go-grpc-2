package types

import (
	"context"

	"github.com/Nivas-Mekala/go-grpc-2/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
