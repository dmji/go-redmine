// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type Attachment struct {
	Filename    *string `json:"filename,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (o *Attachment) GetFilename() *string {
	if o == nil {
		return nil
	}
	return o.Filename
}

func (o *Attachment) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

type UpdateAttachmentRequestBody struct {
	Attachment *Attachment `json:"attachment,omitempty"`
}

func (o *UpdateAttachmentRequestBody) GetAttachment() *Attachment {
	if o == nil {
		return nil
	}
	return o.Attachment
}

type UpdateAttachmentRequest struct {
	Format             components.Format            `pathParam:"style=simple,explode=false,name=format"`
	AttachmentID       float64                      `pathParam:"style=simple,explode=false,name=attachment_id"`
	XRedmineSwitchUser *string                      `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	RequestBody        *UpdateAttachmentRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateAttachmentRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *UpdateAttachmentRequest) GetAttachmentID() float64 {
	if o == nil {
		return 0.0
	}
	return o.AttachmentID
}

func (o *UpdateAttachmentRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *UpdateAttachmentRequest) GetRequestBody() *UpdateAttachmentRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type UpdateAttachmentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *UpdateAttachmentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
