package cmd

import (
	"github.com/mvr-garcia/kafikinha/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Start the consumer",
	Run: func(cmd *cobra.Command, args []string) {
		logger.L().Info("starting consumer",
			zap.String("brokers", brokers),
			zap.String("topic", topic),
		)
	},
}

func init() {
	rootCmd.AddCommand(consumerCmd)
}
