package cli

import (
	"github.com/rafaelmf3/parse/internal"
	"github.com/spf13/cobra"
)

var reportGamesCmd = &cobra.Command{
	Use:   "gamereport",
	Short: "create games structured data by game report file",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		internal.CreateGameReport()
	},
}

var reportDeathCausesCmd = &cobra.Command{
	Use:   "deathcausereport",
	Short: "create death cause by game report file",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		internal.CreateJsonDeathCauseReport()
	},
}

func init() {
	rootCmd.AddCommand(reportGamesCmd)
	rootCmd.AddCommand(reportDeathCausesCmd)
}
