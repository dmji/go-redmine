# Versions
(*Versions*)

## Overview

Status: Alpha, Availablity: 1.3
<https://www.redmine.org/projects/redmine/wiki/Rest_Versions>

### Available Operations

* [ListByProject](#listbyproject) - List versions by project
* [Create](#create) - Create version
* [Get](#get) - Show version
* [Update](#update) - Update version
* [Delete](#delete) - Delete version

## ListByProject

List versions by project

<https://www.redmine.org/projects/redmine/wiki/Rest_Versions#GET>

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

    res, err := s.Versions.ListByProject(ctx, operations.GetVersionsByProjectRequest{
        Format: components.FormatJSON,
        ProjectID: 131775,
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetVersionsByProjectRequest](../../models/operations/getversionsbyprojectrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetVersionsByProjectResponse](../../models/operations/getversionsbyprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create version

<https://www.redmine.org/projects/redmine/wiki/Rest_Versions#POST>

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

    res, err := s.Versions.Create(ctx, components.FormatJSON, 331583, goredmine.String("jsmith"), &operations.CreateVersionRequestBody{
        Version: operations.Version{
            Name: "<value>",
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

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 | Example                                                                                     |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |                                                                                             |
| `format`                                                                                    | [components.Format](../../models/components/format.md)                                      | :heavy_check_mark:                                                                          | N/A                                                                                         |                                                                                             |
| `projectID`                                                                                 | *int64*                                                                                     | :heavy_check_mark:                                                                          | N/A                                                                                         |                                                                                             |
| `xRedmineSwitchUser`                                                                        | **string*                                                                                   | :heavy_minus_sign:                                                                          | N/A                                                                                         | jsmith                                                                                      |
| `requestBody`                                                                               | [*operations.CreateVersionRequestBody](../../models/operations/createversionrequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |                                                                                             |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |                                                                                             |

### Response

**[*operations.CreateVersionResponse](../../models/operations/createversionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Errors   | 422                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Show version

<https://www.redmine.org/projects/redmine/wiki/Rest_Versions#GET-2>

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

    res, err := s.Versions.Get(ctx, components.FormatJSON, 158895, goredmine.String("jsmith"))
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
| `versionID`                                              | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xRedmineSwitchUser`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | jsmith                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetVersionsResponse](../../models/operations/getversionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update version

<https://www.redmine.org/projects/redmine/wiki/Rest_Versions#PUT>

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

    res, err := s.Versions.Update(ctx, components.FormatJSON, 55046, goredmine.String("jsmith"), &operations.UpdateVersionRequestBody{
        Version: &operations.UpdateVersionVersion{
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

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 | Example                                                                                     |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |                                                                                             |
| `format`                                                                                    | [components.Format](../../models/components/format.md)                                      | :heavy_check_mark:                                                                          | N/A                                                                                         |                                                                                             |
| `versionID`                                                                                 | *int64*                                                                                     | :heavy_check_mark:                                                                          | N/A                                                                                         |                                                                                             |
| `xRedmineSwitchUser`                                                                        | **string*                                                                                   | :heavy_minus_sign:                                                                          | N/A                                                                                         | jsmith                                                                                      |
| `requestBody`                                                                               | [*operations.UpdateVersionRequestBody](../../models/operations/updateversionrequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |                                                                                             |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |                                                                                             |

### Response

**[*operations.UpdateVersionResponse](../../models/operations/updateversionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Errors   | 422                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete version

<https://www.redmine.org/projects/redmine/wiki/Rest_Versions#DELETE>

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

    res, err := s.Versions.Delete(ctx, components.FormatXML, 964250, goredmine.String("jsmith"))
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
| `versionID`                                              | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xRedmineSwitchUser`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | jsmith                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteVersionResponse](../../models/operations/deleteversionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Errors   | 422                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |