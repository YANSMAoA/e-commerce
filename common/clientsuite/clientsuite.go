package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	//"github.com/cloudwego/kitex/internal/server"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"

	//"github.com/cloudwego/kitex/server"

	//"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"
	//"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	// "github.com/hertz-contrib/obs-opentelemetry/provider"
	// "github.com/kitex-contrib/obs-opentelemetry/tracing"
	//tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

type CommonGrpcClientSuite struct {
	// DestServiceName    string
	// DestServiceAddr    string
	CurrentServiceName string
	RegistryAddr string
	// TracerProvider     *tracesdk.TracerProvider
}

func (s CommonGrpcClientSuite) Options() []client.Option {
	opts := []client.Option{
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
		client.WithSuite(tracing.NewClientSuite()),
	}
	
	r, err := consul.NewConsulResolver(s.RegistryAddr)
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	// _ = provider.NewOpenTelemetryProvider(provider.WithSdkTracerProvider(s.TracerProvider), provider.WithEnableMetrics(false))

	// opts = append(opts,
	// 	client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
	// 		ServiceName: s.CurrentServiceName,
	// 	}),
	// 	client.WithSuite(tracing.NewClientSuite()),
	// )

	return opts
}