// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type IssuePriority struct {
	ID           int64              `json:"id"`
	Name         string             `json:"name"`
	IsDefault    bool               `json:"is_default"`
	Active       bool               `json:"active"`
	CustomFields []CustomFieldValue `json:"custom_fields,omitempty"`
}

func (o *IssuePriority) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *IssuePriority) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *IssuePriority) GetIsDefault() bool {
	if o == nil {
		return false
	}
	return o.IsDefault
}

func (o *IssuePriority) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *IssuePriority) GetCustomFields() []CustomFieldValue {
	if o == nil {
		return nil
	}
	return o.CustomFields
}
