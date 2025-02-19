// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type GetCustomFieldsRequest struct {
	Format             components.Format `pathParam:"style=simple,explode=false,name=format"`
	XRedmineSwitchUser *string           `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
}

func (o *GetCustomFieldsRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *GetCustomFieldsRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

type GetCustomFieldsResponseBody struct {
	CustomFields []components.CustomField `json:"custom_fields"`
}

func (o *GetCustomFieldsResponseBody) GetCustomFields() []components.CustomField {
	if o == nil {
		return []components.CustomField{}
	}
	return o.CustomFields
}

type GetCustomFieldsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *GetCustomFieldsResponseBody
}

func (o *GetCustomFieldsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCustomFieldsResponse) GetObject() *GetCustomFieldsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
