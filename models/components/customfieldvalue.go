// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/dmji/go-redmine/internal/utils"
)

type CustomFieldValue struct {
	ID                   int64          `json:"id"`
	Name                 string         `json:"name"`
	Multiple             *bool          `json:"multiple,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (c CustomFieldValue) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CustomFieldValue) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CustomFieldValue) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *CustomFieldValue) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CustomFieldValue) GetMultiple() *bool {
	if o == nil {
		return nil
	}
	return o.Multiple
}

func (o *CustomFieldValue) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}
