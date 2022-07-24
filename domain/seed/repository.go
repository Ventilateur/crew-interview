package seed

import "context"

type TalentRepository interface {
	AddTalents(ctx context.Context, talents []Talent) error
}

type CrewRepository interface {
	ListTalents(page int, limit int) ([]Talent, error)
}
