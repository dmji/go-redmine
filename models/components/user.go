// Code generated by Speakeasy (https://speakeasy.com).

package components

import (
	"time"

	"github.com/dmji/go-redmine/internal/utils"
)

type UserTwofaScheme struct{}

type Roles struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Inherited *bool  `json:"inherited,omitempty"`
}

func (o *Roles) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Roles) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Roles) GetInherited() *bool {
	if o == nil {
		return nil
	}
	return o.Inherited
}

type Memberships struct {
	ID      int64   `json:"id"`
	Project IDName  `json:"project"`
	Roles   []Roles `json:"roles"`
}

func (o *Memberships) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Memberships) GetProject() IDName {
	if o == nil {
		return IDName{}
	}
	return o.Project
}

func (o *Memberships) GetRoles() []Roles {
	if o == nil {
		return []Roles{}
	}
	return o.Roles
}

type User struct {
	ID              int64              `json:"id"`
	Login           string             `json:"login"`
	Admin           bool               `json:"admin"`
	Firstname       string             `json:"firstname"`
	Lastname        string             `json:"lastname"`
	Mail            string             `json:"mail"`
	CreatedOn       time.Time          `json:"created_on"`
	UpdatedOn       time.Time          `json:"updated_on"`
	LastLoginOn     *time.Time         `json:"last_login_on"`
	PasswdChangedOn *time.Time         `json:"passwd_changed_on"`
	TwofaScheme     *UserTwofaScheme   `json:"twofa_scheme"`
	APIKey          string             `json:"api_key"`
	Status          int64              `json:"status"`
	CustomFields    []CustomFieldValue `json:"custom_fields,omitempty"`
	Groups          []IDName           `json:"groups,omitempty"`
	Memberships     []Memberships      `json:"memberships,omitempty"`
}

func (u User) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *User) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *User) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *User) GetLogin() string {
	if o == nil {
		return ""
	}
	return o.Login
}

func (o *User) GetAdmin() bool {
	if o == nil {
		return false
	}
	return o.Admin
}

func (o *User) GetFirstname() string {
	if o == nil {
		return ""
	}
	return o.Firstname
}

func (o *User) GetLastname() string {
	if o == nil {
		return ""
	}
	return o.Lastname
}

func (o *User) GetMail() string {
	if o == nil {
		return ""
	}
	return o.Mail
}

func (o *User) GetCreatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedOn
}

func (o *User) GetUpdatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedOn
}

func (o *User) GetLastLoginOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastLoginOn
}

func (o *User) GetPasswdChangedOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.PasswdChangedOn
}

func (o *User) GetTwofaScheme() *UserTwofaScheme {
	if o == nil {
		return nil
	}
	return o.TwofaScheme
}

func (o *User) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *User) GetStatus() int64 {
	if o == nil {
		return 0
	}
	return o.Status
}

func (o *User) GetCustomFields() []CustomFieldValue {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *User) GetGroups() []IDName {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *User) GetMemberships() []Memberships {
	if o == nil {
		return nil
	}
	return o.Memberships
}
