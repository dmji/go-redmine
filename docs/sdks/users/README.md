# Users
(*Users*)

## Overview

Status: Stable, Availablity: 1.1
<https://www.redmine.org/projects/redmine/wiki/Rest_Users>

### Available Operations

* [GetUsers](#getusers) - List users
* [CreateUser](#createuser) - Create user
* [GetUser](#getuser) - Show user
* [UpdateUser](#updateuser) - Update user
* [DeleteUser](#deleteuser) - Delete user
* [GetCurrentUser](#getcurrentuser) - Show current user

## GetUsers

List users

<https://www.redmine.org/projects/redmine/wiki/Rest_Users#GET>

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

    res, err := s.Users.GetUsers(ctx, operations.GetUsersRequest{
        Format: components.FormatJSON,
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetUsersRequest](../../models/operations/getusersrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.GetUsersResponse](../../models/operations/getusersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateUser

Create user

<https://www.redmine.org/projects/redmine/wiki/Rest_Users#POST>

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

    res, err := s.Users.CreateUser(ctx, components.FormatXML, goredmine.String("jsmith"), &operations.CreateUserRequestBody{
        User: operations.User{
            Login: "Jimmy12",
            Firstname: "Domenico",
            Lastname: "Cormier",
            Mail: "<value>",
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

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |                                                                                       |
| `format`                                                                              | [components.Format](../../models/components/format.md)                                | :heavy_check_mark:                                                                    | N/A                                                                                   |                                                                                       |
| `xRedmineSwitchUser`                                                                  | **string*                                                                             | :heavy_minus_sign:                                                                    | N/A                                                                                   | jsmith                                                                                |
| `requestBody`                                                                         | [*operations.CreateUserRequestBody](../../models/operations/createuserrequestbody.md) | :heavy_minus_sign:                                                                    | N/A                                                                                   |                                                                                       |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |                                                                                       |

### Response

**[*operations.CreateUserResponse](../../models/operations/createuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Errors   | 422                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetUser

Show user

<https://www.redmine.org/projects/redmine/wiki/Rest_Users#GET-2>

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

    res, err := s.Users.GetUser(ctx, components.FormatXML, 161855, goredmine.String("jsmith"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `format`                                                                                     | [components.Format](../../models/components/format.md)                                       | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `userID`                                                                                     | *int64*                                                                                      | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `xRedmineSwitchUser`                                                                         | **string*                                                                                    | :heavy_minus_sign:                                                                           | N/A                                                                                          | jsmith                                                                                       |
| `include`                                                                                    | [][operations.GetUserQueryParamInclude](../../models/operations/getuserqueryparaminclude.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.GetUserResponse](../../models/operations/getuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateUser

Update user

<https://www.redmine.org/projects/redmine/wiki/Rest_Users#PUT>

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

    res, err := s.Users.UpdateUser(ctx, components.FormatJSON, 713914, goredmine.String("jsmith"), &operations.UpdateUserRequestBody{
        User: &operations.UpdateUserUser{
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

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |                                                                                       |
| `format`                                                                              | [components.Format](../../models/components/format.md)                                | :heavy_check_mark:                                                                    | N/A                                                                                   |                                                                                       |
| `userID`                                                                              | *int64*                                                                               | :heavy_check_mark:                                                                    | N/A                                                                                   |                                                                                       |
| `xRedmineSwitchUser`                                                                  | **string*                                                                             | :heavy_minus_sign:                                                                    | N/A                                                                                   | jsmith                                                                                |
| `requestBody`                                                                         | [*operations.UpdateUserRequestBody](../../models/operations/updateuserrequestbody.md) | :heavy_minus_sign:                                                                    | N/A                                                                                   |                                                                                       |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |                                                                                       |

### Response

**[*operations.UpdateUserResponse](../../models/operations/updateuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Errors   | 422                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteUser

Delete user

<https://www.redmine.org/projects/redmine/wiki/Rest_Users#DELETE>

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

    res, err := s.Users.DeleteUser(ctx, components.FormatJSON, 335223, goredmine.String("jsmith"))
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
| `userID`                                                 | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xRedmineSwitchUser`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | jsmith                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteUserResponse](../../models/operations/deleteuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetCurrentUser

Show current user

<https://www.redmine.org/projects/redmine/wiki/Rest_Users#GET-2>

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

    res, err := s.Users.GetCurrentUser(ctx, components.FormatXML, goredmine.String("jsmith"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `format`                                                                                                   | [components.Format](../../models/components/format.md)                                                     | :heavy_check_mark:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `xRedmineSwitchUser`                                                                                       | **string*                                                                                                  | :heavy_minus_sign:                                                                                         | N/A                                                                                                        | jsmith                                                                                                     |
| `include`                                                                                                  | [][operations.GetCurrentUserQueryParamInclude](../../models/operations/getcurrentuserqueryparaminclude.md) | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |                                                                                                            |

### Response

**[*operations.GetCurrentUserResponse](../../models/operations/getcurrentuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |