package main

import (
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/labstack/echo"
	"github.com/thoas/stats"
	"go.iondynamics.net/iDechoLog"
	idl "go.iondynamics.net/iDlogger"

	"go.iondynamics.net/helpdesk/persistence"
	"go.iondynamics.net/helpdesk/persistence/backend"
	"go.iondynamics.net/helpdesk/route"
	"go.iondynamics.net/helpdesk/template"
)

func main() {
	idl.Info("Starting...")

	pp, err := backend.InitBolt("helpdesk.db")
	if err != nil {
		idl.Emerg(err)
	}
	persistence.Init(pp)

	tpl, err := template.New()
	if err != nil {
		idl.Emerg(err)
	}

	e := echo.New()
	e.SetRenderer(tpl)
	e.Use(iDechoLog.New())
	e.Use(nosurf.NewPure)

	s := stats.New()
	e.Use(s.Handler)
	e.Get("/stats", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, s.Data())
	})

	route.Init(e)

	e.Run(":3001")

}
