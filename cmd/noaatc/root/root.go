package root

import (
	"os"

	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData"
	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use:   "noaatc",
	Short: "An unofficial CLI for the NOAA CO-OPS API",
	Long: `An unofficial CLI for the NOAA CO-OPS API.

Example usage:
noaatc getData predictions --station-id 9414290 --begin yesterday --hours 24 --interval 60 --units feet
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(getData.GetDataCmd)
	// rootCmd.AddCommand(predict.PredictCmd)
}
