// Code generated by Speakeasy (https://speakeasy.com).

package components

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dmji/go-redmine/internal/utils"
	"github.com/dmji/go-redmine/types"
)

type VersionStatus string

const (
	VersionStatusOpen   VersionStatus = "open"
	VersionStatusLocked VersionStatus = "locked"
	VersionStatusClosed VersionStatus = "closed"
)

func (e VersionStatus) ToPointer() *VersionStatus {
	return &e
}

func (e *VersionStatus) UnmarshalJSON(data []byte) error {
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
		*e = VersionStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VersionStatus: %v", v)
	}
}

type VersionSharing string

const (
	VersionSharingNone        VersionSharing = "none"
	VersionSharingDescendants VersionSharing = "descendants"
	VersionSharingHierarchy   VersionSharing = "hierarchy"
	VersionSharingTree        VersionSharing = "tree"
	VersionSharingSystem      VersionSharing = "system"
)

func (e VersionSharing) ToPointer() *VersionSharing {
	return &e
}

func (e *VersionSharing) UnmarshalJSON(data []byte) error {
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
		*e = VersionSharing(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VersionSharing: %v", v)
	}
}

type Version struct {
	ID             int64              `json:"id"`
	Project        IDName             `json:"project"`
	Name           string             `json:"name"`
	Description    *string            `json:"description"`
	Status         VersionStatus      `json:"status"`
	DueDate        *types.Date        `json:"due_date"`
	Sharing        VersionSharing     `json:"sharing"`
	WikiPageTitle  *string            `json:"wiki_page_title"`
	EstimatedHours *float64           `json:"estimated_hours,omitempty"`
	SpentHours     *float64           `json:"spent_hours,omitempty"`
	CustomFields   []CustomFieldValue `json:"custom_fields,omitempty"`
	CreatedOn      time.Time          `json:"created_on"`
	UpdatedOn      time.Time          `json:"updated_on"`
}

func (v Version) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *Version) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Version) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Version) GetProject() IDName {
	if o == nil {
		return IDName{}
	}
	return o.Project
}

func (o *Version) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Version) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Version) GetStatus() VersionStatus {
	if o == nil {
		return VersionStatus("")
	}
	return o.Status
}

func (o *Version) GetDueDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *Version) GetSharing() VersionSharing {
	if o == nil {
		return VersionSharing("")
	}
	return o.Sharing
}

func (o *Version) GetWikiPageTitle() *string {
	if o == nil {
		return nil
	}
	return o.WikiPageTitle
}

func (o *Version) GetEstimatedHours() *float64 {
	if o == nil {
		return nil
	}
	return o.EstimatedHours
}

func (o *Version) GetSpentHours() *float64 {
	if o == nil {
		return nil
	}
	return o.SpentHours
}

func (o *Version) GetCustomFields() []CustomFieldValue {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *Version) GetCreatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedOn
}

func (o *Version) GetUpdatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedOn
}
