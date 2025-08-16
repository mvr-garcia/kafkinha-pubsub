package cmd

import (
	"github.com/mvr-garcia/kafikinha/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	brokers string
	topic   string
	env     string
)

var rootCmd = &cobra.Command{
	Use:   "kafka-cli",
	Short: "CLI to produce and consumer kafka messages",
	Long:  "A simple CLI (using cobra) to produce and consumer messages from a kafka topic",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.Init(env)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&brokers, "brokers", "localhost:9092", "comma separated list of brokers")
	rootCmd.PersistentFlags().StringVar(&topic, "topic", "events", "Kafka topic to use")
	rootCmd.PersistentFlags().StringVar(&env, "env", "dev", "environment: dev or prod")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.L().Fatal("failed to execute root command", zap.Error(err))
	}
}
