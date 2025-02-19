// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type CreateNewsUploads struct {
	Token       *string `json:"token,omitempty"`
	Filename    *string `json:"filename,omitempty"`
	Description *string `json:"description,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
}

func (o *CreateNewsUploads) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *CreateNewsUploads) GetFilename() *string {
	if o == nil {
		return nil
	}
	return o.Filename
}

func (o *CreateNewsUploads) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateNewsUploads) GetContentType() *string {
	if o == nil {
		return nil
	}
	return o.ContentType
}

type CreateNewsNews struct {
	Title       *string             `json:"title,omitempty"`
	Summary     *string             `json:"summary,omitempty"`
	Description *string             `json:"description,omitempty"`
	Uploads     []CreateNewsUploads `json:"uploads,omitempty"`
}

func (o *CreateNewsNews) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *CreateNewsNews) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *CreateNewsNews) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateNewsNews) GetUploads() []CreateNewsUploads {
	if o == nil {
		return nil
	}
	return o.Uploads
}

type CreateNewsRequestBody struct {
	News CreateNewsNews `json:"news"`
}

func (o *CreateNewsRequestBody) GetNews() CreateNewsNews {
	if o == nil {
		return CreateNewsNews{}
	}
	return o.News
}

type CreateNewsRequest struct {
	Format             components.Format      `pathParam:"style=simple,explode=false,name=format"`
	ProjectID          int64                  `pathParam:"style=simple,explode=false,name=project_id"`
	XRedmineSwitchUser *string                `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	RequestBody        *CreateNewsRequestBody `request:"mediaType=application/json"`
}

func (o *CreateNewsRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *CreateNewsRequest) GetProjectID() int64 {
	if o == nil {
		return 0
	}
	return o.ProjectID
}

func (o *CreateNewsRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *CreateNewsRequest) GetRequestBody() *CreateNewsRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type CreateNewsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *CreateNewsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
