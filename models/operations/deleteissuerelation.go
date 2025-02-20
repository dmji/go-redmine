// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type DeleteIssueRelationRequest struct {
	Format             components.Format `pathParam:"style=simple,explode=false,name=format"`
	IssueRelationID    int64             `pathParam:"style=simple,explode=false,name=issue_relation_id"`
	XRedmineSwitchUser *string           `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
}

func (o *DeleteIssueRelationRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *DeleteIssueRelationRequest) GetIssueRelationID() int64 {
	if o == nil {
		return 0
	}
	return o.IssueRelationID
}

func (o *DeleteIssueRelationRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

type DeleteIssueRelationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteIssueRelationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
