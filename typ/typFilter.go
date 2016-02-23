package typ //import "go.iondynamics.net/helpdesk/typ"

type TicketFilter struct {
	ClientEmail   *StringFilter
	AssigneeEmail *StringFilter
	Subject       *StringFilter
	State         *StateFilter
	Tag           *StringFilter
}

func (f *TicketFilter) Check(t *Ticket) bool {
	result := true

	if f.ClientEmail != nil {
		result = result && f.ClientEmail.Check(t.ClientEmail)
	}

	if f.AssigneeEmail != nil {
		result = result && f.AssigneeEmail.Check(t.AssigneeEmail)
	}

	if f.Subject != nil {
		result = result && f.Subject.Check(t.Subject)
	}

	if f.State != nil {
		result = result && f.State.Check(t.State)
	}

	if f.Tag != nil && result == true {
		tagCheck := false
		if f.Tag.Contains {
			for _, tag := range t.Tags {
				if f.Tag.Check(tag) {
					tagCheck = true
					break
				}
			}
		} else if len(t.Tags) > 0 {
			tagCheck = true
			for _, tag := range t.Tags {
				if !f.Tag.Check(tag) {
					tagCheck = false
					break
				}
			}
		}
		result = result && tagCheck
	}

	return result
}

type NoteFilter struct {
	CreatorEmail *StringFilter
	Content      *StringFilter
	Internal     *BoolFilter
	NewerThan    *TimeFilter
	OlderThan    *TimeFilter
}

func (f *NoteFilter) Check(n *Note) bool {
	result := true

	if f.CreatorEmail != nil {
		result = f.CreatorEmail.Check(n.CreatorEmail)
	}

	if result && f.Content != nil {
		result = f.Content.Check(n.Content)
	}

	if result && f.Internal != nil {
		result = f.Internal.Check(n.Internal)
	}

	if result && f.NewerThan != nil && f.NewerThan.IsSet() {
		result = n.Timestamp.After(f.NewerThan.Get())
	}

	if result && f.OlderThan != nil && f.OlderThan.IsSet() {
		result = f.OlderThan.Get().After(n.Timestamp)
	}

	return result
}

type AttachmentFilter struct {
	Uploader  *StringFilter
	Name      *StringFilter
	NewerThan *TimeFilter
	OlderThan *TimeFilter
}

func (f *AttachmentFilter) Check(a *Attachment) bool {
	result := true

	if f.Uploader != nil {
		result = f.Uploader.Check(a.Uploader)
	}

	if result && f.Name != nil {
		result = f.Name.Check(a.Name)
	}

	if result && f.NewerThan != nil && f.NewerThan.IsSet() {
		result = a.Timestamp.After(f.NewerThan.Get())
	}

	if result && f.OlderThan != nil && f.OlderThan.IsSet() {
		result = f.OlderThan.Get().After(a.Timestamp)
	}

	return result
}

type UserFilter struct {
	Email      *StringFilter
	Hash       *StringFilter
	Role       *RoleFilter
	AllowedTag *StringFilter
}

func (f *UserFilter) Check(u *User) bool {
	result := true

	if f.Email != nil {
		result = f.Email.Check(u.Email)
	}

	if result && f.Hash != nil {
		result = f.Hash.Check(u.Hash)
	}

	if result && f.Role != nil {
		result = f.Role.Check(u.Role)
	}

	if result && f.AllowedTag != nil {
		tagCheck := false
		if f.AllowedTag.Contains {
			for _, tag := range u.AllowedTags {
				if f.AllowedTag.Check(tag) {
					tagCheck = true
					break
				}
			}
		} else if len(u.AllowedTags) > 0 {
			tagCheck = true
			for _, tag := range u.AllowedTags {
				if !f.AllowedTag.Check(tag) {
					tagCheck = false
					break
				}
			}
		}
		result = tagCheck
	}

	return result
}
