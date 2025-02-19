// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type DeleteGroupRequest struct {
	Format             components.Format `pathParam:"style=simple,explode=false,name=format"`
	GroupID            int64             `pathParam:"style=simple,explode=false,name=group_id"`
	XRedmineSwitchUser *string           `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
}

func (o *DeleteGroupRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *DeleteGroupRequest) GetGroupID() int64 {
	if o == nil {
		return 0
	}
	return o.GroupID
}

func (o *DeleteGroupRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

type DeleteGroupResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteGroupResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
