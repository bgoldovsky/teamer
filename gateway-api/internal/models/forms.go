package models

const (
	RoleNone     = 0
	RoleBackEnd  = 1
	RoleFrontEnd = 2
	RoleMobile   = 3
	RoleQA       = 4
)

type TeamForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Slack       string `json:"slack"`
}
