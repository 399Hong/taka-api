package singlesignon

type AuthType int

const (
	Password AuthType = iota + 1
	Google
)

type Claim struct {
	Id    string
	Email string
}
