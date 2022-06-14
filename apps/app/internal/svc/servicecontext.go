package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mall-go-zero-learn.com/apps/app/internal/config"
	"mall-go-zero-learn.com/apps/order/rpc/order"
	"mall-go-zero-learn.com/apps/product/rpc/product"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   order.Order
	ProductRPC product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
