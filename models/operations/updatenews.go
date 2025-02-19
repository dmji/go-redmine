// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type Uploads struct {
	Token       *string `json:"token,omitempty"`
	Filename    *string `json:"filename,omitempty"`
	Description *string `json:"description,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
}

func (o *Uploads) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *Uploads) GetFilename() *string {
	if o == nil {
		return nil
	}
	return o.Filename
}

func (o *Uploads) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Uploads) GetContentType() *string {
	if o == nil {
		return nil
	}
	return o.ContentType
}

type News struct {
	Title       *string   `json:"title,omitempty"`
	Summary     *string   `json:"summary,omitempty"`
	Description *string   `json:"description,omitempty"`
	Uploads     []Uploads `json:"uploads,omitempty"`
}

func (o *News) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *News) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *News) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *News) GetUploads() []Uploads {
	if o == nil {
		return nil
	}
	return o.Uploads
}

type UpdateNewsRequestBody struct {
	News *News `json:"news,omitempty"`
}

func (o *UpdateNewsRequestBody) GetNews() *News {
	if o == nil {
		return nil
	}
	return o.News
}

type UpdateNewsRequest struct {
	Format             components.Format      `pathParam:"style=simple,explode=false,name=format"`
	NewsID             int64                  `pathParam:"style=simple,explode=false,name=news_id"`
	XRedmineSwitchUser *string                `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	RequestBody        *UpdateNewsRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateNewsRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *UpdateNewsRequest) GetNewsID() int64 {
	if o == nil {
		return 0
	}
	return o.NewsID
}

func (o *UpdateNewsRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *UpdateNewsRequest) GetRequestBody() *UpdateNewsRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type UpdateNewsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *UpdateNewsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
