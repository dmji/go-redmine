# ProjectMemberships
(*ProjectMemberships*)

## Overview

### Available Operations

* [List](#list) - List memberships
* [Create](#create) - Create membership
* [GetMembership](#getmembership) - Show membership
* [UpdateMembership](#updatemembership) - Update membership
* [Delete](#delete) - Delete membership

## List

List memberships

<https://www.redmine.org/projects/redmine/wiki/Rest_Memberships#GET>

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

    res, err := s.ProjectMemberships.List(ctx, operations.GetMembershipsRequest{
        Format: components.FormatXML,
        ProjectID: 35144,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetMembershipsRequest](../../models/operations/getmembershipsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetMembershipsResponse](../../models/operations/getmembershipsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create membership

<https://www.redmine.org/projects/redmine/wiki/Rest_Memberships#POST>

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

    res, err := s.ProjectMemberships.Create(ctx, components.FormatJSON, 825309, goredmine.String("jsmith"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       | Example                                                                                           |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |                                                                                                   |
| `format`                                                                                          | [components.Format](../../models/components/format.md)                                            | :heavy_check_mark:                                                                                | N/A                                                                                               |                                                                                                   |
| `projectID`                                                                                       | *int64*                                                                                           | :heavy_check_mark:                                                                                | N/A                                                                                               |                                                                                                   |
| `xRedmineSwitchUser`                                                                              | **string*                                                                                         | :heavy_minus_sign:                                                                                | N/A                                                                                               | jsmith                                                                                            |
| `requestBody`                                                                                     | [*operations.CreateMembershipRequestBody](../../models/operations/createmembershiprequestbody.md) | :heavy_minus_sign:                                                                                | N/A                                                                                               |                                                                                                   |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |                                                                                                   |

### Response

**[*operations.CreateMembershipResponse](../../models/operations/createmembershipresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetMembership

Show membership

<https://www.redmine.org/projects/redmine/wiki/Rest_Memberships#GET-2>

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

    res, err := s.ProjectMemberships.GetMembership(ctx, components.FormatXML, 551384, goredmine.String("jsmith"))
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
| `membershipID`                                           | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xRedmineSwitchUser`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | jsmith                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetMembershipResponse](../../models/operations/getmembershipresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateMembership

Update membership

<https://www.redmine.org/projects/redmine/wiki/Rest_Memberships#PUT>

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

    res, err := s.ProjectMemberships.UpdateMembership(ctx, components.FormatJSON, 404505, goredmine.String("jsmith"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       | Example                                                                                           |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |                                                                                                   |
| `format`                                                                                          | [components.Format](../../models/components/format.md)                                            | :heavy_check_mark:                                                                                | N/A                                                                                               |                                                                                                   |
| `membershipID`                                                                                    | *int64*                                                                                           | :heavy_check_mark:                                                                                | N/A                                                                                               |                                                                                                   |
| `xRedmineSwitchUser`                                                                              | **string*                                                                                         | :heavy_minus_sign:                                                                                | N/A                                                                                               | jsmith                                                                                            |
| `requestBody`                                                                                     | [*operations.UpdateMembershipRequestBody](../../models/operations/updatemembershiprequestbody.md) | :heavy_minus_sign:                                                                                | N/A                                                                                               |                                                                                                   |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |                                                                                                   |

### Response

**[*operations.UpdateMembershipResponse](../../models/operations/updatemembershipresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Errors   | 422                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete membership

<https://www.redmine.org/projects/redmine/wiki/Rest_Memberships#DELETE>

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

    res, err := s.ProjectMemberships.Delete(ctx, components.FormatJSON, 165422, goredmine.String("jsmith"))
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
| `membershipID`                                           | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xRedmineSwitchUser`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | jsmith                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteMembershipResponse](../../models/operations/deletemembershipresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |