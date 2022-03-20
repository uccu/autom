package body

import (
	"strings"
)

type GithubBody struct {
	Ref        string               `json:"ref"`
	After      string               `json:"after"`
	Repository GithubBodyRepository `json:"repository"`
}

type GithubBodyRepository struct {
	Name string `json:"name"`
}

func (b *GithubBody) GetName() string {
	return b.Repository.Name
}

func (b *GithubBody) IsTagPsuh() bool {
	return strings.HasPrefix(b.Ref, "refs/tags/")
}

func (b *GithubBody) IsPush() bool {
	return strings.HasPrefix(b.Ref, "refs/heads/")
}

func (b *GithubBody) IsInvalid() bool {
	return b.After != "0000000000000000000000000000000000000000"
}

func (b *GithubBody) GetBranch() string {
	if b.IsPush() {
		return strings.Replace(b.Ref, "refs/heads/", "", 1)
	}

	if b.IsTagPsuh() {
		return strings.Replace(b.Ref, "refs/tags/", "", 1)
	}
	return ""
}
