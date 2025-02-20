// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"github.com/dmji/go-redmine/internal/utils"
	"github.com/dmji/go-redmine/models/components"
	"github.com/dmji/go-redmine/types"
)

type UpdateTimeEntryTimeEntry struct {
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

func (u UpdateTimeEntryTimeEntry) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateTimeEntryTimeEntry) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateTimeEntryTimeEntry) GetIssueID() *int64 {
	if o == nil {
		return nil
	}
	return o.IssueID
}

func (o *UpdateTimeEntryTimeEntry) GetProjectID() *int64 {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *UpdateTimeEntryTimeEntry) GetSpentOn() *types.Date {
	if o == nil {
		return nil
	}
	return o.SpentOn
}

func (o *UpdateTimeEntryTimeEntry) GetHours() float64 {
	if o == nil {
		return 0.0
	}
	return o.Hours
}

func (o *UpdateTimeEntryTimeEntry) GetActivityID() *int64 {
	if o == nil {
		return nil
	}
	return o.ActivityID
}

func (o *UpdateTimeEntryTimeEntry) GetComments() *string {
	if o == nil {
		return nil
	}
	return o.Comments
}

func (o *UpdateTimeEntryTimeEntry) GetUserID() *int64 {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *UpdateTimeEntryTimeEntry) GetCustomFields() []components.CustomFields {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *UpdateTimeEntryTimeEntry) GetCustomFieldValues() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomFieldValues
}

type UpdateTimeEntryRequestBody struct {
	TimeEntry *UpdateTimeEntryTimeEntry `json:"time_entry,omitempty"`
}

func (o *UpdateTimeEntryRequestBody) GetTimeEntry() *UpdateTimeEntryTimeEntry {
	if o == nil {
		return nil
	}
	return o.TimeEntry
}

type UpdateTimeEntryRequest struct {
	Format             components.Format           `pathParam:"style=simple,explode=false,name=format"`
	TimeEntryID        int64                       `pathParam:"style=simple,explode=false,name=time_entry_id"`
	XRedmineSwitchUser *string                     `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	RequestBody        *UpdateTimeEntryRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateTimeEntryRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *UpdateTimeEntryRequest) GetTimeEntryID() int64 {
	if o == nil {
		return 0
	}
	return o.TimeEntryID
}

func (o *UpdateTimeEntryRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *UpdateTimeEntryRequest) GetRequestBody() *UpdateTimeEntryRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type UpdateTimeEntryResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *UpdateTimeEntryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
