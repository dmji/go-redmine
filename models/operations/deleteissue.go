// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type DeleteIssueRequest struct {
	Format             components.Format `pathParam:"style=simple,explode=false,name=format"`
	IssueID            int64             `pathParam:"style=simple,explode=false,name=issue_id"`
	XRedmineSwitchUser *string           `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
}

func (o *DeleteIssueRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *DeleteIssueRequest) GetIssueID() int64 {
	if o == nil {
		return 0
	}
	return o.IssueID
}

func (o *DeleteIssueRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

type DeleteIssueResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteIssueResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
