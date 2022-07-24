package talent

import "context"

type Application interface {
	ListTalents(ctx context.Context, fromId string, size int) ([]Talent, error)
	AddTalent(ctx context.Context, talent Talent) error
}

type app struct {
	repo Repository
}

func NewApplication(repo Repository) Application {
	return &app{
		repo: repo,
	}
}

func (a *app) ListTalents(ctx context.Context, fromId string, size int) ([]Talent, error) {
	return a.repo.ListTalents(ctx, fromId, size)
}

func (a *app) AddTalent(ctx context.Context, talent Talent) error {
	return a.repo.AddTalent(ctx, talent)
}
