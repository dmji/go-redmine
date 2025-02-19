// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type GetVersionsByProjectRequest struct {
	Format             components.Format          `pathParam:"style=simple,explode=false,name=format"`
	ProjectID          int64                      `pathParam:"style=simple,explode=false,name=project_id"`
	Nometa             *components.Nometa         `queryParam:"style=form,explode=true,name=nometa"`
	XRedmineSwitchUser *string                    `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	XRedmineNometa     *components.XRedmineNometa `header:"style=simple,explode=false,name=X-Redmine-Nometa"`
}

func (o *GetVersionsByProjectRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *GetVersionsByProjectRequest) GetProjectID() int64 {
	if o == nil {
		return 0
	}
	return o.ProjectID
}

func (o *GetVersionsByProjectRequest) GetNometa() *components.Nometa {
	if o == nil {
		return nil
	}
	return o.Nometa
}

func (o *GetVersionsByProjectRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *GetVersionsByProjectRequest) GetXRedmineNometa() *components.XRedmineNometa {
	if o == nil {
		return nil
	}
	return o.XRedmineNometa
}

type GetVersionsByProjectResponseBody struct {
	Versions   []components.VersionSimple `json:"versions"`
	TotalCount *int64                     `json:"total_count,omitempty"`
}

func (o *GetVersionsByProjectResponseBody) GetVersions() []components.VersionSimple {
	if o == nil {
		return []components.VersionSimple{}
	}
	return o.Versions
}

func (o *GetVersionsByProjectResponseBody) GetTotalCount() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalCount
}

type GetVersionsByProjectResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *GetVersionsByProjectResponseBody
}

func (o *GetVersionsByProjectResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetVersionsByProjectResponse) GetObject() *GetVersionsByProjectResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
