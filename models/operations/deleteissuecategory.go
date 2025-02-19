// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"github.com/dmji/go-redmine/models/components"
)

type DeleteIssueCategoryRequest struct {
	Format             components.Format `pathParam:"style=simple,explode=false,name=format"`
	IssueCategoryID    int64             `pathParam:"style=simple,explode=false,name=issue_category_id"`
	XRedmineSwitchUser *string           `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	ReassignToID       *int64            `queryParam:"style=form,explode=true,name=reassign_to_id"`
}

func (o *DeleteIssueCategoryRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *DeleteIssueCategoryRequest) GetIssueCategoryID() int64 {
	if o == nil {
		return 0
	}
	return o.IssueCategoryID
}

func (o *DeleteIssueCategoryRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *DeleteIssueCategoryRequest) GetReassignToID() *int64 {
	if o == nil {
		return nil
	}
	return o.ReassignToID
}

type DeleteIssueCategoryResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteIssueCategoryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
