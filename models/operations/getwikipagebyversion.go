// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"encoding/json"
	"fmt"

	"github.com/dmji/go-redmine/models/components"
)

type GetWikiPageByVersionQueryParamInclude string

const (
	GetWikiPageByVersionQueryParamIncludeAttachments GetWikiPageByVersionQueryParamInclude = "attachments"
)

func (e GetWikiPageByVersionQueryParamInclude) ToPointer() *GetWikiPageByVersionQueryParamInclude {
	return &e
}

func (e *GetWikiPageByVersionQueryParamInclude) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "attachments":
		*e = GetWikiPageByVersionQueryParamInclude(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetWikiPageByVersionQueryParamInclude: %v", v)
	}
}

type GetWikiPageByVersionRequest struct {
	Format             components.Format                      `pathParam:"style=simple,explode=false,name=format"`
	ProjectID          int64                                  `pathParam:"style=simple,explode=false,name=project_id"`
	WikiPageTitle      string                                 `pathParam:"style=simple,explode=false,name=wiki_page_title"`
	VersionID          int64                                  `pathParam:"style=simple,explode=false,name=version_id"`
	XRedmineSwitchUser *string                                `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	Include            *GetWikiPageByVersionQueryParamInclude `queryParam:"style=form,explode=true,name=include"`
}

func (o *GetWikiPageByVersionRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *GetWikiPageByVersionRequest) GetProjectID() int64 {
	if o == nil {
		return 0
	}
	return o.ProjectID
}

func (o *GetWikiPageByVersionRequest) GetWikiPageTitle() string {
	if o == nil {
		return ""
	}
	return o.WikiPageTitle
}

func (o *GetWikiPageByVersionRequest) GetVersionID() int64 {
	if o == nil {
		return 0
	}
	return o.VersionID
}

func (o *GetWikiPageByVersionRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *GetWikiPageByVersionRequest) GetInclude() *GetWikiPageByVersionQueryParamInclude {
	if o == nil {
		return nil
	}
	return o.Include
}

type GetWikiPageByVersionResponseBody struct {
	WikiPage components.WikiPage `json:"wiki_page"`
}

func (o *GetWikiPageByVersionResponseBody) GetWikiPage() components.WikiPage {
	if o == nil {
		return components.WikiPage{}
	}
	return o.WikiPage
}

type GetWikiPageByVersionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *GetWikiPageByVersionResponseBody
}

func (o *GetWikiPageByVersionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetWikiPageByVersionResponse) GetObject() *GetWikiPageByVersionResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
