package kafka_client

import (
	"context"
	"github.com/segmentio/kafka-go"
	"strconv"
)
type cudEvent uint8

const (
	createEvent cudEvent = iota
	updateEvent
	deleteEvent
	getEvent
)

type IKafkaClient interface {
	Connect(ctx context.Context, dsn string, topic string, partition int) error
	SendMessage(bytes []byte) error
}

type KafkaClient struct {
	conn *kafka.Conn
}

func NewKafkaClient() IKafkaClient {
	return &KafkaClient{}
}

func (client *KafkaClient) Connect(ctx context.Context, dsn string, topic string, partition int) error {
	conn, err := kafka.DialLeader(ctx, "tcp", dsn, topic, partition)
	if err != nil {
		return err
	}
	client.conn = conn
	return nil
}

func (client *KafkaClient) SendMessage(message []byte) error {
	_, err := client.conn.Write(message)
	return err
}

func SendKafkaCreateEvent(k IKafkaClient) error {
	return k.SendMessage([]byte(strconv.Itoa(int(createEvent))))
}

func SendKafkaUpdateEvent(k IKafkaClient) error {
	return k.SendMessage([]byte(strconv.Itoa(int(updateEvent))))
}

func SendKafkaDeleteEvent(k IKafkaClient) error {
	return k.SendMessage([]byte(strconv.Itoa(int(deleteEvent))))
}

func SendKafkaGetEvent(k IKafkaClient) error {
	return k.SendMessage([]byte(strconv.Itoa(int(getEvent))))
}
