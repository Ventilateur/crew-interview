package talent

import "github.com/google/uuid"

type Talent struct {
	Id        string
	FirstName string
	LastName  string
	Picture   string
	Job       string
	Location  string
	LinkedIn  string
	Github    string
	Twitter   string
	Tags      []string
	Stage     string
}

func (t *Talent) GenerateId() {
	// TODO: Replace with correct id generator
	t.Id = uuid.NewString()
}
