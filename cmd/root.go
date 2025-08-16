package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	brokers string
	topic   string
	rootCmd = &cobra.Command{
		Use:   "kafka-cli",
		Short: "CLI to produce and consumer kafka messages",
		Long:  "A simple CLI (using cobra) to produce and consumer messages from a kafka topic",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&brokers, "brokers", "localhost:9092", "brokers list (ex: localhost:9092)")
	rootCmd.PersistentFlags().StringVar(&topic, "topics", "event", "kafka topic to be used")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
