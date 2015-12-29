package persistence //import "go.iondynamics.net/helpdesk/persistence"

import (
	"go.iondynamics.net/helpdesk/typ"
)

var dbo PersistenceProvider

func Init(instance PersistenceProvider) {
	dbo = instance
}

func Close() error {
	return dbo.Close()
}

func UpsertUser(email string, u *typ.User) error {
	return dbo.UpsertUser(email, u)
}

func ReadUser(email string) (*typ.User, error) {
	return dbo.ReadUser(email)
}

func DeleteUser(email string, u *typ.User) error {
	return dbo.DeleteUser(email, u)
}

func UserExists(email string) (bool, error) {
	return dbo.UserExists(email)
}

//

func UpsertTicket(ID typ.GUID, t *typ.Ticket) error {
	return dbo.UpsertTicket(ID, t)
}

func ReadTicket(ID typ.GUID) (*typ.Ticket, error) {
	return dbo.ReadTicket(ID)
}

func DeleteTicket(ID typ.GUID, t *typ.Ticket) error {
	return dbo.DeleteTicket(ID, t)
}

func TicketExists(ID typ.GUID) (bool, error) {
	return dbo.TicketExists(ID)
}

//

func UpsertNote(ID typ.GUID, n *typ.Note) error {
	return dbo.UpsertNote(ID, n)
}

func ReadNote(ID typ.GUID) (*typ.Note, error) {
	return dbo.ReadNote(ID)
}

func DeleteNote(ID typ.GUID, n *typ.Note) error {
	return dbo.DeleteNote(ID, n)
}

func NoteExists(ID typ.GUID) (bool, error) {
	return dbo.NoteExists(ID)
}

//

func UpsertAttachment(ID typ.GUID, a *typ.Attachment) error {
	return dbo.UpsertAttachment(ID, a)
}

func ReadAttachment(ID typ.GUID) (*typ.Attachment, error) {
	return dbo.ReadAttachment(ID)
}

func DeleteAttachment(ID typ.GUID, a *typ.Attachment) error {
	return dbo.DeleteAttachment(ID, a)
}

func AttachmentExists(ID typ.GUID) (bool, error) {
	return dbo.AttachmentExists(ID)
}
