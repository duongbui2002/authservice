package app

import (
	"github.com/duongbui2002/core-package/core/fxapp"
	"github.com/duongbui2002/core-package/core/fxapp/contracts"
)

type AuthApplicationBuilder struct {
	contracts.ApplicationBuilder
}

func NewAuthApplicationBuilder() *AuthApplicationBuilder {
	return &AuthApplicationBuilder{fxapp.NewApplicationBuilder()}

}

func (a *AuthApplicationBuilder) Build() *AuthApplication {
	return
}
