package main

import (
	"context"
	"log"

	redmine "github.com/dmji/go-redmine"
	"github.com/dmji/go-redmine/models/components"
	"github.com/dmji/go-redmine/models/operations"
)

func sptr(s string) *string {
	return &s
}

func main() {
	ctx := context.Background()

	s := redmine.New(
		redmine.WithSecurity(components.Security{
			APIKeyAuth: sptr("58af54977203eb4fc28a8eae088eff0d4f855153"),
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
