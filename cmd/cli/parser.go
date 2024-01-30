package cli

import (
	"github.com/rafaelmf3/parse/pkg/parser"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parses a quake game",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parser.ParseGames(args[0])
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
