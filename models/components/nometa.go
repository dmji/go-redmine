// Code generated by Speakeasy (https://speakeasy.com).

package components

import (
	"encoding/json"
	"fmt"
)

type Nometa int64

const (
	NometaOne Nometa = 1
)

func (e Nometa) ToPointer() *Nometa {
	return &e
}

func (e *Nometa) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		*e = Nometa(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Nometa: %v", v)
	}
}
