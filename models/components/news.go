// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/dmji/go-redmine/internal/utils"
	"time"
)

type News struct {
	ID          int64         `json:"id"`
	Project     IDName        `json:"project"`
	Author      IDName        `json:"author"`
	Title       string        `json:"title"`
	Summary     *string       `json:"summary,omitempty"`
	Description string        `json:"description"`
	CreatedOn   time.Time     `json:"created_on"`
	Attachments []Attachment  `json:"attachments,omitempty"`
	Comments    []NewsComment `json:"comments,omitempty"`
}

func (n News) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *News) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *News) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *News) GetProject() IDName {
	if o == nil {
		return IDName{}
	}
	return o.Project
}

func (o *News) GetAuthor() IDName {
	if o == nil {
		return IDName{}
	}
	return o.Author
}

func (o *News) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *News) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *News) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *News) GetCreatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedOn
}

func (o *News) GetAttachments() []Attachment {
	if o == nil {
		return nil
	}
	return o.Attachments
}

func (o *News) GetComments() []NewsComment {
	if o == nil {
		return nil
	}
	return o.Comments
}
