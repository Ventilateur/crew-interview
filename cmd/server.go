package cmd

import (
	"context"
	"fmt"
	"github.com/Ventilateur/crew-interview/api"
	"github.com/Ventilateur/crew-interview/config"
	"github.com/Ventilateur/crew-interview/domain/talent"
	"github.com/Ventilateur/crew-interview/infra/mongodb"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	cfg     config.Config
	cfgFile string
)

var serverCmd = &cobra.Command{
	Use: "server",

	// Load config
	PreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		cfg, err = config.GetConfig(cfgFile)
		if err != nil {
			return fmt.Errorf("failed to load config file '%s'", cfgFile)
		}
		return nil
	},

	// Run server
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo.URI))
		if err != nil {
			return fmt.Errorf("failed to create mongo client: %w", err)
		}
		defer func() {
			_ = mongoClient.Disconnect(ctx)
		}()

		talentRepo := mongodb.NewTalentRepo(mongoClient.
			Database(cfg.Mongo.Database).
			Collection(cfg.Mongo.Collections.Talents))

		talentApp := talent.NewApplication(talentRepo)

		handler := api.NewHTTPHAndler(talentApp)

		server := api.NewServer(cfg.Server, handler)

		server.Serve()

		return nil
	},
}

func initServerCmdFlags() {
	serverCmd.Flags().StringVarP(&cfgFile, "config", "", "", "config file path (required)")
}
