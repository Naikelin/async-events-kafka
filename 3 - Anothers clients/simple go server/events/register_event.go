package events

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	Topic = "register-event"
)

func RegisterEvent(registerToPublish []byte, partition int) error {

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", Topic, partition)
	if err != nil {
		return err
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		/* kafka.Message{Value: []byte(fmt.Sprint("%v", *toRegister))}, */
		kafka.Message{Value: registerToPublish},
	)
	if err != nil {
		return err
	}

	if err := conn.Close(); err != nil {
		return err
	}

	return nil
}
