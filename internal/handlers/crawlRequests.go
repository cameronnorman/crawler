package handlers

import (
	"net/http"

	faktory "github.com/contribsys/faktory/client"
	"github.com/labstack/echo/v4"
)

func (m *manager) IndexCrawlRequests(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "")
}

func (m *manager) CreateCrawlRequest(c echo.Context) error {
	cl, err := faktory.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	request := map[string]string{}
	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	job := faktory.NewJob("get_html_worker", request["website"])
	job.ReserveFor = 1
	err = cl.Push(job)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusCreated, "")
}

func (m *manager) GetApiV1CrawlRequestsUuid(c echo.Context, uuid string) error {
	return c.JSON(http.StatusNotImplemented, "")
}
