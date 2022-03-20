package body

type Body interface {
	GetName() string
	IsTagPsuh() bool
	IsPush() bool
	IsInvalid() bool
	GetBranch() string
}
