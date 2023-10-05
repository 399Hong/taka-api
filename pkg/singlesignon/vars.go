package singlesignon

import "errors"

var ErrCannotUsePasswordAuth = errors.New("cannot use single sign on to validate password")
