package controllers

import (
	"encoding/json"
	"guild/models"
	"net/http"

	"guild/events"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	toRegister := new(models.PatentToRegister)

	err := c.Bind(&toRegister)
	if err != nil {
		return c.String(http.StatusBadRequest, "You sent a bad request")
	}

	b, _ := json.Marshal(toRegister)
	err = events.RegisterEvent(b, 0)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server errror")
	}

	return c.String(http.StatusAccepted, "Registration sended")
}
