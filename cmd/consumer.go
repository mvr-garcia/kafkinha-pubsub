package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Start the consumer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("[consumer] brokers=%s topic=%s (TODO: implementar)\n", brokers, topic)
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)
}
