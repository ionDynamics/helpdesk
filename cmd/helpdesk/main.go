package main

import (
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/labstack/echo"
	"github.com/thoas/stats"
	"go.iondynamics.net/iDechoLog"
	idl "go.iondynamics.net/iDlogger"

	"go.iondynamics.net/helpdesk/template"
)

func main() {
	idl.Info("Starting...")

	e := echo.New()

	tpl, err := template.New()
	if err != nil {
		idl.Emerg(err)
	}
	e.SetRenderer(tpl)

	e.Use(iDechoLog.New())

	s := stats.New()
	e.Use(s.Handler)
	e.Get("/stats", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, s.Data())
	})

	e.Use(nosurf.NewPure)

	e.Get("/login", func(c *echo.Context) error {
		return c.Render(http.StatusOK, "loginGet.tpl", nil)
	})

	e.Run(":3001")
}
