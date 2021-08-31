package handlers

import (
	"crawler/internal/openapi"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *manager) GetApiV1Crawls(c echo.Context, params openapi.GetApiV1CrawlsParams) error {
	return c.JSON(http.StatusNotImplemented, "")
}

func (m *manager) GetApiV1CrawlsUuid(c echo.Context, uuid string) error {
	return c.JSON(http.StatusNotImplemented, "")
}
