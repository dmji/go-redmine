# Journals
(*Journals*)

## Overview

Status: Alpha, Availablity: 5.0
<https://www.redmine.org/projects/redmine/wiki/Rest_Journals?parent=Rest_api>

### Available Operations

* [UpdateJournal](#updatejournal) - Update journal

## UpdateJournal

Update journal

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

    res, err := s.Journals.UpdateJournal(ctx, components.FormatJSON, 169354, goredmine.String("jsmith"), nil)
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
| `journalID`                                                                                 | *int64*                                                                                     | :heavy_check_mark:                                                                          | N/A                                                                                         |                                                                                             |
| `xRedmineSwitchUser`                                                                        | **string*                                                                                   | :heavy_minus_sign:                                                                          | N/A                                                                                         | jsmith                                                                                      |
| `requestBody`                                                                               | [*operations.UpdateJournalRequestBody](../../models/operations/updatejournalrequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |                                                                                             |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |                                                                                             |

### Response

**[*operations.UpdateJournalResponse](../../models/operations/updatejournalresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |