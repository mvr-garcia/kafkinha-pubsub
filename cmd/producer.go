package cmd

import (
	"github.com/mvr-garcia/kafikinha/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var producerCmd = &cobra.Command{
	Use:   "producer",
	Short: "Start the producer",
	Run: func(cmd *cobra.Command, args []string) {
		logger.L().Info("starting producer",
			zap.String("brokers", brokers),
			zap.String("topic", topic),
		)
	},
}

func init() {
	rootCmd.AddCommand(producerCmd)
}
