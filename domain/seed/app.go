package seed

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Application interface {
	Seed(ctx context.Context) error
}

type app struct {
	talentRepo TalentRepository
	crewRepo   CrewRepository
}

func NewApplication(talentRepo TalentRepository, crewRepo CrewRepository) Application {
	return &app{
		talentRepo: talentRepo,
		crewRepo:   crewRepo,
	}
}

func (a *app) Seed(ctx context.Context) error {
	startPage := 0
	pageLimit := 100

	for {
		talents, err := a.crewRepo.ListTalents(startPage, pageLimit)
		if err != nil {
			return fmt.Errorf("failed to get talents from Crew: %w", err)
		}

		if len(talents) == 0 {
			return nil
		}

		err = a.talentRepo.AddTalents(ctx, talents)
		if err != nil {
			return fmt.Errorf("failed add talents: %w", err)
		}
		logrus.Infof("Added %d talents to database", len(talents))

		startPage += 1
	}
}
