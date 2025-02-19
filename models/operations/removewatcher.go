// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type RemoveWatcherRequest struct {
	Format             components.Format `pathParam:"style=simple,explode=false,name=format"`
	IssueID            int64             `pathParam:"style=simple,explode=false,name=issue_id"`
	UserID             int64             `pathParam:"style=simple,explode=false,name=user_id"`
	XRedmineSwitchUser *string           `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
}

func (o *RemoveWatcherRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *RemoveWatcherRequest) GetIssueID() int64 {
	if o == nil {
		return 0
	}
	return o.IssueID
}

func (o *RemoveWatcherRequest) GetUserID() int64 {
	if o == nil {
		return 0
	}
	return o.UserID
}

func (o *RemoveWatcherRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

type RemoveWatcherResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *RemoveWatcherResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
