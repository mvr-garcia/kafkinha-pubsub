package kafka

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/mvr-garcia/kafikinha/pkg/logger"
	"go.uber.org/zap"
)

type Producer struct {
	asyncProducer sarama.AsyncProducer
	topic         string
	done          chan struct{}
}

func NewProducer(brokers []string, topic string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Idempotent = true
	config.Net.MaxOpenRequests = 1

	asyncProducer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	p := &Producer{
		asyncProducer: asyncProducer,
		topic:         topic,
		done:          make(chan struct{}),
	}

	// Handle errors and successes asynchronously
	go func() {
		for {
			select {
			case err := <-asyncProducer.Errors():
				if err != nil {
					logger.L().Error("failed to send message", zap.Error(err))
				}
			case success := <-asyncProducer.Successes():
				if success != nil {
					logger.L().Info("message delivered",
						zap.String("topic", success.Topic),
						zap.Int32("partition", success.Partition),
						zap.Int64("offset", success.Offset),
					)
				}
			case <-p.done:
				return
			}
		}
	}()

	return p, nil
}

func (p *Producer) SendMessage(value string) {
	msg := sarama.ProducerMessage{
		Topic:     p.topic,
		Value:     sarama.StringEncoder(value),
		Timestamp: time.Now(),
	}

	p.asyncProducer.Input() <- &msg
}

func (p *Producer) Close() error {
	close(p.done)
	return p.asyncProducer.Close()
}
