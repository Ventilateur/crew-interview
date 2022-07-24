package mongodb

import (
	"context"
	"fmt"
	"github.com/Ventilateur/crew-interview/domain/seed"
	"go.mongodb.org/mongo-driver/bson"

	talentdomain "github.com/Ventilateur/crew-interview/domain/talent"
	"go.mongodb.org/mongo-driver/mongo"
)

type talentRecord struct {
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

func (t *talentRecord) Entity() talentdomain.Talent {
	entity := talentdomain.Talent{
		Id:        t.Id,
		FirstName: t.FirstName,
		LastName:  t.LastName,
		Picture:   t.Picture,
		Job:       t.Job,
		Location:  t.Location,
		LinkedIn:  t.LinkedIn,
		Github:    t.Github,
		Twitter:   t.Twitter,
		Tags:      make([]string, len(t.Tags)),
		Stage:     t.Stage,
	}

	copy(entity.Tags, t.Tags)

	return entity
}

type talentCollection struct {
	collection *mongo.Collection
}

func NewTalentRepo(collection *mongo.Collection) talentdomain.Repository {
	return &talentCollection{
		collection: collection,
	}
}

func NewSeedTalentRepo(collection *mongo.Collection) seed.TalentRepository {
	return &talentCollection{
		collection: collection,
	}
}

func (t *talentCollection) ListTalents(ctx context.Context, fromId string, size int) ([]talentdomain.Talent, error) {
	cursor, err := t.collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, fmt.Errorf("failed to find talent data from database: %w", err)
	}

	var result []talentdomain.Talent
	for cursor.Next(ctx) {
		var record talentRecord
		err := cursor.Decode(&record)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal talent data from database: %w", err)
		}
		result = append(result, record.Entity())
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("failed to fetch talent data from database: %w", err)
	}

	return result, nil
}

func (t *talentCollection) AddTalent(ctx context.Context, talent talentdomain.Talent) error {
	_, err := t.collection.InsertOne(ctx, &talent)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return talentdomain.NewDuplicateIdError(talent.Id)
		}
		return fmt.Errorf("failed to add talent to database: %w", err)
	}
	return nil
}

func (t *talentCollection) AddTalents(ctx context.Context, talents []seed.Talent) error {
	docs := make([]interface{}, len(talents))
	for i, talent := range talents {
		docs[i] = talent
	}
	_, err := t.collection.InsertMany(ctx, docs)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return talentdomain.NewDuplicateIdError("n/a")
		}
		return fmt.Errorf("failed to add talent to database: %w", err)
	}
	return nil
}
