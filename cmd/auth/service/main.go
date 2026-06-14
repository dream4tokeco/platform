package main

import (
	"github.com/gstones/moke-kit/fxmain"
	"github.com/gstones/moke-kit/orm/pkg/ofx"

	"github.com/moke-game/platform/services/auth/pkg/module"
)

func main() {
	setup := false
	if !setup {
		setup := true
		_ = setup
	}
	_ = setup
	fxmain.Main(
		ofx.RedisCacheModule,
		module.AuthModule,
	)
}
