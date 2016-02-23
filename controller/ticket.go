package controller //import "go.iondynamics.net/helpdesk/controller"

import (
	"errors"

	idl "go.iondynamics.net/iDlogger"

	"go.iondynamics.net/helpdesk/helper"
	"go.iondynamics.net/helpdesk/persistence"
	"go.iondynamics.net/helpdesk/typ"
)

var NotAllowedErr = errors.New("not allowed")

func log(err error) error {
	if err == nil {
		return err
	}

	if !persistence.IsNotFound(err) {
		idl.Crit(err)
	}
	idl.Debug(err)
	return err
}

func UpsertTicket(u *typ.User, t *typ.Ticket) error {
	id := helper.GenerateGUID()
	if t.ID != "" {
		id = t.ID
		o, err := persistence.ReadTicket(id)
		if err != nil {
			log(err)
			if !persistence.IsNotFound(err) {
				return err
			}
		} else if !helper.TicketAccessAllowed(u, o) {
			return NotAllowedErr
		}
	} else {
		t.ID = id
	}

	if !helper.TicketAccessAllowed(u, t) {
		return NotAllowedErr
	}

	return log(persistence.UpsertTicket(id, t))

}

func ReadTicket(u *typ.User, ID typ.GUID) (*typ.Ticket, error) {
	t, err := persistence.ReadTicket(ID)
	if err != nil {
		return nil, log(err)
	}

	if helper.TicketAccessAllowed(u, t) {
		return t, nil
	}

	idl.Debug("not allowed")
	return nil, NotAllowedErr
}

func DeleteTicket(u *typ.User, ID typ.GUID) error {
	_, err := ReadTicket(u, ID)
	if err != nil {
		return err
	}

	return persistence.DeleteTicket(ID)
}

func TicketExists(ID typ.GUID) (bool, error) {
	return persistence.TicketExists(ID)
}

func GetTickets(u *typ.User, filters []*typ.TicketFilter) ([]*typ.Ticket, error) {
	var ret []*typ.Ticket

	tickets, err := persistence.GetTickets(filters)
	if err != nil {
		return ret, log(err)
	}

	for _, ticket := range tickets {
		if helper.TicketAccessAllowed(u, ticket) {
			ret = append(ret, ticket)
		}
	}

	return ret, nil
}
