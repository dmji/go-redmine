// Code generated by Speakeasy (https://speakeasy.com).

package components

import (
	"encoding/json"
	"fmt"
)

type XRedmineNometa int64

const (
	XRedmineNometaOne XRedmineNometa = 1
)

func (e XRedmineNometa) ToPointer() *XRedmineNometa {
	return &e
}

func (e *XRedmineNometa) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		*e = XRedmineNometa(v)
		return nil
	default:
		return fmt.Errorf("invalid value for XRedmineNometa: %v", v)
	}
}
