package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var producerCmd = &cobra.Command{
	Use:   "producer",
	Short: "Start the producer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[producer] brokers=sei la topic=sei la (TODO: implementar)")
	},
}

func init() {
	rootCmd.AddCommand(producerCmd)
}
