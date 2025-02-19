// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type Format string

const (
	FormatJSON Format = "json"
	FormatXML  Format = "xml"
)

func (e Format) ToPointer() *Format {
	return &e
}
func (e *Format) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "xml":
		*e = Format(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Format: %v", v)
	}
}
