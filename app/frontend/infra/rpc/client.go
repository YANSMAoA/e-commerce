package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/conf"
	frontendutils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
	ServiceName = frontendutils.ServiceName
	RegistryAddr = conf.GetConf().Hertz.RegistryAddr
	err error
)

func InitClient() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr: RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr: RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr: RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr: RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonGrpcClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr: RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}
