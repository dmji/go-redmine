# Groups
(*Groups*)

## Overview

Status: Alpha, Availablity: 2.1
<https://www.redmine.org/projects/redmine/wiki/Rest_Groups>

### Available Operations

* [GetGroups](#getgroups) - List groups
* [CreateGroup](#creategroup) - Create group
* [GetGroup](#getgroup) - Show group
* [UpdateGroup](#updategroup) - Update group
* [DeleteGroup](#deletegroup) - Delete group
* [AddUserToGroup](#addusertogroup) - Add user to group
* [RemoveUserFromGroup](#removeuserfromgroup) - Remove user from group

## GetGroups

List groups

<https://www.redmine.org/projects/redmine/wiki/Rest_Groups#GET>

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

    res, err := s.Groups.GetGroups(ctx, components.FormatJSON, goredmine.String("jsmith"))
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
| `xRedmineSwitchUser`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | jsmith                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetGroupsResponse](../../models/operations/getgroupsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateGroup

Create group

<https://www.redmine.org/projects/redmine/wiki/Rest_Groups#POST>

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

    res, err := s.Groups.CreateGroup(ctx, components.FormatJSON, goredmine.String("jsmith"), &operations.CreateGroupRequestBody{
        Group: operations.Group{
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
| `requestBody`                                                                           | [*operations.CreateGroupRequestBody](../../models/operations/creategrouprequestbody.md) | :heavy_minus_sign:                                                                      | N/A                                                                                     |                                                                                         |
| `opts`                                                                                  | [][operations.Option](../../models/operations/option.md)                                | :heavy_minus_sign:                                                                      | The options for this request.                                                           |                                                                                         |

### Response

**[*operations.CreateGroupResponse](../../models/operations/creategroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Errors   | 422                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetGroup

Show group

<https://www.redmine.org/projects/redmine/wiki/Rest_Groups#GET-2>

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

    res, err := s.Groups.GetGroup(ctx, components.FormatXML, 797868, goredmine.String("jsmith"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    | Example                                                                                        |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |                                                                                                |
| `format`                                                                                       | [components.Format](../../models/components/format.md)                                         | :heavy_check_mark:                                                                             | N/A                                                                                            |                                                                                                |
| `groupID`                                                                                      | *int64*                                                                                        | :heavy_check_mark:                                                                             | N/A                                                                                            |                                                                                                |
| `xRedmineSwitchUser`                                                                           | **string*                                                                                      | :heavy_minus_sign:                                                                             | N/A                                                                                            | jsmith                                                                                         |
| `include`                                                                                      | [][operations.GetGroupQueryParamInclude](../../models/operations/getgroupqueryparaminclude.md) | :heavy_minus_sign:                                                                             | N/A                                                                                            |                                                                                                |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |                                                                                                |

### Response

**[*operations.GetGroupResponse](../../models/operations/getgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateGroup

Update group

<https://www.redmine.org/projects/redmine/wiki/Rest_Groups#PUT>

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

    res, err := s.Groups.UpdateGroup(ctx, components.FormatJSON, 750192, goredmine.String("jsmith"), &operations.UpdateGroupRequestBody{
        Group: &operations.UpdateGroupGroup{
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
| `groupID`                                                                               | *int64*                                                                                 | :heavy_check_mark:                                                                      | N/A                                                                                     |                                                                                         |
| `xRedmineSwitchUser`                                                                    | **string*                                                                               | :heavy_minus_sign:                                                                      | N/A                                                                                     | jsmith                                                                                  |
| `requestBody`                                                                           | [*operations.UpdateGroupRequestBody](../../models/operations/updategrouprequestbody.md) | :heavy_minus_sign:                                                                      | N/A                                                                                     |                                                                                         |
| `opts`                                                                                  | [][operations.Option](../../models/operations/option.md)                                | :heavy_minus_sign:                                                                      | The options for this request.                                                           |                                                                                         |

### Response

**[*operations.UpdateGroupResponse](../../models/operations/updategroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteGroup

Delete group

<https://www.redmine.org/projects/redmine/wiki/Rest_Groups#DELETE>

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

    res, err := s.Groups.DeleteGroup(ctx, components.FormatJSON, 114822, goredmine.String("jsmith"))
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
| `groupID`                                                | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xRedmineSwitchUser`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | jsmith                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteGroupResponse](../../models/operations/deletegroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddUserToGroup

Add user to group

<https://www.redmine.org/projects/redmine/wiki/Rest_Groups#POST-2>

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

    res, err := s.Groups.AddUserToGroup(ctx, components.FormatJSON, 699024, goredmine.String("jsmith"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   | Example                                                                                       |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |                                                                                               |
| `format`                                                                                      | [components.Format](../../models/components/format.md)                                        | :heavy_check_mark:                                                                            | N/A                                                                                           |                                                                                               |
| `groupID`                                                                                     | *int64*                                                                                       | :heavy_check_mark:                                                                            | N/A                                                                                           |                                                                                               |
| `xRedmineSwitchUser`                                                                          | **string*                                                                                     | :heavy_minus_sign:                                                                            | N/A                                                                                           | jsmith                                                                                        |
| `requestBody`                                                                                 | [*operations.AddUserToGroupRequestBody](../../models/operations/addusertogrouprequestbody.md) | :heavy_minus_sign:                                                                            | N/A                                                                                           |                                                                                               |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |                                                                                               |

### Response

**[*operations.AddUserToGroupResponse](../../models/operations/addusertogroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RemoveUserFromGroup

Remove user from group

<https://www.redmine.org/projects/redmine/wiki/Rest_Groups#DELETE-2>

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

    res, err := s.Groups.RemoveUserFromGroup(ctx, components.FormatXML, 982330, 948975, goredmine.String("jsmith"))
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
| `groupID`                                                | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `userID`                                                 | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xRedmineSwitchUser`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | jsmith                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RemoveUserFromGroupResponse](../../models/operations/removeuserfromgroupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |