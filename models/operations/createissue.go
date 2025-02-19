// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"github.com/dmji/go-redmine/internal/utils"
	"github.com/dmji/go-redmine/models/components"
	"github.com/dmji/go-redmine/types"
)

type Issue struct {
	ProjectID         int64                     `json:"project_id"`
	TrackerID         *int64                    `json:"tracker_id,omitempty"`
	StatusID          *int64                    `json:"status_id,omitempty"`
	PriorityID        *int64                    `json:"priority_id,omitempty"`
	Subject           string                    `json:"subject"`
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
	Uploads           []components.Uploads      `json:"uploads,omitempty"`
}

func (i Issue) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *Issue) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Issue) GetProjectID() int64 {
	if o == nil {
		return 0
	}
	return o.ProjectID
}

func (o *Issue) GetTrackerID() *int64 {
	if o == nil {
		return nil
	}
	return o.TrackerID
}

func (o *Issue) GetStatusID() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusID
}

func (o *Issue) GetPriorityID() *int64 {
	if o == nil {
		return nil
	}
	return o.PriorityID
}

func (o *Issue) GetSubject() string {
	if o == nil {
		return ""
	}
	return o.Subject
}

func (o *Issue) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Issue) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *Issue) GetDueDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *Issue) GetDoneRatio() *int64 {
	if o == nil {
		return nil
	}
	return o.DoneRatio
}

func (o *Issue) GetCategoryID() *int64 {
	if o == nil {
		return nil
	}
	return o.CategoryID
}

func (o *Issue) GetFixedVersionID() *int64 {
	if o == nil {
		return nil
	}
	return o.FixedVersionID
}

func (o *Issue) GetAssignedToID() *int64 {
	if o == nil {
		return nil
	}
	return o.AssignedToID
}

func (o *Issue) GetParentIssueID() *int64 {
	if o == nil {
		return nil
	}
	return o.ParentIssueID
}

func (o *Issue) GetCustomFields() []components.CustomFields {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *Issue) GetCustomFieldValues() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomFieldValues
}

func (o *Issue) GetWatcherUserIds() []int64 {
	if o == nil {
		return nil
	}
	return o.WatcherUserIds
}

func (o *Issue) GetIsPrivate() *bool {
	if o == nil {
		return nil
	}
	return o.IsPrivate
}

func (o *Issue) GetEstimatedHours() *float64 {
	if o == nil {
		return nil
	}
	return o.EstimatedHours
}

func (o *Issue) GetUploads() []components.Uploads {
	if o == nil {
		return nil
	}
	return o.Uploads
}

type CreateIssueRequestBody struct {
	Issue Issue `json:"issue"`
}

func (o *CreateIssueRequestBody) GetIssue() Issue {
	if o == nil {
		return Issue{}
	}
	return o.Issue
}

type CreateIssueRequest struct {
	Format             components.Format       `pathParam:"style=simple,explode=false,name=format"`
	XRedmineSwitchUser *string                 `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	RequestBody        *CreateIssueRequestBody `request:"mediaType=application/json"`
}

func (o *CreateIssueRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *CreateIssueRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *CreateIssueRequest) GetRequestBody() *CreateIssueRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type CreateIssueResponseBody struct {
	Issue components.IssueSimple `json:"issue"`
}

func (o *CreateIssueResponseBody) GetIssue() components.IssueSimple {
	if o == nil {
		return components.IssueSimple{}
	}
	return o.Issue
}

type CreateIssueResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *CreateIssueResponseBody
}

func (o *CreateIssueResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateIssueResponse) GetObject() *CreateIssueResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
