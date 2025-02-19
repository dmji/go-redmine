// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/dmji/go-redmine/internal/utils"
	"github.com/dmji/go-redmine/models/components"
	"github.com/dmji/go-redmine/types"
)

type UpdateIssueIssue struct {
	ProjectID         *int64                    `json:"project_id,omitempty"`
	TrackerID         *int64                    `json:"tracker_id,omitempty"`
	StatusID          *int64                    `json:"status_id,omitempty"`
	PriorityID        *int64                    `json:"priority_id,omitempty"`
	Subject           *string                   `json:"subject,omitempty"`
	Description       *string                   `json:"description,omitempty"`
	StartDate         *types.Date               `json:"start_date,omitempty"`
	DueDate           *types.Date               `json:"due_date,omitempty"`
	DoneRatio         *int64                    `json:"done_ratio,omitempty"`
	CategoryID        *int64                    `json:"category_id,omitempty"`
	FixedVersionID    *int64                    `json:"fixed_version_id,omitempty"`
	AssignedToID      *int64                    `json:"assigned_to_id,omitempty"`
	ParentIssueID     *int64                    `json:"parent_issue_id,omitempty"`
	CustomFields      []components.CustomFields `json:"custom_fields,omitempty"`
	CustomFieldValues map[string]any            `json:"custom_field_values,omitempty"`
	WatcherUserIds    []int64                   `json:"watcher_user_ids,omitempty"`
	IsPrivate         *bool                     `json:"is_private,omitempty"`
	EstimatedHours    *float64                  `json:"estimated_hours,omitempty"`
	Notes             *string                   `json:"notes,omitempty"`
	PrivateNotes      *string                   `json:"private_notes,omitempty"`
	Uploads           []components.Uploads      `json:"uploads,omitempty"`
}

func (u UpdateIssueIssue) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateIssueIssue) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateIssueIssue) GetProjectID() *int64 {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *UpdateIssueIssue) GetTrackerID() *int64 {
	if o == nil {
		return nil
	}
	return o.TrackerID
}

func (o *UpdateIssueIssue) GetStatusID() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusID
}

func (o *UpdateIssueIssue) GetPriorityID() *int64 {
	if o == nil {
		return nil
	}
	return o.PriorityID
}

func (o *UpdateIssueIssue) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *UpdateIssueIssue) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateIssueIssue) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *UpdateIssueIssue) GetDueDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *UpdateIssueIssue) GetDoneRatio() *int64 {
	if o == nil {
		return nil
	}
	return o.DoneRatio
}

func (o *UpdateIssueIssue) GetCategoryID() *int64 {
	if o == nil {
		return nil
	}
	return o.CategoryID
}

func (o *UpdateIssueIssue) GetFixedVersionID() *int64 {
	if o == nil {
		return nil
	}
	return o.FixedVersionID
}

func (o *UpdateIssueIssue) GetAssignedToID() *int64 {
	if o == nil {
		return nil
	}
	return o.AssignedToID
}

func (o *UpdateIssueIssue) GetParentIssueID() *int64 {
	if o == nil {
		return nil
	}
	return o.ParentIssueID
}

func (o *UpdateIssueIssue) GetCustomFields() []components.CustomFields {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *UpdateIssueIssue) GetCustomFieldValues() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomFieldValues
}

func (o *UpdateIssueIssue) GetWatcherUserIds() []int64 {
	if o == nil {
		return nil
	}
	return o.WatcherUserIds
}

func (o *UpdateIssueIssue) GetIsPrivate() *bool {
	if o == nil {
		return nil
	}
	return o.IsPrivate
}

func (o *UpdateIssueIssue) GetEstimatedHours() *float64 {
	if o == nil {
		return nil
	}
	return o.EstimatedHours
}

func (o *UpdateIssueIssue) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}

func (o *UpdateIssueIssue) GetPrivateNotes() *string {
	if o == nil {
		return nil
	}
	return o.PrivateNotes
}

func (o *UpdateIssueIssue) GetUploads() []components.Uploads {
	if o == nil {
		return nil
	}
	return o.Uploads
}

type UpdateIssueRequestBody struct {
	Issue *UpdateIssueIssue `json:"issue,omitempty"`
}

func (o *UpdateIssueRequestBody) GetIssue() *UpdateIssueIssue {
	if o == nil {
		return nil
	}
	return o.Issue
}

type UpdateIssueRequest struct {
	Format             components.Format       `pathParam:"style=simple,explode=false,name=format"`
	IssueID            int64                   `pathParam:"style=simple,explode=false,name=issue_id"`
	XRedmineSwitchUser *string                 `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	RequestBody        *UpdateIssueRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateIssueRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *UpdateIssueRequest) GetIssueID() int64 {
	if o == nil {
		return 0
	}
	return o.IssueID
}

func (o *UpdateIssueRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *UpdateIssueRequest) GetRequestBody() *UpdateIssueRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type UpdateIssueResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *UpdateIssueResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
