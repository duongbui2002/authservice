package configurations

import (
	"github.com/duongbui2002/core-package/core/web/route"
	contracts2 "github.com/duongbui2002/core-package/fxapp/contracts"
	echocontracts "github.com/duongbui2002/core-package/http/customecho/contracts"
	"github.com/duongbui2002/core-package/logger"
)

type AuthModuleConfigurator struct {
	contracts2.Application
}

func NewAuthModuleConfigurator(app contracts2.Application) *AuthModuleConfigurator {
	return &AuthModuleConfigurator{app}
}

func (c *AuthModuleConfigurator) ConfigureAuthModule() {
	c.ResolveFunc(
		func(logger logger.Logger,
			server echocontracts.EchoHttpServer,

		) error {
			return nil
		})
}

func (c *AuthModuleConfigurator) MapAuthEndpoints() {
	c.ResolveFuncWithParamTag(func(endpoints []route.Endpoint) {
		for _, endpoint := range endpoints {
			endpoint.MapEndpoint()
		}
	}, `group:"auth-routes"`)
}
