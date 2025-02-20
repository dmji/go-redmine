// Code generated by Speakeasy (https://speakeasy.com).

package components

type Query struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	IsPublic  bool   `json:"is_public"`
	ProjectID *int64 `json:"project_id"`
}

func (o *Query) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Query) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Query) GetIsPublic() bool {
	if o == nil {
		return false
	}
	return o.IsPublic
}

func (o *Query) GetProjectID() *int64 {
	if o == nil {
		return nil
	}
	return o.ProjectID
}
