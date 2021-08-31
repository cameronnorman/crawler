//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -generate "types,server" -o openapi/openapi.gen.go --package openapi ../docs/swagger.yaml
package internal

import (
	"crawler/internal/handlers"
	"crawler/internal/openapi"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type App struct {
	Server *echo.Echo
}

func NewApp(db *gorm.DB, loggingEnabled bool) *App {
	app := &App{}

	e := echo.New()
	if loggingEnabled {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	handlerManager := handlers.NewManager(db)
	openapi.RegisterHandlers(e, handlerManager)

	app.Server = e

	return app

}

func (a *App) Run() error {
	return a.Server.Start(":3000")
}
