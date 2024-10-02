package cmd

import (
	"github.com/gkwa/bravesalsa/core"
	"github.com/spf13/cobra"
)

var reverse bool

var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Sort file paths based on modification time",
	Long:  `Sort a list of file paths based on their most recent modification time.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		err := core.SortFiles(cmd.InOrStdin(), cmd.OutOrStdout(), reverse)
		if err != nil {
			logger.Error(err, "Failed to sort files")
		}
	},
}

func init() {
	rootCmd.AddCommand(sortCmd)
	sortCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "Sort in reverse order")
}
