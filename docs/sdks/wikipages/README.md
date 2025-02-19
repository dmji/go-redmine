# WikiPages
(*WikiPages*)

## Overview

### Available Operations

* [List](#list) - List wiki pages
* [Get](#get) - Show wiki page
* [Update](#update) - Create or update wiki page
* [Delete](#delete) - Delete wiki page
* [GetByVersion](#getbyversion) - Show wiki page by specific version

## List

List wiki pages

<https://www.redmine.org/projects/redmine/wiki/Rest_WikiPages#Getting-the-pages-list-of-a-wiki>

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

    res, err := s.WikiPages.List(ctx, components.FormatJSON, 221236, goredmine.String("jsmith"))
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
| `projectID`                                              | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xRedmineSwitchUser`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | jsmith                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetWikiPagesResponse](../../models/operations/getwikipagesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Show wiki page

<https://www.redmine.org/projects/redmine/wiki/Rest_WikiPages#Getting-a-wiki-page>

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

    res, err := s.WikiPages.Get(ctx, operations.GetWikiPageRequest{
        Format: components.FormatJSON,
        ProjectID: 584534,
        WikiPageTitle: "<value>",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetWikiPageRequest](../../models/operations/getwikipagerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetWikiPageResponse](../../models/operations/getwikipageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Create or update wiki page

<https://www.redmine.org/projects/redmine/wiki/Rest_WikiPages#Creating-or-updating-a-wiki-page>

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

    res, err := s.WikiPages.Update(ctx, operations.UpdateWikiPageRequest{
        Format: components.FormatJSON,
        ProjectID: 44462,
        WikiPageTitle: "<value>",
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
| `request`                                                                            | [operations.UpdateWikiPageRequest](../../models/operations/updatewikipagerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateWikiPageResponse](../../models/operations/updatewikipageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Errors   | 409, 422           | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete wiki page

<https://www.redmine.org/projects/redmine/wiki/Rest_WikiPages#Deleting-a-wiki-page>

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

    res, err := s.WikiPages.Delete(ctx, components.FormatJSON, 16238, "<value>", goredmine.String("jsmith"))
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
| `projectID`                                              | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `wikiPageTitle`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xRedmineSwitchUser`                                     | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | jsmith                                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteWikiPageResponse](../../models/operations/deletewikipageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetByVersion

Show wiki page by specific version

<https://www.redmine.org/projects/redmine/wiki/Rest_WikiPages#Getting-an-old-version-of-a-wiki-page>

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

    res, err := s.WikiPages.GetByVersion(ctx, operations.GetWikiPageByVersionRequest{
        Format: components.FormatXML,
        ProjectID: 876217,
        WikiPageTitle: "<value>",
        VersionID: 975705,
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
| `request`                                                                                        | [operations.GetWikiPageByVersionRequest](../../models/operations/getwikipagebyversionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetWikiPageByVersionResponse](../../models/operations/getwikipagebyversionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |