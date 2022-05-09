package cmd

import (
	"fmt"
	"os"

	"github.com/pincheng0101/likecoin-chain-tx-exporter/export"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "./likecoin-chain-tx-exporter",
	Short: "Export from LikeCoin chain tx data",
	Run: func(cmd *cobra.Command, args []string) {
		likedPath, err := cmd.PersistentFlags().GetString("liked-path")
		if err != nil {
			fmt.Println("Cannot get liked data folder path from command line parameters")
			os.Exit(0)
		}
		export.Run(likedPath)
	},
}

func Execute() {
	if err := exportCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	exportCmd.PersistentFlags().String("liked-path", "./.liked", "location of the LikeCoin chain database folder (.liked)")
}
