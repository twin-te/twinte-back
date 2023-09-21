package apperr

import "fmt"

var ErrUnauthenticated error = fmt.Errorf("ErrUnauthenticated")
var ErrNotFound error = fmt.Errorf("ErrNotFound")
