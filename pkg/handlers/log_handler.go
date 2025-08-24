package handlers

import (
	"github.com/mvr-garcia/kafikinha/pkg/logger"
	"go.uber.org/zap"
)

type LogHandler struct{}

func NewLogHandler() *LogHandler {
	return &LogHandler{}
}

func (h *LogHandler) ProcessMessage(topic string, partition int32, offset int64, key, value []byte) error {
	logger.L().Info("message processed",
		zap.String("topic", topic),
		zap.Int32("partition", partition),
		zap.Int64("offset", offset),
		zap.ByteString("key", key),
		zap.ByteString("value", value),
	)
	return nil
}
