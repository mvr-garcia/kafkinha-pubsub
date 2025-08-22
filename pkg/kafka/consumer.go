package kafka

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/mvr-garcia/kafikinha/pkg/logger"
	"go.uber.org/zap"
)

type Consumer struct {
	group   sarama.ConsumerGroup
	topic   string
	groupID string
	done    chan struct{}
}

func NewConsumer(brokers []string, topic, groupID string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	group, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	c := &Consumer{
		group:   group,
		topic:   topic,
		groupID: groupID,
		done:    make(chan struct{}),
	}

	return c, nil
}

func (c *Consumer) Start() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer cancel()
		for {
			if err := c.group.Consume(ctx, []string{c.topic}, c); err != nil {
				logger.L().Error("error consuming", zap.Error(err))
			}

			if ctx.Err() != nil {
				return
			}
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigterm:
		logger.L().Info("shutdown signal received")
		cancel()
	case <-c.done:
		cancel()
	}

	if err := c.Close(); err != nil {
		logger.L().Error("error closing consumer group", zap.Error(err))
	}
}

func (c *Consumer) Close() error {
	close(c.done)
	return c.group.Close()
}

// --- sarama.ConsumerGroupHandler implementation ---

func (c *Consumer) Setup(_ sarama.ConsumerGroupSession) error {
	logger.L().Info("consumer group setup",
		zap.String("groupID", c.groupID),
		zap.String("topic", c.topic))
	return nil
}

func (c *Consumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	logger.L().Info("consumer group cleanup",
		zap.String("groupID", c.groupID),
		zap.String("topic", c.topic))
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		logger.L().Info("message consumed",
			zap.String("topic", message.Topic),
			zap.Int32("partition", message.Partition),
			zap.Int64("offset", message.Offset),
			zap.ByteString("value", message.Value),
		)
		session.MarkMessage(message, "") // offset commit
	}
	return nil
}
