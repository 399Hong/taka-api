package signon

type AuthType int

const (
	Native AuthType = iota + 1
	Google
)

type Claim struct {
	Id    string
	Email string
}
