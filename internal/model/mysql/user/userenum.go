package user

type Gender int64

const (
	Unknown Gender = iota
	Male
	Female
)

type RoleType int64

// need to revisit
const (
	Candidate RoleType = iota + 1
	Employer
)
