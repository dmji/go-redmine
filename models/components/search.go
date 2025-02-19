// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/dmji/go-redmine/internal/utils"
	"time"
)

type Search struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Datetime    time.Time `json:"datetime"`
}

func (s Search) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Search) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Search) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Search) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *Search) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *Search) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *Search) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *Search) GetDatetime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Datetime
}
