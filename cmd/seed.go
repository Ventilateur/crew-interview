package cmd

import (
	"context"
	"fmt"
	"github.com/Ventilateur/crew-interview/domain/seed"
	"github.com/Ventilateur/crew-interview/infra/crewapi"
	"github.com/Ventilateur/crew-interview/infra/mongodb"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI   string
	database   string
	collection string
	crewURI    string
)

var seedCmd = &cobra.Command{
	Use: "seed",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
		if err != nil {
			return fmt.Errorf("failed to create mongo client: %w", err)
		}
		defer func() {
			_ = mongoClient.Disconnect(ctx)
		}()

		talentRepo := mongodb.NewSeedTalentRepo(mongoClient.
			Database(database).
			Collection(collection))

		crewRepo := crewapi.NewTalentAPIClient(crewURI)

		seedApp := seed.NewApplication(talentRepo, crewRepo)

		return seedApp.Seed(ctx)
	},
}

func initSeedCmdFlags() {
	seedCmd.Flags().StringVarP(&mongoURI, "mongo-uri", "", "", "Mongo's URI")
	seedCmd.Flags().StringVarP(&database, "database", "", "", "Database to seed")
	seedCmd.Flags().StringVarP(&collection, "collection", "", "", "Collection to seed")
	seedCmd.Flags().StringVarP(&crewURI, "crew-uri", "", "", "Crew's URI")
	_ = seedCmd.MarkFlagRequired("mongo-uri")
	_ = seedCmd.MarkFlagRequired("database")
	_ = seedCmd.MarkFlagRequired("collection")
	_ = seedCmd.MarkFlagRequired("crew-uri")
}
