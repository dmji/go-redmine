// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"encoding/json"
	"fmt"

	"github.com/dmji/go-redmine/models/components"
)

type GetNewsQueryParamInclude string

const (
	GetNewsQueryParamIncludeAttachments GetNewsQueryParamInclude = "attachments"
	GetNewsQueryParamIncludeComments    GetNewsQueryParamInclude = "comments"
)

func (e GetNewsQueryParamInclude) ToPointer() *GetNewsQueryParamInclude {
	return &e
}

func (e *GetNewsQueryParamInclude) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "attachments":
		fallthrough
	case "comments":
		*e = GetNewsQueryParamInclude(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetNewsQueryParamInclude: %v", v)
	}
}

type GetNewsRequest struct {
	Format             components.Format          `pathParam:"style=simple,explode=false,name=format"`
	NewsID             int64                      `pathParam:"style=simple,explode=false,name=news_id"`
	XRedmineSwitchUser *string                    `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	Include            []GetNewsQueryParamInclude `queryParam:"style=form,explode=false,name=include"`
}

func (o *GetNewsRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *GetNewsRequest) GetNewsID() int64 {
	if o == nil {
		return 0
	}
	return o.NewsID
}

func (o *GetNewsRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *GetNewsRequest) GetInclude() []GetNewsQueryParamInclude {
	if o == nil {
		return nil
	}
	return o.Include
}

type GetNewsResponseBody struct {
	News components.News `json:"news"`
}

func (o *GetNewsResponseBody) GetNews() components.News {
	if o == nil {
		return components.News{}
	}
	return o.News
}

type GetNewsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *GetNewsResponseBody
}

func (o *GetNewsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetNewsResponse) GetObject() *GetNewsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
