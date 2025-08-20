package cmd

import (
	"fmt"
	"log/slog"

	"github.com/linuxoid69/go-pantry/internal/config"
	"github.com/linuxoid69/go-pantry/internal/dbase"
	"github.com/linuxoid69/go-pantry/internal/logger"
	"github.com/linuxoid69/go-pantry/migrations"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:              "run",
	Short:            "Run server",
	PersistentPreRun: func(cmd *cobra.Command, args []string) { logger.InitLogger() },
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig("pantry_conf.yaml")

		fmt.Println("Start server")
		db := dbase.NewDb(
			viper.GetString("db.type"),
			viper.GetString("db.host"),
			viper.GetString("db.name"),
			viper.GetString("db.user"),
			viper.GetString("db.password"),
		)

		if err := db.RunMigration(migrations.Dir); err != nil {
			slog.Error("migrations error", "error", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
