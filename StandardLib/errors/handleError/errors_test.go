package handleError

import (
	"github.com/pkg/errors"
)

type temporary interface {
	Temporary() bool
}

// IsTemporary returns true if err is temporary.
func IsTemporary(err error) bool {
	te, ok := errors.Cause(err).(temporary)
	return ok && te.Temporary()
}

