# github.com/dmji/go-redmine

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/dmji/go-redmine* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/dmji/go-redmine&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/dmji/jikan). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Redmine API: Unofficial OpenAPI specification for the Redmine API.

For more information about the API: [Redmine API Official Developer Guide](https://www.redmine.org/projects/redmine/wiki/Rest_api)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/dmji/go-redmine](#githubcomdmjigo-redmine)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Special Types](#special-types)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

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

	res, err := s.Issues.List(ctx, operations.GetIssuesRequest{
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

	res, err := s.Issues.List(ctx, operations.GetIssuesRequest{
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

* [Get](docs/sdks/attachments/README.md#get) - Show attachment
* [Update](docs/sdks/attachments/README.md#update) - Update attachment
* [Delete](docs/sdks/attachments/README.md#delete) - Delete attachment
* [Download](docs/sdks/attachments/README.md#download) - Download attachment file
* [DownloadThumbnail](docs/sdks/attachments/README.md#downloadthumbnail) - Download thumbnail
* [UploadFile](docs/sdks/attachments/README.md#uploadfile) - Upload attachment file

### [CustomFields](docs/sdks/customfields/README.md)

* [List](docs/sdks/customfields/README.md#list) - List custom fields

### [Enumerations](docs/sdks/enumerations/README.md)

* [GetIssuePriorities](docs/sdks/enumerations/README.md#getissuepriorities) - List issue priorities
* [GetTimeEntryActivities](docs/sdks/enumerations/README.md#gettimeentryactivities) - List time entry activities
* [GetDocumentCategories](docs/sdks/enumerations/README.md#getdocumentcategories) - List document categories

### [Files](docs/sdks/files/README.md)

* [List](docs/sdks/files/README.md#list) - List files
* [Create](docs/sdks/files/README.md#create) - Create file


### [Groups](docs/sdks/groups/README.md)

* [List](docs/sdks/groups/README.md#list) - List groups
* [Create](docs/sdks/groups/README.md#create) - Create group
* [Get](docs/sdks/groups/README.md#get) - Show group
* [Update](docs/sdks/groups/README.md#update) - Update group
* [Delete](docs/sdks/groups/README.md#delete) - Delete group
* [AddUser](docs/sdks/groups/README.md#adduser) - Add user to group
* [RemoveUser](docs/sdks/groups/README.md#removeuser) - Remove user from group

### [IssueCategories](docs/sdks/issuecategories/README.md)

* [List](docs/sdks/issuecategories/README.md#list) - List issue categories
* [Create](docs/sdks/issuecategories/README.md#create) - Create issue category
* [Show](docs/sdks/issuecategories/README.md#show) - Show issue category
* [Update](docs/sdks/issuecategories/README.md#update) - Update issue category
* [Delete](docs/sdks/issuecategories/README.md#delete) - Delete issue category

### [IssueRelations](docs/sdks/issuerelations/README.md)

* [List](docs/sdks/issuerelations/README.md#list) - List issue relations
* [Create](docs/sdks/issuerelations/README.md#create) - Create issue relation
* [Get](docs/sdks/issuerelations/README.md#get) - Show issue relation
* [Delete](docs/sdks/issuerelations/README.md#delete) - Delete issue relation

### [Issues](docs/sdks/issues/README.md)

* [List](docs/sdks/issues/README.md#list) - List issues
* [Create](docs/sdks/issues/README.md#create) - Create issue
* [Get](docs/sdks/issues/README.md#get) - Show issue
* [Update](docs/sdks/issues/README.md#update) - Update issue
* [Delete](docs/sdks/issues/README.md#delete) - Delete issue
* [AddWatcher](docs/sdks/issues/README.md#addwatcher) - Add watcher
* [RemoveWatcher](docs/sdks/issues/README.md#removewatcher) - Remove watcher

### [IssueStatuses](docs/sdks/issuestatuses/README.md)

* [List](docs/sdks/issuestatuses/README.md#list) - List issue statuses

### [Journals](docs/sdks/journals/README.md)

* [Update](docs/sdks/journals/README.md#update) - Update journal

### [MyAccount](docs/sdks/myaccount/README.md)

* [Get](docs/sdks/myaccount/README.md#get) - Show my account
* [Update](docs/sdks/myaccount/README.md#update) - Update my account

### [News](docs/sdks/news/README.md)

* [List](docs/sdks/news/README.md#list) - List news
* [Get](docs/sdks/news/README.md#get) - Show news
* [Update](docs/sdks/news/README.md#update) - Update news
* [Delete](docs/sdks/news/README.md#delete) - Delete news
* [ListByProject](docs/sdks/news/README.md#listbyproject) - List news by project
* [Create](docs/sdks/news/README.md#create) - Create news

### [ProjectMemberships](docs/sdks/projectmemberships/README.md)

* [List](docs/sdks/projectmemberships/README.md#list) - List memberships
* [Create](docs/sdks/projectmemberships/README.md#create) - Create membership
* [GetMembership](docs/sdks/projectmemberships/README.md#getmembership) - Show membership
* [UpdateMembership](docs/sdks/projectmemberships/README.md#updatemembership) - Update membership
* [Delete](docs/sdks/projectmemberships/README.md#delete) - Delete membership

### [Projects](docs/sdks/projects/README.md)

* [List](docs/sdks/projects/README.md#list) - List projects
* [Create](docs/sdks/projects/README.md#create) - Crete project
* [Get](docs/sdks/projects/README.md#get) - Show project
* [Update](docs/sdks/projects/README.md#update) - Update project
* [Delete](docs/sdks/projects/README.md#delete) - Delete project
* [Archive](docs/sdks/projects/README.md#archive) - Archive project
* [Unarchive](docs/sdks/projects/README.md#unarchive) - Unarchive project
* [Close](docs/sdks/projects/README.md#close) - Close project
* [Reopen](docs/sdks/projects/README.md#reopen) - Reopen project

### [Queries](docs/sdks/queries/README.md)

* [List](docs/sdks/queries/README.md#list) - List queries

### [Repositories](docs/sdks/repositories/README.md)

* [AddRelatedIssue](docs/sdks/repositories/README.md#addrelatedissue) - Add related issue
* [RemoveRelatedIssue](docs/sdks/repositories/README.md#removerelatedissue) - Remove related issue

### [Roles](docs/sdks/roles/README.md)

* [List](docs/sdks/roles/README.md#list) - List roles
* [Get](docs/sdks/roles/README.md#get) - Show role

### [Search](docs/sdks/search/README.md)

* [Execute](docs/sdks/search/README.md#execute) - Search

### [TimeEntries](docs/sdks/timeentries/README.md)

* [List](docs/sdks/timeentries/README.md#list) - List time entries
* [Create](docs/sdks/timeentries/README.md#create) - Create time entry
* [Get](docs/sdks/timeentries/README.md#get) - Show time entry
* [Update](docs/sdks/timeentries/README.md#update) - Update time entry
* [Delete](docs/sdks/timeentries/README.md#delete) - Delete time entry

### [Trackers](docs/sdks/trackers/README.md)

* [List](docs/sdks/trackers/README.md#list) - List trackers

### [Users](docs/sdks/users/README.md)

* [List](docs/sdks/users/README.md#list) - List users
* [Create](docs/sdks/users/README.md#create) - Create user
* [Show](docs/sdks/users/README.md#show) - Show user
* [Update](docs/sdks/users/README.md#update) - Update user
* [Delete](docs/sdks/users/README.md#delete) - Delete user
* [GetCurrent](docs/sdks/users/README.md#getcurrent) - Show current user

### [Versions](docs/sdks/versions/README.md)

* [ListByProject](docs/sdks/versions/README.md#listbyproject) - List versions by project
* [Create](docs/sdks/versions/README.md#create) - Create version
* [Get](docs/sdks/versions/README.md#get) - Show version
* [Update](docs/sdks/versions/README.md#update) - Update version
* [Delete](docs/sdks/versions/README.md#delete) - Delete version

### [WikiPages](docs/sdks/wikipages/README.md)

* [List](docs/sdks/wikipages/README.md#list) - List wiki pages
* [Get](docs/sdks/wikipages/README.md#get) - Show wiki page
* [Update](docs/sdks/wikipages/README.md#update) - Create or update wiki page
* [Delete](docs/sdks/wikipages/README.md#delete) - Delete wiki page
* [GetByVersion](docs/sdks/wikipages/README.md#getbyversion) - Show wiki page by specific version

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

	res, err := s.Issues.List(ctx, operations.GetIssuesRequest{
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

	res, err := s.Issues.List(ctx, operations.GetIssuesRequest{
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

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Create` function may return the following errors:

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

	res, err := s.Projects.Create(ctx, components.FormatXML, goredmine.String("jsmith"), &operations.CreateProjectRequestBody{
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

	res, err := s.Issues.List(ctx, operations.GetIssuesRequest{
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
