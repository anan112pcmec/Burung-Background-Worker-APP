package mb_cud_serializer

type ConsumeDataJson struct {
	TableName string `json:"table_name"`
	Protocol  string `json:"protocol"`
	Data      interface{}
}

type ConsumeDataProto struct {
	TableName string `json:"table_name"`
	Protocol  string `json:"protocol"`
	Payload   []byte
}
