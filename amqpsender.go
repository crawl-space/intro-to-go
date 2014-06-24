// AmqpSender sends delayd entries over amqp after their timeout
type AmqpSender struct {
	AmqpBase

	C chan<- EntryWrapper
}

// Send sends a delayd entry over AMQP, using the entry's Target as the publish
// exchange.
func (s AmqpSender) Send(e Entry) (err error) {
	msg := amqp.Publishing{
		DeliveryMode:    amqp.Persistent,
		Timestamp:       time.Now(),
		ContentType:     e.ContentType,
		ContentEncoding: e.ContentEncoding,
		CorrelationId:   e.CorrelationID,
		Body:            e.Body,
	}

	err = s.channel.Publish(e.Target, "", true, false, msg)
	return
}
