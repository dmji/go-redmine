package main

import (
	"context"
	"log"

	goredmine "github.com/dmji/go-redmine"
	"github.com/dmji/go-redmine/models/components"
	"github.com/dmji/go-redmine/models/operations"
)

func sptr(s string) *string {
	return &s
}

func main() {
	ctx := context.Background()

	s := goredmine.New(
		goredmine.WithSecurity(components.Security{
			APIKeyAuth: sptr("58af54977203eb4fc28a8eae088eff0d4f855153"),
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
