// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/dmji/go-redmine/internal/utils"
	"time"
)

type TwofaScheme struct {
}

type UserSimple struct {
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
	TwofaScheme     *TwofaScheme       `json:"twofa_scheme"`
	CustomFields    []CustomFieldValue `json:"custom_fields,omitempty"`
}

func (u UserSimple) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UserSimple) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UserSimple) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *UserSimple) GetLogin() string {
	if o == nil {
		return ""
	}
	return o.Login
}

func (o *UserSimple) GetAdmin() bool {
	if o == nil {
		return false
	}
	return o.Admin
}

func (o *UserSimple) GetFirstname() string {
	if o == nil {
		return ""
	}
	return o.Firstname
}

func (o *UserSimple) GetLastname() string {
	if o == nil {
		return ""
	}
	return o.Lastname
}

func (o *UserSimple) GetMail() string {
	if o == nil {
		return ""
	}
	return o.Mail
}

func (o *UserSimple) GetCreatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedOn
}

func (o *UserSimple) GetUpdatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedOn
}

func (o *UserSimple) GetLastLoginOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastLoginOn
}

func (o *UserSimple) GetPasswdChangedOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.PasswdChangedOn
}

func (o *UserSimple) GetTwofaScheme() *TwofaScheme {
	if o == nil {
		return nil
	}
	return o.TwofaScheme
}

func (o *UserSimple) GetCustomFields() []CustomFieldValue {
	if o == nil {
		return nil
	}
	return o.CustomFields
}
