// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/dmji/go-redmine/internal/utils"
	"github.com/dmji/go-redmine/types"
	"time"
)

type TimeEntryIssue struct {
	ID int64 `json:"id"`
}

func (o *TimeEntryIssue) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type TimeEntry struct {
	ID           int64              `json:"id"`
	Project      *IDName            `json:"project,omitempty"`
	Issue        *TimeEntryIssue    `json:"issue,omitempty"`
	User         IDName             `json:"user"`
	Activity     IDName             `json:"activity"`
	Hours        float64            `json:"hours"`
	Comments     *string            `json:"comments"`
	SpentOn      types.Date         `json:"spent_on"`
	CreatedOn    time.Time          `json:"created_on"`
	UpdatedOn    time.Time          `json:"updated_on"`
	CustomFields []CustomFieldValue `json:"custom_fields,omitempty"`
}

func (t TimeEntry) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TimeEntry) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TimeEntry) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *TimeEntry) GetProject() *IDName {
	if o == nil {
		return nil
	}
	return o.Project
}

func (o *TimeEntry) GetIssue() *TimeEntryIssue {
	if o == nil {
		return nil
	}
	return o.Issue
}

func (o *TimeEntry) GetUser() IDName {
	if o == nil {
		return IDName{}
	}
	return o.User
}

func (o *TimeEntry) GetActivity() IDName {
	if o == nil {
		return IDName{}
	}
	return o.Activity
}

func (o *TimeEntry) GetHours() float64 {
	if o == nil {
		return 0.0
	}
	return o.Hours
}

func (o *TimeEntry) GetComments() *string {
	if o == nil {
		return nil
	}
	return o.Comments
}

func (o *TimeEntry) GetSpentOn() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.SpentOn
}

func (o *TimeEntry) GetCreatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedOn
}

func (o *TimeEntry) GetUpdatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedOn
}

func (o *TimeEntry) GetCustomFields() []CustomFieldValue {
	if o == nil {
		return nil
	}
	return o.CustomFields
}
