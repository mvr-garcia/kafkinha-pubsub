package cmd

import (
	"github.com/mvr-garcia/kafikinha/pkg/kafka"
	"github.com/mvr-garcia/kafikinha/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Start the consumer",
	Run: func(cmd *cobra.Command, args []string) {
		brokers, _ := cmd.Flags().GetStringSlice("brokers")
		topic, _ := cmd.Flags().GetString("topic")
		groupID, _ := cmd.Flags().GetString("group")

		cons, err := kafka.NewConsumer(brokers, topic, groupID)
		if err != nil {
			logger.L().Fatal("failed to create consumer", zap.Error(err))
		}

		cons.Start()

		// Block until termination
		select {}
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)
	consumerCmd.Flags().StringSlice("brokers", []string{"localhost:9092"}, "Kafka brokers")
	consumerCmd.Flags().String("topic", "events", "Kafka topic to consume from")
	consumerCmd.Flags().String("group", "kafikinha-group", "Kafka consumer group ID")
}
