package body

import (
	"strings"
)

type GitlabBody struct {
	EventName   string               `json:"event_name"`
	Ref         string               `json:"ref"`
	CheckoutSha *string              `json:"checkout_sha"`
	Repository  GitlabBodyRepository `json:"repository"`
}

type GitlabBodyRepository struct {
	Name string `json:"name"`
}

func (b *GitlabBody) GetName() string {
	return b.Repository.Name
}

func (b *GitlabBody) IsTagPsuh() bool {
	return b.EventName == "tag_push"
}

func (b *GitlabBody) IsPush() bool {
	return b.EventName == "push"
}

func (b *GitlabBody) IsInvalid() bool {
	return b.CheckoutSha != nil
}

func (b *GitlabBody) GetBranch() string {
	if b.IsPush() {
		return strings.Replace(b.Ref, "refs/heads/", "", 1)
	}

	if b.IsTagPsuh() {
		return strings.Replace(b.Ref, "refs/tags/", "", 1)
	}
	return ""
}
