package user

type Gender int64

const (
	Unknown Gender = iota
	Male
	Female
)

type Type int64

// need to revisit
const (
	Candidate Type = iota
	employer
)
