# IssueCategories
(*IssueCategories*)

## Overview

### Available Operations

* [List](#list) - List issue categories
* [Create](#create) - Create issue category
* [Show](#show) - Show issue category
* [Update](#update) - Update issue category
* [Delete](#delete) - Delete issue category

## List

List issue categories

<https://www.redmine.org/projects/redmine/wiki/Rest_IssueCategories#GET>

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

    res, err := s.IssueCategories.List(ctx, operations.GetIssueCategoriesRequest{
        Format: components.FormatJSON,
        ProjectID: 823814,
        XRedmineSwitchUser: goredmine.String("jsmith"),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetIssueCategoriesRequest](../../models/operations/getissuecategoriesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetIssueCategoriesResponse](../../models/operations/getissuecategoriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create issue category

<https://www.redmine.org/projects/redmine/wiki/Rest_IssueCategories#POST>

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

    res, err := s.IssueCategories.Create(ctx, components.FormatXML, 714556, goredmine.String("jsmith"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |                                                                                                         |
| `format`                                                                                                | [components.Format](../../models/components/format.md)                                                  | :heavy_check_mark:                                                                                      | N/A                                                                                                     |                                                                                                         |
| `projectID`                                                                                             | *int64*                                                                                                 | :heavy_check_mark:                                                                                      | N/A                                                                                                     |                                                                                                         |
| `xRedmineSwitchUser`                                                                                    | **string*                                                                                               | :heavy_minus_sign:                                                                                      | N/A                                                                                                     | jsmith                                                                                                  |
| `requestBody`                                                                                           | [*operations.CreateIssueCategoryRequestBody](../../models/operations/createissuecategoryrequestbody.md) | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |                                                                                                         |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |                                                                                                         |

### Response

**[*operations.CreateIssueCategoryResponse](../../models/operations/createissuecategoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Errors   | 422                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Show

Show issue category

<https://www.redmine.org/projects/redmine/wiki/Rest_IssueCategories#GET-2>

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

    res, err := s.IssueCategories.Show(ctx, components.FormatJSON, 735932, goredmine.String("jsmith"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `format`                                                 | [components.Format](../../models/components/format.md)   | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `issueCategoryID`                                        | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xRedmineSwitchUser`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | jsmith                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetIssueCategoryResponse](../../models/operations/getissuecategoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update issue category

<https://www.redmine.org/projects/redmine/wiki/Rest_IssueCategories#PUT>

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

    res, err := s.IssueCategories.Update(ctx, components.FormatJSON, 774678, goredmine.String("jsmith"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |                                                                                                         |
| `format`                                                                                                | [components.Format](../../models/components/format.md)                                                  | :heavy_check_mark:                                                                                      | N/A                                                                                                     |                                                                                                         |
| `issueCategoryID`                                                                                       | *int64*                                                                                                 | :heavy_check_mark:                                                                                      | N/A                                                                                                     |                                                                                                         |
| `xRedmineSwitchUser`                                                                                    | **string*                                                                                               | :heavy_minus_sign:                                                                                      | N/A                                                                                                     | jsmith                                                                                                  |
| `requestBody`                                                                                           | [*operations.UpdateIssueCategoryRequestBody](../../models/operations/updateissuecategoryrequestbody.md) | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |                                                                                                         |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |                                                                                                         |

### Response

**[*operations.UpdateIssueCategoryResponse](../../models/operations/updateissuecategoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Errors   | 422                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete issue category

<https://www.redmine.org/projects/redmine/wiki/Rest_IssueCategories#DELETE>

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

    res, err := s.IssueCategories.Delete(ctx, components.FormatJSON, 147277, goredmine.String("jsmith"), nil)
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
| `issueCategoryID`                                        | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xRedmineSwitchUser`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | jsmith                                                   |
| `reassignToID`                                           | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteIssueCategoryResponse](../../models/operations/deleteissuecategoryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |