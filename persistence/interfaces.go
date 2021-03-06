package persistence //import "go.iondynamics.net/helpdesk/persistence"

import (
	"go.iondynamics.net/helpdesk/typ"
)

type User interface {
	RegisterUser(email, password string) error
	LoginUser(email, password string) (*User, error)
	LogoutUser() error
}

type Ticket interface {
	CreateTicket(clientEmail, subject string, initialNote typ.GUID) (typ.GUID, error)
	TicketList(clientFilter, asigneeFilter string, stateFilter typ.State) map[typ.GUID]*Ticket
	OneTicket(ID typ.GUID) (*Ticket, error)
	UpdateTicket(ID typ.GUID, ClientEmail, Subject string, State typ.State) error
	DeleteTicket(ID typ.GUID) error
	Assign(ID typ.GUID, AssigneeEmail string) error
	AddTag(ID typ.GUID, Tag string) error
	DelTag(ID typ.GUID, Tag string) error
}

type Note interface {
	CreateNote(creatorEmail, content string) (typ.GUID, error)
	OneNote(ID typ.GUID) (*Note, error)
	UpdateNote(ID typ.GUID, content string) error
	DeleteNote(ID typ.GUID) error
	Internal(ID typ.GUID, internal bool) error
	AddAttachment(ID typ.GUID, AttachmentID typ.GUID) error
	DelAttachment(ID typ.GUID, AttachmentID typ.GUID) error
}

type Provider interface {
	UpsertUser(email string, u *typ.User) error
	ReadUser(email string) (*typ.User, error)
	DeleteUser(email string) error
	UserExists(email string) (bool, error)
	GetUsers(filters []*typ.UserFilter) ([]*typ.User, error)

	UpsertTicket(ID typ.GUID, t *typ.Ticket) error
	ReadTicket(ID typ.GUID) (*typ.Ticket, error)
	DeleteTicket(ID typ.GUID) error
	TicketExists(ID typ.GUID) (bool, error)
	GetTickets(filters []*typ.TicketFilter) ([]*typ.Ticket, error)

	UpsertNote(ID typ.GUID, n *typ.Note) error
	ReadNote(ID typ.GUID) (*typ.Note, error)
	DeleteNote(ID typ.GUID) error
	NoteExists(ID typ.GUID) (bool, error)
	GetNotes(filters []*typ.NoteFilter) ([]*typ.Note, error)

	UpsertAttachment(ID typ.GUID, a *typ.Attachment) error
	ReadAttachment(ID typ.GUID) (*typ.Attachment, error)
	DeleteAttachment(ID typ.GUID) error
	AttachmentExists(ID typ.GUID) (bool, error)
	GetAttachments(filters []*typ.AttachmentFilter) ([]*typ.Attachment, error)

	Close() error
}

type NotFoundError interface {
	error
	IsNotFoundError()
}
