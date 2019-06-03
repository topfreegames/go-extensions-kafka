package kafka

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/opentracing/opentracing-go"
	tracing "github.com/topfreegames/go-extensions-tracing"
)

type SyncProducer struct {
	sarama.SyncProducer
	ctx context.Context
}

func NewSyncProducer(brokers []string, kafkaConfig *sarama.Config) (*SyncProducer, error) {
	p, err := sarama.NewSyncProducer(brokers, kafkaConfig)
	if err != nil {
		return nil, err
	}
	return &SyncProducer{SyncProducer: p, ctx: context.Background()}, nil
}

func NewSyncProducerFromClient(c sarama.SyncProducer) *SyncProducer {
	return &SyncProducer{SyncProducer: c, ctx: context.Background()}
}

func (s *SyncProducer) WithContext(ctx context.Context) *SyncProducer {
	ss := *s
	ss.ctx = ctx
	return &ss
}

func (s *SyncProducer) Produce(topic string, message []byte) (int32, int64, error) {
	m := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}
	return s.SendMessage(m)
}

func (s *SyncProducer) SendMessage(msg *sarama.ProducerMessage) (int32, int64, error) {
	span := s.buildSpan()
	defer span.Finish()
	defer tracing.LogPanic(span)

	partition, offset, err := s.SyncProducer.SendMessage(msg)
	if err != nil {
		message := err.Error()
		tracing.LogError(span, message)
	}
	return partition, offset, err
}

func (s *SyncProducer) SendMessages(msgs []*sarama.ProducerMessage) error {
	span := s.buildSpan()
	defer span.Finish()
	defer tracing.LogPanic(span)

	err := s.SyncProducer.SendMessages(msgs)
	if err != nil {
		message := err.Error()
		tracing.LogError(span, message)
	}
	return err
}

func (s *SyncProducer) buildSpan() opentracing.Span {
	var parent opentracing.SpanContext

	if span := opentracing.SpanFromContext(s.ctx); span != nil {
		parent = span.Context()
	}

	operationName := fmt.Sprintf("Kafka Produce")
	reference := opentracing.ChildOf(parent)
	tags := opentracing.Tags{
		"span.kind": "client",
	}

	return opentracing.StartSpan(operationName, reference, tags)
}
