package typ //import "go.iondynamics.net/helpdesk/typ"

import (
	"time"
)

type GUID string

type State uint8

const (
	Open State = iota
	Assigned
	Closed
)

type Ticket struct {
	ID            GUID
	ClientEmail   string
	AssigneeEmail string
	Subject       string
	Notes         map[GUID]*Note
	State         State
	Tags          []string
}

type Note struct {
	ID           GUID
	CreatorEmail string
	Content      string
	Internal     bool
	Attachments  map[GUID]*Attachment
	Timestamp    time.Time
	History      map[time.Time]string
}

type Attachment struct {
	ID        GUID
	Uploader  string
	Name      string
	Content   []byte
	Timestamp time.Time
}

type User struct {
	Email       string
	Hash        string
	Role        Role
	AllowedTags []string
}

type Role uint8

const (
	Visitor Role = 1 << iota
	Client
	Team
	Admin
)
