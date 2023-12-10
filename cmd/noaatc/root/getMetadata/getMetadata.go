package getMetadata

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"github.com/spf13/cobra"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getMetadata/harcon"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getMetadata/tidepredoffsets"
)

var GetMetadataCmd = &cobra.Command{
	Use:   "getMetadata",
	Short: "Get data from NOAA CO-OPS Metadata API",
	Long: `Get data from NOAA CO-OPS Metadata API`,
}

func init() {
	GetMetadataCmd.AddCommand(harcon.HarconCmd)
	GetMetadataCmd.AddCommand(tidepredoffsets.TidepredoffsetsCmd)
}

