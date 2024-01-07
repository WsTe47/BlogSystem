package main

import (
	"BlogSystem/internal/app"
	"BlogSystem/internal/infra/impl"
	"BlogSystem/internal/infra/server"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	app.NewUser,
	impl.NewUserImpl,
	server.NewHTTPServer,
)
