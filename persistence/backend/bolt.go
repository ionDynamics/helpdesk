package backend //import "go.iondynamics.net/helpdesk/persistence/backend"

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"

	"go.iondynamics.net/helpdesk/persistence"
	"go.iondynamics.net/helpdesk/typ"
)

type NotFoundErr struct{}

func (e *NotFoundErr) Error() string {
	return "not found"
}

func (e *NotFoundErr) IsNotFoundError() {}

type Bolt struct {
	db *bolt.DB
}

func InitBolt(path string) (persistence.Provider, error) {
	var err error
	b := &Bolt{}
	b.db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	return b, err
}

func (blt *Bolt) upsert(bucket, key []byte, val interface{}) error {
	return blt.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		byt, err := json.Marshal(val)
		if err != nil {
			return err
		}

		return b.Put([]byte(key), byt)
	})
}

func (blt *Bolt) read(bucket, key []byte, ptr interface{}) error {
	return blt.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return &NotFoundErr{}
		}

		byt := b.Get([]byte(key))
		if byt == nil {
			return &NotFoundErr{}
		}

		return json.Unmarshal(byt, ptr)
	})
}

func (blt *Bolt) delete(bucket, key []byte) error {
	return blt.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		return b.Delete(key)
	})
}

func (blt *Bolt) exists(bucket, key []byte) (bool, error) {
	return blt.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return &NotFoundErr{}
		}

		byt := b.Get([]byte(key))
		if byt == nil {
			return &NotFoundErr{}
		}
		return nil
	}) == nil, nil
}

func (blt *Bolt) getAll(bucket []byte, gen func() interface{}) ([]interface{}, error) {
	var slice []interface{}
	return slice, blt.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return &NotFoundErr{}
		}

		return b.ForEach(func(k, v []byte) error {
			ptr := gen()
			err := json.Unmarshal(v, ptr)
			if err != nil {
				return err
			}
			slice = append(slice, ptr)
			return nil
		})
	})
}

//

func (blt *Bolt) UpsertUser(email string, u *typ.User) error {
	return blt.upsert([]byte("users"), []byte(email), u)
}

func (blt *Bolt) ReadUser(email string) (*typ.User, error) {
	usr := &typ.User{}
	return usr, blt.read([]byte("users"), []byte(email), usr)
}

func (blt *Bolt) DeleteUser(email string) error {
	return blt.delete([]byte("users"), []byte(email))
}

func (blt *Bolt) UserExists(email string) (bool, error) {
	return blt.exists([]byte("users"), []byte(email))
}

func (blt *Bolt) GetUsers(filters []*typ.UserFilter) (u []*typ.User, e error) {
	var users []*typ.User
	vals, err := blt.getAll([]byte("users"), func() interface{} { return &typ.User{} })
	if err != nil {
		return nil, err
	}
	for _, v := range vals {
		t := v.(*typ.User)

		for _, filter := range filters {
			if !filter.Check(t) {
				t = nil
			}
		}

		if t != nil {
			users = append(users, t)
		}
	}
	return users, nil
}

//

func (blt *Bolt) UpsertTicket(ID typ.GUID, t *typ.Ticket) error {
	return blt.upsert([]byte("tickets"), []byte(ID), t)
}

func (blt *Bolt) ReadTicket(ID typ.GUID) (*typ.Ticket, error) {
	tic := &typ.Ticket{}
	return tic, blt.read([]byte("tickets"), []byte(ID), tic)
}

func (blt *Bolt) DeleteTicket(ID typ.GUID) error {
	return blt.delete([]byte("tickets"), []byte(ID))
}

func (blt *Bolt) TicketExists(ID typ.GUID) (bool, error) {
	return blt.exists([]byte("tickets"), []byte(ID))
}

func (blt *Bolt) GetTickets(filters []*typ.TicketFilter) ([]*typ.Ticket, error) {
	var tickets []*typ.Ticket
	vals, err := blt.getAll([]byte("tickets"), func() interface{} { return &typ.Ticket{} })
	if err != nil {
		return nil, err
	}
	for _, v := range vals {
		t := v.(*typ.Ticket)

		for _, filter := range filters {
			if !filter.Check(t) {
				t = nil
			}
		}

		if t != nil {
			tickets = append(tickets, t)
		}
	}
	return tickets, nil
}

//

func (blt *Bolt) UpsertNote(ID typ.GUID, n *typ.Note) error {
	return blt.upsert([]byte("notes"), []byte(ID), n)
}

func (blt *Bolt) ReadNote(ID typ.GUID) (*typ.Note, error) {
	not := &typ.Note{}
	return not, blt.read([]byte("notes"), []byte(ID), not)
}

func (blt *Bolt) DeleteNote(ID typ.GUID) error {
	return blt.delete([]byte("notes"), []byte(ID))
}

func (blt *Bolt) NoteExists(ID typ.GUID) (bool, error) {
	return blt.exists([]byte("notes"), []byte(ID))
}

func (blt *Bolt) GetNotes(filters []*typ.NoteFilter) (n []*typ.Note, e error) {
	var notes []*typ.Note
	vals, err := blt.getAll([]byte("notes"), func() interface{} { return &typ.Note{} })
	if err != nil {
		return nil, err
	}
	for _, v := range vals {
		t := v.(*typ.Note)

		for _, filter := range filters {
			if !filter.Check(t) {
				t = nil
			}
		}

		if t != nil {
			notes = append(notes, t)
		}
	}
	return notes, nil
}

//

func (blt *Bolt) UpsertAttachment(ID typ.GUID, a *typ.Attachment) error {
	return blt.upsert([]byte("attachments"), []byte(ID), a)
}

func (blt *Bolt) ReadAttachment(ID typ.GUID) (*typ.Attachment, error) {
	att := &typ.Attachment{}
	return att, blt.read([]byte("attachments"), []byte(ID), att)
}

func (blt *Bolt) DeleteAttachment(ID typ.GUID) error {
	return blt.delete([]byte("attachments"), []byte(ID))
}

func (blt *Bolt) AttachmentExists(ID typ.GUID) (bool, error) {
	return blt.exists([]byte("attachments"), []byte(ID))
}

func (blt *Bolt) GetAttachments(filters []*typ.AttachmentFilter) (a []*typ.Attachment, e error) {
	var attachments []*typ.Attachment
	vals, err := blt.getAll([]byte("attachments"), func() interface{} { return &typ.Attachment{} })
	if err != nil {
		return nil, err
	}
	for _, v := range vals {
		t := v.(*typ.Attachment)

		for _, filter := range filters {
			if !filter.Check(t) {
				t = nil
			}
		}

		if t != nil {
			attachments = append(attachments, t)
		}
	}
	return attachments, nil
}

//

func (blt *Bolt) Close() error {
	return blt.db.Close()
}
