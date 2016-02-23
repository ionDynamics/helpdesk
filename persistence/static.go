package persistence //import "go.iondynamics.net/helpdesk/persistence"

import (
	"go.iondynamics.net/helpdesk/typ"
)

var dbo Provider

func Init(instance Provider) {
	dbo = instance
}

func IsNotFound(err error) bool {
	_, ok := err.(NotFoundError)
	return ok
}

func Close() error {
	return dbo.Close()
}

//

func UpsertUser(email string, u *typ.User) error {
	return dbo.UpsertUser(email, u)
}

func ReadUser(email string) (*typ.User, error) {
	return dbo.ReadUser(email)
}

func DeleteUser(email string) error {
	return dbo.DeleteUser(email)
}

func UserExists(email string) (bool, error) {
	return dbo.UserExists(email)
}

func GetUsers(filters []*typ.UserFilter) ([]*typ.User, error) {
	return dbo.GetUsers(filters)
}

//

func UpsertTicket(ID typ.GUID, t *typ.Ticket) error {
	return dbo.UpsertTicket(ID, t)
}

func ReadTicket(ID typ.GUID) (*typ.Ticket, error) {
	return dbo.ReadTicket(ID)
}

func DeleteTicket(ID typ.GUID) error {
	return dbo.DeleteTicket(ID)
}

func TicketExists(ID typ.GUID) (bool, error) {
	return dbo.TicketExists(ID)
}

func GetTickets(filters []*typ.TicketFilter) ([]*typ.Ticket, error) {
	return dbo.GetTickets(filters)
}

//

func UpsertNote(ID typ.GUID, n *typ.Note) error {
	return dbo.UpsertNote(ID, n)
}

func ReadNote(ID typ.GUID) (*typ.Note, error) {
	return dbo.ReadNote(ID)
}

func DeleteNote(ID typ.GUID) error {
	return dbo.DeleteNote(ID)
}

func NoteExists(ID typ.GUID) (bool, error) {
	return dbo.NoteExists(ID)
}

func GetNotes(filters []*typ.NoteFilter) ([]*typ.Note, error) {
	return dbo.GetNotes(filters)
}

//

func UpsertAttachment(ID typ.GUID, a *typ.Attachment) error {
	return dbo.UpsertAttachment(ID, a)
}

func ReadAttachment(ID typ.GUID) (*typ.Attachment, error) {
	return dbo.ReadAttachment(ID)
}

func DeleteAttachment(ID typ.GUID) error {
	return dbo.DeleteAttachment(ID)
}

func AttachmentExists(ID typ.GUID) (bool, error) {
	return dbo.AttachmentExists(ID)
}

func GetAttachments(filters []*typ.AttachmentFilter) ([]*typ.Attachment, error) {
	return dbo.GetAttachments(filters)
}
