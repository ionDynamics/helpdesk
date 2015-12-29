package typ //import "go.iondynamics.net/helpdesk/typ"

type TicketFilter struct {
	ClientEmail   *StringFilter
	AssigneeEmail *StringFilter
	Subject       *StringFilter
	State         *StateFilter
	Tag           *StringFilter
}

type NoteFilter struct {
	CreatorEmail *StringFilter
	Content      *StringFilter
	Internal     *BoolFilter
	NewerThan    *TimeFilter
	OlderThan    *TimeFilter
}

type AttachmentFilter struct {
	Uploader  *StringFilter
	Name      *StringFilter
	NewerThan *TimeFilter
	OlderThan *TimeFilter
}

type UserFilter struct {
	Email      *StringFilter
	Hash       *StringFilter
	Role       *RoleFilter
	AllowedTag *StringFilter
}
