package talent

import "context"

type Repository interface {
	ListTalents(ctx context.Context, fromId string, size int) ([]Talent, error)
	AddTalent(ctx context.Context, talent Talent) error
}
