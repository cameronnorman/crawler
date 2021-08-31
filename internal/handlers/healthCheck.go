package handlers

import (
	"crawler/internal/openapi"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *manager) GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, openapi.HealthCheckResponse{Status: "RUNNING"})
}
