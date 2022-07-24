package talent

import "fmt"

type DuplicateIdError struct {
	id string
}

func NewDuplicateIdError(id string) error {
	return &DuplicateIdError{id: id}
}

func (e *DuplicateIdError) Error() string {
	return fmt.Sprintf("talent id %s has already existed", e.id)
}
