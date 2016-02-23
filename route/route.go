package route //import "go.iondynamics.net/helpdesk/route"

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	idl "go.iondynamics.net/iDlogger"

	"go.iondynamics.net/helpdesk/controller"
	"go.iondynamics.net/helpdesk/helper"
	"go.iondynamics.net/helpdesk/persistence"
	"go.iondynamics.net/helpdesk/typ"
)

func Init(e *echo.Echo) {
	e.Get("/login", func(c *echo.Context) error {
		return c.Render(http.StatusOK, "loginGet.tpl", nil)
	})
}

func InitAPI(e *echo.Echo) {
	InitAPIv1(e.Group("/api"))
}

func getUser(c *echo.Context) *typ.User {
	usr := c.Get("api_usr")
	if usr == nil || usr.(type) != *typ.User {
		if usr.(type) != *typ.User {
			idl.Err("context.Get(\"api_usr\") is not a *typ.User")
		}
		return nil
	}

	return usr.(*typ.User)
}

func InitAPIv1(g *echo.Group) {
	g.Use(mw.BasicAuth(func(eml, pwd) bool {
		return helper.ValidLogin(eml, pwd)
	}))

	g.Use(func(c *echo.Context) error {
		if (c.Request().Header.Get(echo.Upgrade)) == echo.WebSocket {
			return nil
		}

		auth := c.Request().Header.Get(echo.Authorization)
		l := len(Basic)

		if len(auth) > l+1 && auth[:l] == Basic {
			b, err := base64.StdEncoding.DecodeString(auth[l+1:])
			if err == nil {
				cred := string(b)
				for i := 0; i < len(cred); i++ {
					if cred[i] == ':' {
						usr, err := persistence.ReadUser(cred[:i])
						if err != nil {
							return err
						}

						if bcrypt.CompareHashAndPassword([]byte(usr.Hash), []byte(cred[i+1:])) == nil {
							c.Set("api_usr", usr)
							return nil
						}
					}
				}
			}
		}
		c.Response().Header().Set(echo.WWWAuthenticate, Basic+" realm=Restricted")
		return echo.NewHTTPError(http.StatusUnauthorized)
	})

	g.Get("/ticket", func(c *echo.Context) error {
		usr := getUser(c)
		if usr == nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		var filters []*typ.TicketFilter
		filtersString := c.Form("filters")

		if filtersString != "" && json.Unmarshal([]byte(filtersString), &filters) != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		tickets, err := controller.GetTickets(usr, filters)
		if err != nil {
			if err == controller.NotAllowedErr {
				return c.NoContent(http.StatusForbidden)
			}
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, tickets)
	})

	g.Get("/ticket/:guid", func(c *echo.Context) error {
		usr := getUser(c)
		if usr == nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		t, err := controller.ReadTicket(usr, typ.GUID(c.Param("guid")))
		if err != nil {
			if err == controller.NotAllowedErr {
				return c.NoContent(http.StatusForbidden)
			}
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, t)
	})

	g.Post("/ticket", func(c *echo.Context) error {
		usr := getUser(c)
		if usr == nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		defer c.Request().Body.Close()

		var t *typ.Ticket
		byt, err := ioutil.ReadAll(c.Request().Body)
		if err != nil || json.Unmarshal(byt, t) != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		err = controller.UpsertTicket(usr, t)
		if err == nil {
			return c.NoContent(http.StatusOK)
		}

		if err == controller.NotAllowedErr {
			return c.NoContent(http.StatusForbidden)
		}

		return c.NoContent(http.StatusInternalServerError)
	})
}
