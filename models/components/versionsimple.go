// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/dmji/go-redmine/internal/utils"
	"github.com/dmji/go-redmine/types"
	"time"
)

type Status string

const (
	StatusOpen   Status = "open"
	StatusLocked Status = "locked"
	StatusClosed Status = "closed"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "open":
		fallthrough
	case "locked":
		fallthrough
	case "closed":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type Sharing string

const (
	SharingNone        Sharing = "none"
	SharingDescendants Sharing = "descendants"
	SharingHierarchy   Sharing = "hierarchy"
	SharingTree        Sharing = "tree"
	SharingSystem      Sharing = "system"
)

func (e Sharing) ToPointer() *Sharing {
	return &e
}
func (e *Sharing) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "descendants":
		fallthrough
	case "hierarchy":
		fallthrough
	case "tree":
		fallthrough
	case "system":
		*e = Sharing(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Sharing: %v", v)
	}
}

type VersionSimple struct {
	ID            int64              `json:"id"`
	Project       IDName             `json:"project"`
	Name          string             `json:"name"`
	Description   *string            `json:"description"`
	Status        Status             `json:"status"`
	DueDate       *types.Date        `json:"due_date"`
	Sharing       Sharing            `json:"sharing"`
	WikiPageTitle *string            `json:"wiki_page_title"`
	CustomFields  []CustomFieldValue `json:"custom_fields,omitempty"`
	CreatedOn     time.Time          `json:"created_on"`
	UpdatedOn     time.Time          `json:"updated_on"`
}

func (v VersionSimple) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *VersionSimple) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *VersionSimple) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *VersionSimple) GetProject() IDName {
	if o == nil {
		return IDName{}
	}
	return o.Project
}

func (o *VersionSimple) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *VersionSimple) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *VersionSimple) GetStatus() Status {
	if o == nil {
		return Status("")
	}
	return o.Status
}

func (o *VersionSimple) GetDueDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *VersionSimple) GetSharing() Sharing {
	if o == nil {
		return Sharing("")
	}
	return o.Sharing
}

func (o *VersionSimple) GetWikiPageTitle() *string {
	if o == nil {
		return nil
	}
	return o.WikiPageTitle
}

func (o *VersionSimple) GetCustomFields() []CustomFieldValue {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *VersionSimple) GetCreatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedOn
}

func (o *VersionSimple) GetUpdatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedOn
}
