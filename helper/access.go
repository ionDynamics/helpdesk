package helper //import "go.iondynamics.net/helpdesk/helper"

import (
	"go.iondynamics.net/helpdesk/typ"
)

func TicketAccessAllowed(u *typ.User, t *typ.Ticket) bool {
	if u.Role != typ.Admin {
		if u.Role != typ.Team && t.ClientEmail != u.Email {
			return false
		}

		if u.Email == t.AssigneeEmail {
			return true
		}

		allowed := false
		for _, allowedTag := range u.AllowedTags {
			for _, ticketTag := range t.Tags {
				if allowedTag == ticketTag {
					allowed = true
				}
			}
		}
		return allowed
	}
	return true
}
