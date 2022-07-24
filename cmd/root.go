package cmd

import (
	"github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{}
)

func Execute() {
	initServerCmdFlags()
	initSeedCmdFlags()
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("error starting application")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(seedCmd)
}
