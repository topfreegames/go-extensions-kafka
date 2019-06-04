package kafka_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/Shopify/sarama"
	"github.com/golang/mock/gomock"
	kafka "github.com/topfreegames/go-extensions-kafka"
	mocks "github.com/topfreegames/go-extensions-kafka/mocks"
)

func TestNewSyncProducerFromClient(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	m := mocks.NewMockSyncProducer(mockCtrl)
	s := kafka.NewSyncProducerFromClient(m)
	if s == nil {
		t.Fatal("Expected SyncProducer to not be nil")
	}
}

func TestSyncProducerWithContext(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	m := mocks.NewMockSyncProducer(mockCtrl)
	s := kafka.NewSyncProducerFromClient(m)
	ctx := context.WithValue(context.Background(), "key", "value")
	ss := s.WithContext(ctx)
	if ss == s {
		t.Fatal("Expected s.WithContext() to return a new *SyncProducer")
	}
	rf := reflect.ValueOf(ss).Elem().FieldByName("ctx")
	ctxI := reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem().Interface()
	val := ctxI.(context.Context).Value("key")
	if _, ok := val.(string); !ok {
		t.Fatal("Expected ctx.Value(`key`) to be string")
	}
	if val.(string) != "value" {
		t.Fatal("Expected ctx.Value(`key`) to be `value`")
	}
}

func TestSyncProducerSendMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	m := mocks.NewMockSyncProducer(mockCtrl)
	s := kafka.NewSyncProducerFromClient(m)
	ctx := context.WithValue(context.Background(), "key", "value")
	ss := s.WithContext(ctx)
	msg := &sarama.ProducerMessage{Topic: "some-topic", Value: sarama.ByteEncoder([]byte("message"))}
	m.EXPECT().SendMessage(msg).Times(1).Return(int32(1), int64(2), nil)
	p, o, err := ss.SendMessage(msg)
	if p != 1 {
		t.Fatalf("Expected partition to be %d. Got %d", 1, p)
	}
	if o != 2 {
		t.Fatalf("Expected partition to be %d. Got %d", 2, p)
	}
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
}

func TestSyncProducerSendMessages(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	m := mocks.NewMockSyncProducer(mockCtrl)
	s := kafka.NewSyncProducerFromClient(m)
	ctx := context.WithValue(context.Background(), "key", "value")
	ss := s.WithContext(ctx)
	msg1 := &sarama.ProducerMessage{Topic: "some-topic", Value: sarama.ByteEncoder([]byte("message"))}
	m.EXPECT().SendMessages([]*sarama.ProducerMessage{msg1}).Times(1).Return(nil)
	var err error
	err = ss.SendMessages([]*sarama.ProducerMessage{msg1})
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	msg2 := &sarama.ProducerMessage{Topic: "some-topic-2", Value: sarama.ByteEncoder([]byte("message"))}
	m.EXPECT().SendMessages([]*sarama.ProducerMessage{msg2}).Times(1).Return(fmt.Errorf("fake fail"))
	err = ss.SendMessages([]*sarama.ProducerMessage{msg2})
	if err == nil {
		t.Fatalf("Expected error: %s", "fake fail")
	}
}

func TestSyncProducerProduce(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	m := mocks.NewMockSyncProducer(mockCtrl)
	s := kafka.NewSyncProducerFromClient(m)
	ctx := context.WithValue(context.Background(), "key", "value")
	ss := s.WithContext(ctx)
	topic := "some-topic"
	message := []byte("some-message")
	m.EXPECT().SendMessage(gomock.Any()).Times(1).Do(func(msg *sarama.ProducerMessage) {
		if msg.Topic != topic {
			t.Fatalf("Expected Topic to be %s. Got %s", topic, msg.Topic)
		}
		val, err := msg.Value.Encode()
		if err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}
		if string(val) != string(message) {
			t.Fatalf("Expected Value to be %s. Got %s", string(message), string(val))
		}
	}).Return(int32(1), int64(2), nil)
	p, o, err := ss.Produce(topic, message)
	if p != 1 {
		t.Fatalf("Expected partition to be %d. Got %d", 1, p)
	}
	if o != 2 {
		t.Fatalf("Expected partition to be %d. Got %d", 2, p)
	}
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
}
