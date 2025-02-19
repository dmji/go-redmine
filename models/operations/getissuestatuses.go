// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type GetIssueStatusesRequest struct {
	Format             components.Format `pathParam:"style=simple,explode=false,name=format"`
	XRedmineSwitchUser *string           `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
}

func (o *GetIssueStatusesRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *GetIssueStatusesRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

type GetIssueStatusesResponseBody struct {
	IssueStatuses []components.IssueStatus `json:"issue_statuses"`
}

func (o *GetIssueStatusesResponseBody) GetIssueStatuses() []components.IssueStatus {
	if o == nil {
		return []components.IssueStatus{}
	}
	return o.IssueStatuses
}

type GetIssueStatusesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *GetIssueStatusesResponseBody
}

func (o *GetIssueStatusesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetIssueStatusesResponse) GetObject() *GetIssueStatusesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
