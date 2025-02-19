# github.com/dmji/go-redmine

Developer-friendly & type-safe Go SDK specifically catered to leverage _github.com/dmji/go-redmine_ API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/dmji/go-redmine&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>

`<br /><br />`

## Summary

Redmine API: Unofficial OpenAPI specification for the Redmine API.

For more information about the API: [Redmine API Official Developer Guide](https://www.redmine.org/projects/redmine/wiki/Rest_api)

<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->

## Table of Contents

<!-- $toc-max-depth=2 -->

- [github.com/dmji/go-redmine](#githubcomdmjigo-redmine)
  - [SDK Installation](#sdk-installation)
  - [SDK Example Usage](#sdk-example-usage)
  - [Authentication](#authentication)
  - [Available Resources and Operations](#available-resources-and-operations)
  - [Retries](#retries)
  - [Error Handling](#error-handling)
  - [Server Selection](#server-selection)
  - [Custom HTTP Client](#custom-http-client)
  - [Special Types](#special-types)
- [Development](#development)
  - [Maturity](#maturity)
  - [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->

## SDK Installation

To add the SDK as a dependency to your project:

```bash
go get github.com/dmji/go-redmine
```

<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->

## SDK Example Usage

### Example

```go
package main

import (
	"context"
	goredmine "github.com/dmji/go-redmine"
	"github.com/dmji/go-redmine/models/components"
	"github.com/dmji/go-redmine/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := goredmine.New(
		goredmine.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Username: "",
				Password: "",
			},
		}),
	)

	res, err := s.Issues.GetIssues(ctx, operations.GetIssuesRequest{
		Format:             components.FormatXML,
		XRedmineSwitchUser: goredmine.String("jsmith"),
		Sort:               goredmine.String("id:desc"),
		CfX: map[string]string{
			"cf_0": "string",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->

## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name                | Type   | Scheme     | Environment Variable              |
| ------------------- | ------ | ---------- | --------------------------------- |
| `BasicAuth`         | http   | HTTP Basic | `GOREDMINE_BASIC_AUTH`            |
| `APIKeyAuth`        | apiKey | API key    | `GOREDMINE_API_KEY_AUTH`          |
| `APIKeyInQueryAuth` | apiKey | API key    | `GOREDMINE_API_KEY_IN_QUERY_AUTH` |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:

```go
package main

import (
	"context"
	goredmine "github.com/dmji/go-redmine"
	"github.com/dmji/go-redmine/models/components"
	"github.com/dmji/go-redmine/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := goredmine.New(
		goredmine.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Username: "",
				Password: "",
			},
		}),
	)

	res, err := s.Issues.GetIssues(ctx, operations.GetIssuesRequest{
		Format:             components.FormatXML,
		XRedmineSwitchUser: goredmine.String("jsmith"),
		Sort:               goredmine.String("id:desc"),
		CfX: map[string]string{
			"cf_0": "string",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->

## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Attachments](docs/sdks/attachments/README.md)

- [GetAttachment](docs/sdks/attachments/README.md#getattachment) - Show attachment
- [UpdateAttachment](docs/sdks/attachments/README.md#updateattachment) - Update attachment
- [DeleteAttachment](docs/sdks/attachments/README.md#deleteattachment) - Delete attachment
- [DownloadAttachmentFile](docs/sdks/attachments/README.md#downloadattachmentfile) - Download attachment file
- [DownloadThumbnail](docs/sdks/attachments/README.md#downloadthumbnail) - Download thumbnail
- [UploadAttachmentFile](docs/sdks/attachments/README.md#uploadattachmentfile) - Upload attachment file

### [CustomFields](docs/sdks/customfields/README.md)

- [GetCustomFields](docs/sdks/customfields/README.md#getcustomfields) - List custom fields

### [Enumerations](docs/sdks/enumerations/README.md)

- [GetIssuePriorities](docs/sdks/enumerations/README.md#getissuepriorities) - List issue priorities
- [GetTimeEntryActivities](docs/sdks/enumerations/README.md#gettimeentryactivities) - List time entry activities
- [GetDocumentCategories](docs/sdks/enumerations/README.md#getdocumentcategories) - List document categories

### [Files](docs/sdks/files/README.md)

- [GetFiles](docs/sdks/files/README.md#getfiles) - List files
- [CreateFile](docs/sdks/files/README.md#createfile) - Create file

### [Groups](docs/sdks/groups/README.md)

- [GetGroups](docs/sdks/groups/README.md#getgroups) - List groups
- [CreateGroup](docs/sdks/groups/README.md#creategroup) - Create group
- [GetGroup](docs/sdks/groups/README.md#getgroup) - Show group
- [UpdateGroup](docs/sdks/groups/README.md#updategroup) - Update group
- [DeleteGroup](docs/sdks/groups/README.md#deletegroup) - Delete group
- [AddUserToGroup](docs/sdks/groups/README.md#addusertogroup) - Add user to group
- [RemoveUserFromGroup](docs/sdks/groups/README.md#removeuserfromgroup) - Remove user from group

### [IssueCategories](docs/sdks/issuecategories/README.md)

- [GetIssueCategories](docs/sdks/issuecategories/README.md#getissuecategories) - List issue categories
- [CreateIssueCategory](docs/sdks/issuecategories/README.md#createissuecategory) - Create issue category
- [GetIssueCategory](docs/sdks/issuecategories/README.md#getissuecategory) - Show issue category
- [UpdateIssueCategory](docs/sdks/issuecategories/README.md#updateissuecategory) - Update issue category
- [DeleteIssueCategory](docs/sdks/issuecategories/README.md#deleteissuecategory) - Delete issue category

### [IssueRelations](docs/sdks/issuerelations/README.md)

- [GetIssueRelations](docs/sdks/issuerelations/README.md#getissuerelations) - List issue relations
- [CreateIssueRelation](docs/sdks/issuerelations/README.md#createissuerelation) - Create issue relation
- [GetIssueRelation](docs/sdks/issuerelations/README.md#getissuerelation) - Show issue relation
- [DeleteIssueRelation](docs/sdks/issuerelations/README.md#deleteissuerelation) - Delete issue relation

### [Issues](docs/sdks/issues/README.md)

- [GetIssues](docs/sdks/issues/README.md#getissues) - List issues
- [CreateIssue](docs/sdks/issues/README.md#createissue) - Create issue
- [GetIssue](docs/sdks/issues/README.md#getissue) - Show issue
- [UpdateIssue](docs/sdks/issues/README.md#updateissue) - Update issue
- [DeleteIssue](docs/sdks/issues/README.md#deleteissue) - Delete issue
- [AddWatcher](docs/sdks/issues/README.md#addwatcher) - Add watcher
- [RemoveWatcher](docs/sdks/issues/README.md#removewatcher) - Remove watcher

### [IssueStatuses](docs/sdks/issuestatuses/README.md)

- [GetIssueStatuses](docs/sdks/issuestatuses/README.md#getissuestatuses) - List issue statuses

### [Journals](docs/sdks/journals/README.md)

- [UpdateJournal](docs/sdks/journals/README.md#updatejournal) - Update journal

### [MyAccount](docs/sdks/myaccount/README.md)

- [GetMyAccount](docs/sdks/myaccount/README.md#getmyaccount) - Show my account
- [UpdateMyAccount](docs/sdks/myaccount/README.md#updatemyaccount) - Update my account

### [News](docs/sdks/news/README.md)

- [GetNewsList](docs/sdks/news/README.md#getnewslist) - List news
- [GetNews](docs/sdks/news/README.md#getnews) - Show news
- [UpdateNews](docs/sdks/news/README.md#updatenews) - Update news
- [DeleteNews](docs/sdks/news/README.md#deletenews) - Delete news
- [GetNewsListByProject](docs/sdks/news/README.md#getnewslistbyproject) - List news by project
- [CreateNews](docs/sdks/news/README.md#createnews) - Create news

### [ProjectMemberships](docs/sdks/projectmemberships/README.md)

- [GetMemberships](docs/sdks/projectmemberships/README.md#getmemberships) - List memberships
- [CreateMembership](docs/sdks/projectmemberships/README.md#createmembership) - Create membership
- [GetMembership](docs/sdks/projectmemberships/README.md#getmembership) - Show membership
- [UpdateMembership](docs/sdks/projectmemberships/README.md#updatemembership) - Update membership
- [DeleteMembership](docs/sdks/projectmemberships/README.md#deletemembership) - Delete membership

### [Projects](docs/sdks/projects/README.md)

- [GetProjects](docs/sdks/projects/README.md#getprojects) - List projects
- [CreateProject](docs/sdks/projects/README.md#createproject) - Crete project
- [GetProject](docs/sdks/projects/README.md#getproject) - Show project
- [UpdateProject](docs/sdks/projects/README.md#updateproject) - Update project
- [DeleteProject](docs/sdks/projects/README.md#deleteproject) - Delete project
- [ArchiveProject](docs/sdks/projects/README.md#archiveproject) - Archive project
- [UnarchiveProject](docs/sdks/projects/README.md#unarchiveproject) - Unarchive project
- [CloseProject](docs/sdks/projects/README.md#closeproject) - Close project
- [ReopenProject](docs/sdks/projects/README.md#reopenproject) - Reopen project

### [Queries](docs/sdks/queries/README.md)

- [GetQueries](docs/sdks/queries/README.md#getqueries) - List queries

### [Repositories](docs/sdks/repositories/README.md)

- [AddRelatedIssue](docs/sdks/repositories/README.md#addrelatedissue) - Add related issue
- [RemoveRelatedIssue](docs/sdks/repositories/README.md#removerelatedissue) - Remove related issue

### [Roles](docs/sdks/roles/README.md)

- [GetRoles](docs/sdks/roles/README.md#getroles) - List roles
- [GetRole](docs/sdks/roles/README.md#getrole) - Show role

### [Search](docs/sdks/search/README.md)

- [Search](docs/sdks/search/README.md#search) - Search

### [TimeEntries](docs/sdks/timeentries/README.md)

- [GetTimeEntries](docs/sdks/timeentries/README.md#gettimeentries) - List time entries
- [CreateTimeEntry](docs/sdks/timeentries/README.md#createtimeentry) - Create time entry
- [GetTimeEntry](docs/sdks/timeentries/README.md#gettimeentry) - Show time entry
- [UpdateTimeEntry](docs/sdks/timeentries/README.md#updatetimeentry) - Update time entry
- [DeleteTimeEntry](docs/sdks/timeentries/README.md#deletetimeentry) - Delete time entry

### [Trackers](docs/sdks/trackers/README.md)

- [GetTrackers](docs/sdks/trackers/README.md#gettrackers) - List trackers

### [Users](docs/sdks/users/README.md)

- [GetUsers](docs/sdks/users/README.md#getusers) - List users
- [CreateUser](docs/sdks/users/README.md#createuser) - Create user
- [GetUser](docs/sdks/users/README.md#getuser) - Show user
- [UpdateUser](docs/sdks/users/README.md#updateuser) - Update user
- [DeleteUser](docs/sdks/users/README.md#deleteuser) - Delete user
- [GetCurrentUser](docs/sdks/users/README.md#getcurrentuser) - Show current user

### [Versions](docs/sdks/versions/README.md)

- [GetVersionsByProject](docs/sdks/versions/README.md#getversionsbyproject) - List versions by project
- [CreateVersion](docs/sdks/versions/README.md#createversion) - Create version
- [GetVersions](docs/sdks/versions/README.md#getversions) - Show version
- [UpdateVersion](docs/sdks/versions/README.md#updateversion) - Update version
- [DeleteVersion](docs/sdks/versions/README.md#deleteversion) - Delete version

### [WikiPages](docs/sdks/wikipages/README.md)

- [GetWikiPages](docs/sdks/wikipages/README.md#getwikipages) - List wiki pages
- [GetWikiPage](docs/sdks/wikipages/README.md#getwikipage) - Show wiki page
- [UpdateWikiPage](docs/sdks/wikipages/README.md#updatewikipage) - Create or update wiki page
- [DeleteWikiPage](docs/sdks/wikipages/README.md#deletewikipage) - Delete wiki page
- [GetWikiPageByVersion](docs/sdks/wikipages/README.md#getwikipagebyversion) - Show wiki page by specific version

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->

## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:

```go
package main

import (
	"context"
	goredmine "github.com/dmji/go-redmine"
	"github.com/dmji/go-redmine/models/components"
	"github.com/dmji/go-redmine/models/operations"
	"github.com/dmji/go-redmine/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := goredmine.New(
		goredmine.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Username: "",
				Password: "",
			},
		}),
	)

	res, err := s.Issues.GetIssues(ctx, operations.GetIssuesRequest{
		Format:             components.FormatXML,
		XRedmineSwitchUser: goredmine.String("jsmith"),
		Sort:               goredmine.String("id:desc"),
		CfX: map[string]string{
			"cf_0": "string",
		},
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:

```go
package main

import (
	"context"
	goredmine "github.com/dmji/go-redmine"
	"github.com/dmji/go-redmine/models/components"
	"github.com/dmji/go-redmine/models/operations"
	"github.com/dmji/go-redmine/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := goredmine.New(
		goredmine.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		goredmine.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Username: "",
				Password: "",
			},
		}),
	)

	res, err := s.Issues.GetIssues(ctx, operations.GetIssuesRequest{
		Format:             components.FormatXML,
		XRedmineSwitchUser: goredmine.String("jsmith"),
		Sort:               goredmine.String("id:desc"),
		CfX: map[string]string{
			"cf_0": "string",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->

## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective _Errors_ tables in SDK docs for more details on possible error types for each operation.

For example, the `CreateProject` function may return the following errors:

| Error Type         | Status Code | Content Type     |
| ------------------ | ----------- | ---------------- |
| apierrors.Errors   | 422         | application/json |
| apierrors.APIError | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	goredmine "github.com/dmji/go-redmine"
	"github.com/dmji/go-redmine/models/apierrors"
	"github.com/dmji/go-redmine/models/components"
	"github.com/dmji/go-redmine/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := goredmine.New(
		goredmine.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Username: "",
				Password: "",
			},
		}),
	)

	res, err := s.Projects.CreateProject(ctx, components.FormatXML, goredmine.String("jsmith"), &operations.CreateProjectRequestBody{
		Project: operations.Project{
			Name:       "<value>",
			Identifier: "<value>",
			CustomFields: []components.CustomFields{
				components.CustomFields{
					ID: 0,
					AdditionalProperties: map[string]any{
						"value": "string",
					},
				},
			},
			CustomFieldValues: map[string]any{
				"0": "string",
			},
		},
	})
	if err != nil {

		var e *apierrors.Errors
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```

<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->

## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:

```go
package main

import (
	"context"
	goredmine "github.com/dmji/go-redmine"
	"github.com/dmji/go-redmine/models/components"
	"github.com/dmji/go-redmine/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := goredmine.New(
		goredmine.WithServerURL("/"),
		goredmine.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Username: "",
				Password: "",
			},
		}),
	)

	res, err := s.Issues.GetIssues(ctx, operations.GetIssuesRequest{
		Format:             components.FormatXML,
		XRedmineSwitchUser: goredmine.String("jsmith"),
		Sort:               goredmine.String("id:desc"),
		CfX: map[string]string{
			"cf_0": "string",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->

## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.

<!-- End Custom HTTP Client [http-client] -->

<!-- Start Special Types [types] -->

## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```

<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation.
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release.

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/dmji/go-redmine&utm_campaign=go)
