package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echo"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal()
	}
	c, err := echo.NewClient("demo_proto", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
	res, err := c.Echo(context.TODO(), &pbapi.Request{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", res)
}