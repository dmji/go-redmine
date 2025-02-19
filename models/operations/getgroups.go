// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type GetGroupsRequest struct {
	Format             components.Format `pathParam:"style=simple,explode=false,name=format"`
	XRedmineSwitchUser *string           `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
}

func (o *GetGroupsRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *GetGroupsRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

type GetGroupsResponseBody struct {
	Groups []components.GroupSimple `json:"groups"`
}

func (o *GetGroupsResponseBody) GetGroups() []components.GroupSimple {
	if o == nil {
		return []components.GroupSimple{}
	}
	return o.Groups
}

type GetGroupsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *GetGroupsResponseBody
}

func (o *GetGroupsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetGroupsResponse) GetObject() *GetGroupsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
