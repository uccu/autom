package body

import "strings"

type Body struct {
	Ref        string     `json:"ref"`
	After      string     `json:"after"`
	Repository repository `json:"repository"`
}

type repository struct {
	Name string `json:"name"`
}

func (b *Body) GetName() string {
	return b.Repository.Name
}

func (b *Body) IsTagPsuh() bool {
	return strings.HasPrefix(b.Ref, "refs/tags/")
}

func (b *Body) IsPush() bool {
	return strings.HasPrefix(b.Ref, "refs/heads/")
}

func (b *Body) IsInvalid() bool {
	return b.After != "0000000000000000000000000000000000000000"
}

func (b *Body) GetBranch() string {
	if b.IsPush() {
		return strings.Replace(b.Ref, "refs/heads/", "", 1)
	}

	if b.IsTagPsuh() {
		return strings.Replace(b.Ref, "refs/tags/", "", 1)
	}
	return ""
}
