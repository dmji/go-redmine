// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type GetFilesRequest struct {
	Format             components.Format `pathParam:"style=simple,explode=false,name=format"`
	ProjectID          int64             `pathParam:"style=simple,explode=false,name=project_id"`
	XRedmineSwitchUser *string           `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
}

func (o *GetFilesRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *GetFilesRequest) GetProjectID() int64 {
	if o == nil {
		return 0
	}
	return o.ProjectID
}

func (o *GetFilesRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

type GetFilesResponseBody struct {
	Files []components.File `json:"files"`
}

func (o *GetFilesResponseBody) GetFiles() []components.File {
	if o == nil {
		return []components.File{}
	}
	return o.Files
}

type GetFilesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *GetFilesResponseBody
}

func (o *GetFilesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetFilesResponse) GetObject() *GetFilesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
