package cmd

import (
	"strings"

	"github.com/mvr-garcia/kafikinha/pkg/kafka"
	"github.com/mvr-garcia/kafikinha/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var message string

var producerCmd = &cobra.Command{
	Use:   "producer",
	Short: "Start the producer",
	Run: func(cmd *cobra.Command, args []string) {
		brokerList := strings.Split(brokers, ",")
		p, err := kafka.NewProducer(brokerList, topic)
		if err != nil {
			logger.L().Fatal("failed to create producer", zap.Error(err))
		}
		defer p.Close()

		if message == "" {
			logger.L().Warn("no message provided, skipping send")
			return
		}

		logger.L().Info("sending message",
			zap.String("topic", topic),
			zap.String("message", message),
		)

		p.SendMessage(message)
	},
}

func init() {
	rootCmd.AddCommand(producerCmd)
	producerCmd.Flags().StringVar(&message, "message", "", "message to send")
}
