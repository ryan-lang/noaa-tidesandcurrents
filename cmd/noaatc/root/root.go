package root

import (
	"os"

	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getMetadata"

	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use:   "noaatc",
	Short: "An unofficial CLI for the NOAA CO-OPS API",
	Long: `An unofficial CLI for the NOAA CO-OPS API.

Example usage:
noaatc getData predictions --station-id 9447130 --begin yesterday --hours 24 --interval 60 --units feet
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(getData.GetDataCmd)
	rootCmd.AddCommand(getMetadata.GetMetadataCmd)
}
