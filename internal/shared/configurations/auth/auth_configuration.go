package auth

import (
	"fmt"
	"github.com/duongbui2002/core-package/fxapp/contracts"
	echocontracts "github.com/duongbui2002/core-package/http/customecho/contracts"
	"github.com/duongbui2002/myblog-authservice/internal/auth/configurations"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthServiceConfigurator struct {
	contracts.Application
	authModuleConfigurator *configurations.AuthModuleConfigurator
}

func NewAuthServiceConfigurator(app contracts.Application) *AuthServiceConfigurator {
	return &AuthServiceConfigurator{
		Application:            app,
		authModuleConfigurator: configurations.NewAuthModuleConfigurator(app),
	}
}

func (ic *AuthServiceConfigurator) ConfigureAuth() {
	ic.authModuleConfigurator.ConfigureAuthModule()

}

func (ic *AuthServiceConfigurator) MapAuthEndpoints() {
	ic.ResolveFunc(func(authServer echocontracts.EchoHttpServer, cfg *config.Config) {
		authServer.SetupDefaultMiddlewares()

		authServer.RouteBuilder().RegisterRoutes(func(e *echo.Echo) {
			e.GET("", func(ec echo.Context) error {
				return ec.String(
					http.StatusOK,
					fmt.Sprintf(
						"%s is running...",
						cfg.AppOptions.GetMicroserviceNameUpper(),
					),
				)
			})
		})
	})

	ic.authModuleConfigurator.MapAuthEndpoints()

}
