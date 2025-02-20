// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type GetUsersRequest struct {
	Format             components.Format          `pathParam:"style=simple,explode=false,name=format"`
	Offset             *int64                     `queryParam:"style=form,explode=true,name=offset"`
	Limit              *int64                     `queryParam:"style=form,explode=true,name=limit"`
	Nometa             *components.Nometa         `queryParam:"style=form,explode=true,name=nometa"`
	XRedmineSwitchUser *string                    `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	XRedmineNometa     *components.XRedmineNometa `header:"style=simple,explode=false,name=X-Redmine-Nometa"`
	Status             *int64                     `queryParam:"style=form,explode=true,name=status"`
	Name               *string                    `queryParam:"style=form,explode=true,name=name"`
	GroupID            *int64                     `queryParam:"style=form,explode=true,name=group_id"`
}

func (o *GetUsersRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *GetUsersRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetUsersRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetUsersRequest) GetNometa() *components.Nometa {
	if o == nil {
		return nil
	}
	return o.Nometa
}

func (o *GetUsersRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *GetUsersRequest) GetXRedmineNometa() *components.XRedmineNometa {
	if o == nil {
		return nil
	}
	return o.XRedmineNometa
}

func (o *GetUsersRequest) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetUsersRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetUsersRequest) GetGroupID() *int64 {
	if o == nil {
		return nil
	}
	return o.GroupID
}

type GetUsersResponseBody struct {
	Users      []components.UserSimple `json:"users"`
	TotalCount *int64                  `json:"total_count,omitempty"`
	Offset     *int64                  `json:"offset,omitempty"`
	Limit      *int64                  `json:"limit,omitempty"`
}

func (o *GetUsersResponseBody) GetUsers() []components.UserSimple {
	if o == nil {
		return []components.UserSimple{}
	}
	return o.Users
}

func (o *GetUsersResponseBody) GetTotalCount() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalCount
}

func (o *GetUsersResponseBody) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetUsersResponseBody) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type GetUsersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *GetUsersResponseBody
}

func (o *GetUsersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUsersResponse) GetObject() *GetUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
