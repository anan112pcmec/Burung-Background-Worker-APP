package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/rabbitmq/amqp091-go"

	mb_cud_serializer "github.com/anan112pcmec/Burung-backend-2/watcher_app/message_broker/serializer"
)

func ParseDataMessage[T mb_cud_serializer.ConsumeDataJson | mb_cud_serializer.ConsumeDataProto](Data amqp091.Delivery) (T, bool) {
	var datanya T

	if err := DecodeDeliveryBody(Data, &datanya); err != nil {
		fmt.Println("Jenis Data Tidak Diketahui")

		return datanya, false
	}

	return datanya, true
}

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

func DecodeJSONBody(data mb_cud_serializer.ParsedDataMessage, dst interface{}) error {

	var payloadBytes []byte

	switch v := data.Data.(type) {
	case []byte:
		payloadBytes = v

	case string:
		payloadBytes = []byte(v)

	default:
		return fmt.Errorf("unsupported payload type: %T", data.Data)
	}

	dec := json.NewDecoder(bytes.NewReader(payloadBytes))
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}

	return nil
}
