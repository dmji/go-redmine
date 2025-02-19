// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type DeleteVersionRequest struct {
	Format             components.Format `pathParam:"style=simple,explode=false,name=format"`
	VersionID          int64             `pathParam:"style=simple,explode=false,name=version_id"`
	XRedmineSwitchUser *string           `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
}

func (o *DeleteVersionRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *DeleteVersionRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

func (o *DeleteVersionRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

type DeleteVersionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteVersionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
