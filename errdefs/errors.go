package errdefs

import "errors"

var ErrMaxWarningsReached = errors.New("max_warnings_reached")
var ErrRoleAddFailed = errors.New("role_add_failed")

var ErrNoWarnsToRemove = errors.New("no_warns_to_remove")

var ErrRoleNotFound = errors.New("role not found")

type RoleNotFoundError struct {
	RoleID string
}

func (e *RoleNotFoundError) Error() string {
	return ErrRoleNotFound.Error()
}
func (e *RoleNotFoundError) Unwrap() error {
	return ErrRoleNotFound
}
