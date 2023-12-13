package resty

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/config"

	"github.com/dubonzi/otelresty"
	"github.com/go-resty/resty/v2"
)

var Client *resty.Client

func init() {
	config.Installations.EnableRestyClient = true
	LoadDependency()
}

func LoadDependency() container.LoadDependency {
	var dependency container.LoadDependency = func() error {
		Client = resty.New()
		opts := []otelresty.Option{otelresty.WithTracerName("resty-http-client")}
		otelresty.TraceClient(Client, opts...)
		return nil
	}
	container.InjectInstallation(dependency)
	return dependency
}
