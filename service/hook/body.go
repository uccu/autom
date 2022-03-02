package hook

import "strings"

type HookBody struct {
	EventName   string             `json:"event_name"`
	Ref         string             `json:"ref"`
	CheckoutSha *string            `json:"checkout_sha"`
	Repository  HookBodyRepository `json:"repository"`
}

type HookBodyRepository struct {
	Name string `json:"name"`
}

func (b *HookBody) GetEvent() string {
	return b.EventName
}

func (b *HookBody) GetName() string {
	return b.Repository.Name
}

func (b *HookBody) GetBranch() string {
	if b.GetEvent() == "push" {
		return strings.Replace(b.Ref, "refs/heads/", "", 1)
	}

	if b.GetEvent() == "tag_push" {
		return strings.Replace(b.Ref, "refs/tags/", "", 1)
	}
	return ""
}
