package app

import (
	"github.com/duongbui2002/core-package/core/config/environment"
	"github.com/duongbui2002/core-package/core/fxapp"
	"github.com/duongbui2002/core-package/core/logger"
	"go.uber.org/fx"
)

type AuthApplication struct {
}

func NewAuthApplication(
	providers []interface{},
	decorates []interface{},
	options []fx.Option,
	logger logger.Logger,
	environment environment.Environment,
) *AuthApplication {

	app := fxapp.NewApplication(providers, decorates, options, logger, environment)
	print("app: ", app)
	return &AuthApplication{}
}