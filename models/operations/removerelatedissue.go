// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type RemoveRelatedIssueRequest struct {
	Format             components.Format `pathParam:"style=simple,explode=false,name=format"`
	ProjectID          int64             `pathParam:"style=simple,explode=false,name=project_id"`
	RepositoryID       int64             `pathParam:"style=simple,explode=false,name=repository_id"`
	Revision           string            `pathParam:"style=simple,explode=false,name=revision"`
	IssueID            int64             `pathParam:"style=simple,explode=false,name=issue_id"`
	XRedmineSwitchUser *string           `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
}

func (o *RemoveRelatedIssueRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *RemoveRelatedIssueRequest) GetProjectID() int64 {
	if o == nil {
		return 0
	}
	return o.ProjectID
}

func (o *RemoveRelatedIssueRequest) GetRepositoryID() int64 {
	if o == nil {
		return 0
	}
	return o.RepositoryID
}

func (o *RemoveRelatedIssueRequest) GetRevision() string {
	if o == nil {
		return ""
	}
	return o.Revision
}

func (o *RemoveRelatedIssueRequest) GetIssueID() int64 {
	if o == nil {
		return 0
	}
	return o.IssueID
}

func (o *RemoveRelatedIssueRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

type RemoveRelatedIssueResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *RemoveRelatedIssueResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
