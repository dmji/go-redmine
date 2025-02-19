# Issues
(*Issues*)

## Overview

Status: Stable, Availablity: 1.0
<https://www.redmine.org/projects/redmine/wiki/Rest_Issues>

### Available Operations

* [List](#list) - List issues
* [Create](#create) - Create issue
* [Get](#get) - Show issue
* [Update](#update) - Update issue
* [Delete](#delete) - Delete issue
* [AddWatcher](#addwatcher) - Add watcher
* [RemoveWatcher](#removewatcher) - Remove watcher

## List

List issues

<https://www.redmine.org/projects/redmine/wiki/Rest_Issues#Listing-issues>

### Example Usage

```go
package main

import(
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
        Format: components.FormatXML,
        XRedmineSwitchUser: goredmine.String("jsmith"),
        Sort: goredmine.String("id:desc"),
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

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetIssuesRequest](../../models/operations/getissuesrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.GetIssuesResponse](../../models/operations/getissuesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create issue

<https://www.redmine.org/projects/redmine/wiki/Rest_Issues#Creating-an-issue>

### Example Usage

```go
package main

import(
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

    res, err := s.Issues.Create(ctx, components.FormatJSON, goredmine.String("jsmith"), &operations.CreateIssueRequestBody{
        Issue: operations.Issue{
            ProjectID: 46402,
            Subject: "<value>",
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
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             | Example                                                                                 |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |                                                                                         |
| `format`                                                                                | [components.Format](../../models/components/format.md)                                  | :heavy_check_mark:                                                                      | N/A                                                                                     |                                                                                         |
| `xRedmineSwitchUser`                                                                    | **string*                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     | jsmith                                                                                  |
| `requestBody`                                                                           | [*operations.CreateIssueRequestBody](../../models/operations/createissuerequestbody.md) | :heavy_minus_sign:                                                                      | N/A                                                                                     |                                                                                         |
| `opts`                                                                                  | [][operations.Option](../../models/operations/option.md)                                | :heavy_minus_sign:                                                                      | The options for this request.                                                           |                                                                                         |

### Response

**[*operations.CreateIssueResponse](../../models/operations/createissueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Show issue

<https://www.redmine.org/projects/redmine/wiki/Rest_Issues#Showing-an-issue>

### Example Usage

```go
package main

import(
	"context"
	goredmine "github.com/dmji/go-redmine"
	"github.com/dmji/go-redmine/models/components"
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

    res, err := s.Issues.Get(ctx, components.FormatXML, 310882, goredmine.String("jsmith"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `format`                                                                       | [components.Format](../../models/components/format.md)                         | :heavy_check_mark:                                                             | N/A                                                                            |                                                                                |
| `issueID`                                                                      | *int64*                                                                        | :heavy_check_mark:                                                             | N/A                                                                            |                                                                                |
| `xRedmineSwitchUser`                                                           | **string*                                                                      | :heavy_minus_sign:                                                             | N/A                                                                            | jsmith                                                                         |
| `include`                                                                      | [][operations.QueryParamInclude](../../models/operations/queryparaminclude.md) | :heavy_minus_sign:                                                             | N/A                                                                            |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.GetIssueResponse](../../models/operations/getissueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update issue

<https://www.redmine.org/projects/redmine/wiki/Rest_Issues#Updating-an-issue>

### Example Usage

```go
package main

import(
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

    res, err := s.Issues.Update(ctx, components.FormatXML, 733216, goredmine.String("jsmith"), &operations.UpdateIssueRequestBody{
        Issue: &operations.UpdateIssueIssue{
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
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             | Example                                                                                 |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |                                                                                         |
| `format`                                                                                | [components.Format](../../models/components/format.md)                                  | :heavy_check_mark:                                                                      | N/A                                                                                     |                                                                                         |
| `issueID`                                                                               | *int64*                                                                                 | :heavy_check_mark:                                                                      | N/A                                                                                     |                                                                                         |
| `xRedmineSwitchUser`                                                                    | **string*                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     | jsmith                                                                                  |
| `requestBody`                                                                           | [*operations.UpdateIssueRequestBody](../../models/operations/updateissuerequestbody.md) | :heavy_minus_sign:                                                                      | N/A                                                                                     |                                                                                         |
| `opts`                                                                                  | [][operations.Option](../../models/operations/option.md)                                | :heavy_minus_sign:                                                                      | The options for this request.                                                           |                                                                                         |

### Response

**[*operations.UpdateIssueResponse](../../models/operations/updateissueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete issue

<https://www.redmine.org/projects/redmine/wiki/Rest_Issues#Deleting-an-issue>

### Example Usage

```go
package main

import(
	"context"
	goredmine "github.com/dmji/go-redmine"
	"github.com/dmji/go-redmine/models/components"
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

    res, err := s.Issues.Delete(ctx, components.FormatXML, 338722, goredmine.String("jsmith"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `format`                                                 | [components.Format](../../models/components/format.md)   | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `issueID`                                                | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xRedmineSwitchUser`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | jsmith                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteIssueResponse](../../models/operations/deleteissueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddWatcher

Add watcher

<https://www.redmine.org/projects/redmine/wiki/Rest_Issues#Adding-a-watcher>

### Example Usage

```go
package main

import(
	"context"
	goredmine "github.com/dmji/go-redmine"
	"github.com/dmji/go-redmine/models/components"
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

    res, err := s.Issues.AddWatcher(ctx, components.FormatJSON, 831564, goredmine.String("jsmith"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |                                                                                       |
| `format`                                                                              | [components.Format](../../models/components/format.md)                                | :heavy_check_mark:                                                                    | N/A                                                                                   |                                                                                       |
| `issueID`                                                                             | *int64*                                                                               | :heavy_check_mark:                                                                    | N/A                                                                                   |                                                                                       |
| `xRedmineSwitchUser`                                                                  | **string*                                                                             | :heavy_minus_sign:                                                                    | N/A                                                                                   | jsmith                                                                                |
| `requestBody`                                                                         | [*operations.AddWatcherRequestBody](../../models/operations/addwatcherrequestbody.md) | :heavy_minus_sign:                                                                    | N/A                                                                                   |                                                                                       |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |                                                                                       |

### Response

**[*operations.AddWatcherResponse](../../models/operations/addwatcherresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RemoveWatcher

Remove watcher

<https://www.redmine.org/projects/redmine/wiki/Rest_Issues#Removing-a-watcher>

### Example Usage

```go
package main

import(
	"context"
	goredmine "github.com/dmji/go-redmine"
	"github.com/dmji/go-redmine/models/components"
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

    res, err := s.Issues.RemoveWatcher(ctx, components.FormatXML, 542647, 434528, goredmine.String("jsmith"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `format`                                                 | [components.Format](../../models/components/format.md)   | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `issueID`                                                | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `userID`                                                 | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xRedmineSwitchUser`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | jsmith                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RemoveWatcherResponse](../../models/operations/removewatcherresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |