package errs

import "fmt"

var (
	ErrUserExist     = fmt.Errorf("user is existed")
	ErrDbError       = fmt.Errorf("db error")
	ErrOutOfCapacity = fmt.Errorf("out of capacity")
	ErrBucketOwner   = fmt.Errorf("you are not the bucket owner")
)
