package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *manager) IndexCrawlRequests(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "")
}

func (m *manager) CreateCrawlRequest(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "")
}

func (m *manager) GetApiV1CrawlRequestsUuid(c echo.Context, uuid string) error {
	return c.JSON(http.StatusNotImplemented, "")
}
