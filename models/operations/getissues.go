// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"encoding/json"
	"fmt"

	"github.com/dmji/go-redmine/models/components"
)

type Include string

const (
	IncludeAttachments Include = "attachments"
	IncludeRelations   Include = "relations"
)

func (e Include) ToPointer() *Include {
	return &e
}

func (e *Include) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "attachments":
		fallthrough
	case "relations":
		*e = Include(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Include: %v", v)
	}
}

type GetIssuesRequest struct {
	Format                components.Format          `pathParam:"style=simple,explode=false,name=format"`
	Offset                *int64                     `queryParam:"style=form,explode=true,name=offset"`
	Limit                 *int64                     `queryParam:"style=form,explode=true,name=limit"`
	Nometa                *components.Nometa         `queryParam:"style=form,explode=true,name=nometa"`
	XRedmineSwitchUser    *string                    `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	XRedmineNometa        *components.XRedmineNometa `header:"style=simple,explode=false,name=X-Redmine-Nometa"`
	Sort                  *string                    `queryParam:"style=form,explode=true,name=sort"`
	Include               []Include                  `queryParam:"style=form,explode=false,name=include"`
	IssueID               []string                   `queryParam:"style=form,explode=false,name=issue_id"`
	ProjectID             []string                   `queryParam:"style=form,explode=false,name=project_id"`
	SubprojectID          []string                   `queryParam:"style=form,explode=false,name=subproject_id"`
	TrackerID             []string                   `queryParam:"style=form,explode=false,name=tracker_id"`
	StatusID              []string                   `queryParam:"style=form,explode=false,name=status_id"`
	AssignedToID          []string                   `queryParam:"style=form,explode=false,name=assigned_to_id"`
	ParentID              []string                   `queryParam:"style=form,explode=false,name=parent_id"`
	CfX                   map[string]string          `queryParam:"style=form,explode=true,name=cf_x"`
	AuthorID              []string                   `queryParam:"style=form,explode=false,name=author_id"`
	MemberOfGroup         []string                   `queryParam:"style=form,explode=false,name=member_of_group"`
	AssignedToRole        []string                   `queryParam:"style=form,explode=false,name=assigned_to_role"`
	FixedVersionID        []string                   `queryParam:"style=form,explode=false,name=fixed_version_id"`
	FixedVersionDueDate   *string                    `queryParam:"style=form,explode=true,name=fixed_version.due_date"`
	FixedVersionStatus    []string                   `queryParam:"style=form,explode=false,name=fixed_version.status"`
	CategoryID            []string                   `queryParam:"style=form,explode=false,name=category_id"`
	Subject               *string                    `queryParam:"style=form,explode=true,name=subject"`
	Description           *string                    `queryParam:"style=form,explode=true,name=description"`
	Notes                 *string                    `queryParam:"style=form,explode=true,name=notes"`
	CreatedOn             *string                    `queryParam:"style=form,explode=true,name=created_on"`
	UpdatedOn             *string                    `queryParam:"style=form,explode=true,name=updated_on"`
	ClosedOn              *string                    `queryParam:"style=form,explode=true,name=closed_on"`
	StartDate             *string                    `queryParam:"style=form,explode=true,name=start_date"`
	DueDate               *string                    `queryParam:"style=form,explode=true,name=due_date"`
	EstimatedHours        *string                    `queryParam:"style=form,explode=true,name=estimated_hours"`
	SpentTime             *string                    `queryParam:"style=form,explode=true,name=spent_time"`
	DoneRatio             *string                    `queryParam:"style=form,explode=true,name=done_ratio"`
	IsPrivate             *string                    `queryParam:"style=form,explode=true,name=is_private"`
	Attachment            *string                    `queryParam:"style=form,explode=true,name=attachment"`
	AttachmentDescription *string                    `queryParam:"style=form,explode=true,name=attachment_description"`
	WatcherID             []string                   `queryParam:"style=form,explode=false,name=watcher_id"`
	UpdatedBy             []string                   `queryParam:"style=form,explode=false,name=updated_by"`
	LastUpdatedBy         []string                   `queryParam:"style=form,explode=false,name=last_updated_by"`
	ProjectStatus         *int64                     `queryParam:"style=form,explode=true,name=project.status"`
	RelationType          []string                   `queryParam:"style=form,explode=false,name=relation_type"`
	ChildID               []string                   `queryParam:"style=form,explode=false,name=child_id"`
	QueryID               *int64                     `queryParam:"style=form,explode=true,name=query_id"`
}

func (o *GetIssuesRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *GetIssuesRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetIssuesRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetIssuesRequest) GetNometa() *components.Nometa {
	if o == nil {
		return nil
	}
	return o.Nometa
}

func (o *GetIssuesRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *GetIssuesRequest) GetXRedmineNometa() *components.XRedmineNometa {
	if o == nil {
		return nil
	}
	return o.XRedmineNometa
}

func (o *GetIssuesRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetIssuesRequest) GetInclude() []Include {
	if o == nil {
		return nil
	}
	return o.Include
}

func (o *GetIssuesRequest) GetIssueID() []string {
	if o == nil {
		return nil
	}
	return o.IssueID
}

func (o *GetIssuesRequest) GetProjectID() []string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *GetIssuesRequest) GetSubprojectID() []string {
	if o == nil {
		return nil
	}
	return o.SubprojectID
}

func (o *GetIssuesRequest) GetTrackerID() []string {
	if o == nil {
		return nil
	}
	return o.TrackerID
}

func (o *GetIssuesRequest) GetStatusID() []string {
	if o == nil {
		return nil
	}
	return o.StatusID
}

func (o *GetIssuesRequest) GetAssignedToID() []string {
	if o == nil {
		return nil
	}
	return o.AssignedToID
}

func (o *GetIssuesRequest) GetParentID() []string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *GetIssuesRequest) GetCfX() map[string]string {
	if o == nil {
		return nil
	}
	return o.CfX
}

func (o *GetIssuesRequest) GetAuthorID() []string {
	if o == nil {
		return nil
	}
	return o.AuthorID
}

func (o *GetIssuesRequest) GetMemberOfGroup() []string {
	if o == nil {
		return nil
	}
	return o.MemberOfGroup
}

func (o *GetIssuesRequest) GetAssignedToRole() []string {
	if o == nil {
		return nil
	}
	return o.AssignedToRole
}

func (o *GetIssuesRequest) GetFixedVersionID() []string {
	if o == nil {
		return nil
	}
	return o.FixedVersionID
}

func (o *GetIssuesRequest) GetFixedVersionDueDate() *string {
	if o == nil {
		return nil
	}
	return o.FixedVersionDueDate
}

func (o *GetIssuesRequest) GetFixedVersionStatus() []string {
	if o == nil {
		return nil
	}
	return o.FixedVersionStatus
}

func (o *GetIssuesRequest) GetCategoryID() []string {
	if o == nil {
		return nil
	}
	return o.CategoryID
}

func (o *GetIssuesRequest) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *GetIssuesRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *GetIssuesRequest) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}

func (o *GetIssuesRequest) GetCreatedOn() *string {
	if o == nil {
		return nil
	}
	return o.CreatedOn
}

func (o *GetIssuesRequest) GetUpdatedOn() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedOn
}

func (o *GetIssuesRequest) GetClosedOn() *string {
	if o == nil {
		return nil
	}
	return o.ClosedOn
}

func (o *GetIssuesRequest) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetIssuesRequest) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *GetIssuesRequest) GetEstimatedHours() *string {
	if o == nil {
		return nil
	}
	return o.EstimatedHours
}

func (o *GetIssuesRequest) GetSpentTime() *string {
	if o == nil {
		return nil
	}
	return o.SpentTime
}

func (o *GetIssuesRequest) GetDoneRatio() *string {
	if o == nil {
		return nil
	}
	return o.DoneRatio
}

func (o *GetIssuesRequest) GetIsPrivate() *string {
	if o == nil {
		return nil
	}
	return o.IsPrivate
}

func (o *GetIssuesRequest) GetAttachment() *string {
	if o == nil {
		return nil
	}
	return o.Attachment
}

func (o *GetIssuesRequest) GetAttachmentDescription() *string {
	if o == nil {
		return nil
	}
	return o.AttachmentDescription
}

func (o *GetIssuesRequest) GetWatcherID() []string {
	if o == nil {
		return nil
	}
	return o.WatcherID
}

func (o *GetIssuesRequest) GetUpdatedBy() []string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *GetIssuesRequest) GetLastUpdatedBy() []string {
	if o == nil {
		return nil
	}
	return o.LastUpdatedBy
}

func (o *GetIssuesRequest) GetProjectStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.ProjectStatus
}

func (o *GetIssuesRequest) GetRelationType() []string {
	if o == nil {
		return nil
	}
	return o.RelationType
}

func (o *GetIssuesRequest) GetChildID() []string {
	if o == nil {
		return nil
	}
	return o.ChildID
}

func (o *GetIssuesRequest) GetQueryID() *int64 {
	if o == nil {
		return nil
	}
	return o.QueryID
}

type GetIssuesResponseBody struct {
	Issues     []components.IssueSimple `json:"issues"`
	TotalCount *int64                   `json:"total_count,omitempty"`
	Offset     *int64                   `json:"offset,omitempty"`
	Limit      *int64                   `json:"limit,omitempty"`
}

func (o *GetIssuesResponseBody) GetIssues() []components.IssueSimple {
	if o == nil {
		return []components.IssueSimple{}
	}
	return o.Issues
}

func (o *GetIssuesResponseBody) GetTotalCount() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalCount
}

func (o *GetIssuesResponseBody) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetIssuesResponseBody) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type GetIssuesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *GetIssuesResponseBody
}

func (o *GetIssuesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetIssuesResponse) GetObject() *GetIssuesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
