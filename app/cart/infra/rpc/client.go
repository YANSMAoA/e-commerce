package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/cart/conf"
	cartutils "github.com/cloudwego/biz-demo/gomall/app/cart/utils"
	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"
	//"github.com/cloudwego/biz-demo/gomall/common/serversuite"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	//consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	ServiceName = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
	err error
)


func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option {
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr: RegistryAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}
