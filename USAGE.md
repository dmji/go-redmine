<!-- Start SDK Example Usage [usage] -->

```go
package main

import (
	"context"
	redmine "github.com/dmji/go-redmine"
	"github.com/dmji/go-redmine/models/components"
	"github.com/dmji/go-redmine/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := redmine.New(
		redmine.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Username: "",
				Password: "",
			},
		}),
	)

	res, err := s.Issues.GetIssues(ctx, operations.GetIssuesRequest{
		Format:             components.FormatXML,
		XRedmineSwitchUser: redmine.String("jsmith"),
		Sort:               redmine.String("id:desc"),
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
