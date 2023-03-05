package errs

import "fmt"

var (
	ErrUserExist = fmt.Errorf("user is existed")
	ErrDbError   = fmt.Errorf("db error")
)
