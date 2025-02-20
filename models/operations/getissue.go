// Code generated by Speakeasy (https://speakeasy.com).

package operations

import (
	"encoding/json"
	"fmt"

	"github.com/dmji/go-redmine/models/components"
)

type QueryParamInclude string

const (
	QueryParamIncludeChildren        QueryParamInclude = "children"
	QueryParamIncludeAttachments     QueryParamInclude = "attachments"
	QueryParamIncludeRelations       QueryParamInclude = "relations"
	QueryParamIncludeChangesets      QueryParamInclude = "changesets"
	QueryParamIncludeJournals        QueryParamInclude = "journals"
	QueryParamIncludeWatchers        QueryParamInclude = "watchers"
	QueryParamIncludeAllowedStatuses QueryParamInclude = "allowed_statuses"
)

func (e QueryParamInclude) ToPointer() *QueryParamInclude {
	return &e
}

func (e *QueryParamInclude) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "children":
		fallthrough
	case "attachments":
		fallthrough
	case "relations":
		fallthrough
	case "changesets":
		fallthrough
	case "journals":
		fallthrough
	case "watchers":
		fallthrough
	case "allowed_statuses":
		*e = QueryParamInclude(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamInclude: %v", v)
	}
}

type GetIssueRequest struct {
	Format             components.Format   `pathParam:"style=simple,explode=false,name=format"`
	IssueID            int64               `pathParam:"style=simple,explode=false,name=issue_id"`
	XRedmineSwitchUser *string             `header:"style=simple,explode=false,name=X-Redmine-Switch-User"`
	Include            []QueryParamInclude `queryParam:"style=form,explode=false,name=include"`
}

func (o *GetIssueRequest) GetFormat() components.Format {
	if o == nil {
		return components.Format("")
	}
	return o.Format
}

func (o *GetIssueRequest) GetIssueID() int64 {
	if o == nil {
		return 0
	}
	return o.IssueID
}

func (o *GetIssueRequest) GetXRedmineSwitchUser() *string {
	if o == nil {
		return nil
	}
	return o.XRedmineSwitchUser
}

func (o *GetIssueRequest) GetInclude() []QueryParamInclude {
	if o == nil {
		return nil
	}
	return o.Include
}

type GetIssueResponseBody struct {
	Issue components.Issue `json:"issue"`
}

func (o *GetIssueResponseBody) GetIssue() components.Issue {
	if o == nil {
		return components.Issue{}
	}
	return o.Issue
}

type GetIssueResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *GetIssueResponseBody
}

func (o *GetIssueResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetIssueResponse) GetObject() *GetIssueResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
