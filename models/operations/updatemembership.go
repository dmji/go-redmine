// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type UpdateMembershipMembership struct {
	RoleIds []int64 `json:"role_ids"`
}

func (o *UpdateMembershipMembership) GetRoleIds() []int64 {
	if o == nil {
		return []int64{}
	}
	return o.RoleIds
}

type UpdateMembershipRequestBody struct {
	Membership *UpdateMembershipMembership `json:"membership,omitempty"`
}

func (o *UpdateMembershipRequestBody) GetMembership() *UpdateMembershipMembership {
	if o == nil {
		return nil
	}
	return o.Membership
}

type UpdateMembershipRequest struct {
	Format             components.Format            `pathParam:"style=simple,explode=false,name=format"`
	MembershipID       int64                        `pathParam:"style=simple,explode=false,name=membership_id"`
	XRedmineSwitchUser *string                      `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	RequestBody        *UpdateMembershipRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateMembershipRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *UpdateMembershipRequest) GetMembershipID() int64 {
	if o == nil {
		return 0
	}
	return o.MembershipID
}

func (o *UpdateMembershipRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *UpdateMembershipRequest) GetRequestBody() *UpdateMembershipRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type UpdateMembershipResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *UpdateMembershipResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
