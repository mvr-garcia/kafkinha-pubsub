package kafka

type MessageHandler interface {
	ProcessMessage(topic string, partition int32, offset int64, key, value []byte) error
}
