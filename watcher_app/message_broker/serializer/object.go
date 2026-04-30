package mb_cud_serializer

type ParseDelveryData interface {
	Parse() ParsedDataMessage
}

type ConsumeDataJson struct {
	TableName string `json:"table_name"`
	Protocol  string `json:"protocol"`
	Role      string `json:"role"`
	Data      interface{}
}

func (c *ConsumeDataJson) Parse() ParsedDataMessage {
	return ParsedDataMessage{
		TableName: c.TableName,
		Protocol:  c.Protocol,
		Data:      c.Data,
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
	TableName string `json:"table_name"`
	Protocol  string `json:"protocol"`
	Data      interface{}
}
