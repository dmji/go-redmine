// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"github.com/dmji/go-redmine/internal/utils"
	"github.com/dmji/go-redmine/models/components"
	"github.com/dmji/go-redmine/types"
)

type TimeEntry struct {
	IssueID           *int64                    `json:"issue_id,omitempty"`
	ProjectID         *int64                    `json:"project_id,omitempty"`
	SpentOn           *types.Date               `json:"spent_on,omitempty"`
	Hours             float64                   `json:"hours"`
	ActivityID        *int64                    `json:"activity_id,omitempty"`
	Comments          *string                   `json:"comments,omitempty"`
	UserID            *int64                    `json:"user_id,omitempty"`
	CustomFields      []components.CustomFields `json:"custom_fields,omitempty"`
	CustomFieldValues map[string]any            `json:"custom_field_values,omitempty"`
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

func (o *TimeEntry) GetIssueID() *int64 {
	if o == nil {
		return nil
	}
	return o.IssueID
}

func (o *TimeEntry) GetProjectID() *int64 {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *TimeEntry) GetSpentOn() *types.Date {
	if o == nil {
		return nil
	}
	return o.SpentOn
}

func (o *TimeEntry) GetHours() float64 {
	if o == nil {
		return 0.0
	}
	return o.Hours
}

func (o *TimeEntry) GetActivityID() *int64 {
	if o == nil {
		return nil
	}
	return o.ActivityID
}

func (o *TimeEntry) GetComments() *string {
	if o == nil {
		return nil
	}
	return o.Comments
}

func (o *TimeEntry) GetUserID() *int64 {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *TimeEntry) GetCustomFields() []components.CustomFields {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *TimeEntry) GetCustomFieldValues() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomFieldValues
}

type CreateTimeEntryRequestBody struct {
	TimeEntry TimeEntry `json:"time_entry"`
}

func (o *CreateTimeEntryRequestBody) GetTimeEntry() TimeEntry {
	if o == nil {
		return TimeEntry{}
	}
	return o.TimeEntry
}

type CreateTimeEntryRequest struct {
	Format             components.Format           `pathParam:"style=simple,explode=false,name=format"`
	XRedmineSwitchUser *string                     `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	RequestBody        *CreateTimeEntryRequestBody `request:"mediaType=application/json"`
}

func (o *CreateTimeEntryRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *CreateTimeEntryRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *CreateTimeEntryRequest) GetRequestBody() *CreateTimeEntryRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type CreateTimeEntryResponseBody struct {
	TimeEntry components.TimeEntry `json:"time_entry"`
}

func (o *CreateTimeEntryResponseBody) GetTimeEntry() components.TimeEntry {
	if o == nil {
		return components.TimeEntry{}
	}
	return o.TimeEntry
}

type CreateTimeEntryResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *CreateTimeEntryResponseBody
}

func (o *CreateTimeEntryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateTimeEntryResponse) GetObject() *CreateTimeEntryResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
