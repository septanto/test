package facade

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"septanto.github.com/omdb/service"
)

func Search(c echo.Context) error {
	pagination := c.QueryParam("pagination")
	searchword := c.QueryParam("searchword")
	if pagination != "" && searchword != "" {
		resp, err := service.RequestSearchToOmdb(searchword, pagination)
		if err != nil {
			log.Info("Error request to omdb: " + err.Error())
			return c.String(http.StatusBadGateway, "Bad Gateway")
		}
		if resp.StatusCode() != http.StatusOK {
			return c.String(resp.StatusCode(), string(resp.Body()))
		}
		return c.String(http.StatusOK, string(resp.Body()))
	} else {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
}

func GetDetail(c echo.Context) error {
	id := c.QueryParam("imdbid")
	if id != "" {
		resp, err := service.RequestDetailToOmdb(id)
		if err != nil {
			log.Info("Error request to omdb: " + err.Error())
			return c.String(http.StatusBadGateway, "Bad Gateway")
		}
		if resp.StatusCode() != http.StatusOK {
			return c.String(http.StatusBadGateway, "Bad Gateway")
		}
		return c.String(http.StatusOK, string(resp.Body()))
	} else {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
}
