package mb_cud_serializer

type ParseDelveryData interface {
	Parse() ParsedDataMessage
}

type ConsumeDataJson struct {
	TableName string      `json:"table_name"`
	Protocol  string      `json:"protocol"`
	Role      string      `json:"role"`
	Payload   interface{} `json:"payload"`
}

func (c *ConsumeDataJson) Parse() ParsedDataMessage {
	return ParsedDataMessage{
		TableName: c.TableName,
		Protocol:  c.Protocol,
		Role:      c.Role,
		Data:      c.Payload,
	}
}

type ConsumeDataProto struct {
	TableName string `json:"table_name"`
	Protocol  string `json:"protocol"`
	Role      string `json:"role"`
	Data      []byte
}

func (c *ConsumeDataProto) Parse() ParsedDataMessage {
	return ParsedDataMessage{
		TableName: c.TableName,
		Protocol:  c.Protocol,
		Data:      c.Data,
	}
}

type ParsedDataMessage struct {
	TableName string      `json:"table_name"`
	Protocol  string      `json:"protocol"`
	Role      string      `json:"role"`
	Data      interface{} `json:"payload"`
}
