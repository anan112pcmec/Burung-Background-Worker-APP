package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	"github.com/rabbitmq/amqp091-go"
)

func DecodeDeliveryBody(data amqp091.Delivery, dst interface{}) error {

	dec := json.NewDecoder(bytes.NewReader(data.Body))
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}

	return nil
}
