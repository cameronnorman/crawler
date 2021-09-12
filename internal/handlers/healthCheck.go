package handlers

import (
	"crawler/internal/openapi"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (m *manager) GetApiV1Health(c echo.Context) error {
	name, _ := os.Hostname()
	return c.JSON(http.StatusOK, openapi.HealthCheckResponse{Status: fmt.Sprintf("RUNNING: %s", name)})
}
