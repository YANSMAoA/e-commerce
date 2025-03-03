package serversuite

import (
	"github.com/cloudwego/biz-demo/gomall/common/mtl"
	//"github.com/cloudwego/kitex/internal/mocks/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"

	//"github.com/hertz-contrib/obs-opentelemetry/provider"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr string
}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
		server.WithSuite(tracing.NewServerSuite()),
	}
	
	r, err := consul.NewConsulRegister(s.RegistryAddr)
	if err != nil {
		return nil
	}
	opts = append(opts, server.WithRegistry(r))

	// _ = provider.NewOpenTelemetryProvider(provider.WithSdkTracerProvider(mtl.TracerProvider), provider.WithEnableMetrics(false))
	//opts = append(opts, server.WithRegistry(r))
	// opts = append(opts,
	// 	server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
	// 		ServiceName: s.CurrentServiceName,
	// 	}),
	// 	server.WithSuite(tracing.NewServerSuite()),
	// 	server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
	// )

	return opts
}