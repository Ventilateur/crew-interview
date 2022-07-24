package api

import (
	entity "github.com/Ventilateur/crew-interview/domain/talent"
	"github.com/Ventilateur/crew-interview/utils"
)

type addTalentRequest struct {
	FirstName *string  `json:"firstName"`
	LastName  *string  `json:"lastName"`
	Picture   *string  `json:"picture"`
	Job       *string  `json:"job"`
	Location  *string  `json:"location"`
	LinkedIn  *string  `json:"linkedin"`
	Github    *string  `json:"github"`
	Twitter   *string  `json:"twitter"`
	Tags      []string `json:"tags"`
	Stage     *string  `json:"stage"`
}

func (t *addTalentRequest) Entity() entity.Talent {
	e := entity.Talent{
		FirstName: utils.PtrString(t.FirstName),
		LastName:  utils.PtrString(t.LastName),
		Picture:   utils.PtrString(t.Picture),
		Job:       utils.PtrString(t.Job),
		Location:  utils.PtrString(t.Location),
		LinkedIn:  utils.PtrString(t.LinkedIn),
		Github:    utils.PtrString(t.Github),
		Twitter:   utils.PtrString(t.Twitter),
		Tags:      make([]string, len(t.Tags)),
		Stage:     utils.PtrString(t.Stage),
	}
	e.GenerateId()
	copy(e.Tags, t.Tags)

	return e
}

type addTalentResponse struct {
	Id string `json:"id"`
}

type getTalentResponse struct {
	Id        *string  `json:"id"`
	FirstName *string  `json:"firstName"`
	LastName  *string  `json:"lastName"`
	Picture   *string  `json:"picture"`
	Job       *string  `json:"job"`
	Location  *string  `json:"location"`
	LinkedIn  *string  `json:"linkedin"`
	Github    *string  `json:"github"`
	Twitter   *string  `json:"twitter"`
	Tags      []string `json:"tags"`
	Stage     *string  `json:"stage"`
}

type status string

const (
	statusOK status = "ok"
)

type statusResponse struct {
	Status status `json:"status"`
}
