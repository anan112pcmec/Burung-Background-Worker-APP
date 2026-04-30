package models

type PayOutKurir struct {
	ID               int64  `gorm:"primaryKey;autoIncrement" json:"id_payout_kurir"`
	IdKurir          int64  `gorm:"column:id_kurir;not null" json:"id_kurir"`
	Kurir            Kurir  `gorm:"foreignKey:IdKurir;references:ID" json:"-"`
	IdDisbursment    int64  `gorm:"index;column:id_disbursment;type:int8;not null" json:"id_disbursment"`
	UserId           int    `gorm:"column:user_id;type:int4;not null" json:"user_id"`
	Amount           int    `gorm:"column:amount;type:int4;not null" json:"amount"`
	Status           string `gorm:"column:status;type:varchar(20);not null" json:"status"`
	Reason           string `gorm:"column:reason;type:text" json:"reason"`
	Timestamp        string `gorm:"column:timestamp;type:text;not null" json:"timestamp"`
	BankCode         string `gorm:"column:bank_code;type:varchar(50);not null" json:"bank_code"`
	AccountNumber    string `gorm:"column:account_number;type:varchar(150);not null" json:"account_number"`
	RecipientName    string `gorm:"column:recipient_name;type:varchar(100);not null" json:"recipient_name"`
	SenderBank       string `gorm:"column:sender_bank;type:varchar(50);not null" json:"sender_bank"`
	Remark           string `gorm:"column:remark;type:text" json:"remark"`
	Receipt          string `gorm:"column:receipt;type:text" json:"receipt"`
	TimeServed       string `gorm:"column:time_served;type:text;not null" json:"time_served"`
	BundleId         int64  `gorm:"column:bundle_id;type:int8;not null;default:0" json:"bundle_id"`
	CompanyId        int64  `gorm:"column:company_id;type:int8;not null;default:0" json:"company_id"`
	RecipientCity    int    `gorm:"column:recipient_city;type:int4;not null" json:"recipient_city"`
	CreatedFrom      string `gorm:"column:created_from;type:text" json:"created_from"`
	Direction        string `gorm:"column:direction;type:text;not null" json:"direction"`
	Sender           string `gorm:"column:sender;type:text;not null" json:"sender"`
	Fee              int    `gorm:"column:fee;type:int4;not null" json:"fee"`
	BeneficiaryEmail string `gorm:"column:beneficiary_email;type:varchar(100);not null" json:"beneficiary_email"`
	IdempotencyKey   string `gorm:"column:idempotency_key;type:varchar(100);not null" json:"idempotency_key"`
	IsVirtualAccount bool   `gorm:"column:is_virtual_account;type:bool;not null;default:false" json:"is_virtual_account"`
}

func (PayOutKurir) TableName() string {
	return "payout_kurir"
}

type PayOutSeller struct {
	ID               int64  `gorm:"primaryKey;autoIncrement" json:"id_payout_kurir"`
	IdSeller         int64  `gorm:"column:id_seller;not null" json:"id_seller"`
	Seller           Seller `gorm:"foreignKey:IdSeller;references:ID" json:"-"`
	IdDisbursment    int64  `gorm:"index;column:id_disbursment;type:int8;not null" json:"id_disbursment"`
	UserId           int    `gorm:"column:user_id;type:int4;not null" json:"user_id"`
	Amount           int    `gorm:"column:amount;type:int4;not null" json:"amount"`
	Status           string `gorm:"column:status;type:varchar(20);not null" json:"status"`
	Reason           string `gorm:"column:reason;type:text" json:"reason"`
	Timestamp        string `gorm:"column:timestamp;type:text;not null" json:"timestamp"`
	BankCode         string `gorm:"column:bank_code;type:varchar(50);not null" json:"bank_code"`
	AccountNumber    string `gorm:"column:account_number;type:varchar(150);not null" json:"account_number"`
	RecipientName    string `gorm:"column:recipient_name;type:varchar(100);not null" json:"recipient_name"`
	SenderBank       string `gorm:"column:sender_bank;type:varchar(50);not null" json:"sender_bank"`
	Remark           string `gorm:"column:remark;type:text" json:"remark"`
	Receipt          string `gorm:"column:receipt;type:text" json:"receipt"`
	TimeServed       string `gorm:"column:time_served;type:text;not null" json:"time_served"`
	BundleId         int64  `gorm:"column:bundle_id;type:int8;not null;default:0" json:"bundle_id"`
	CompanyId        int64  `gorm:"column:company_id;type:int8;not null;default:0" json:"company_id"`
	RecipientCity    int    `gorm:"column:recipient_city;type:int4;not null" json:"recipient_city"`
	CreatedFrom      string `gorm:"column:created_from;type:text" json:"created_from"`
	Direction        string `gorm:"column:direction;type:text;not null" json:"direction"`
	Sender           string `gorm:"column:sender;type:text;not null" json:"sender"`
	Fee              int    `gorm:"column:fee;type:int4;not null" json:"fee"`
	BeneficiaryEmail string `gorm:"column:beneficiary_email;type:varchar(100);not null" json:"beneficiary_email"`
	IdempotencyKey   string `gorm:"column:idempotency_key;type:varchar(100);not null" json:"idempotency_key"`
	IsVirtualAccount bool   `gorm:"column:is_virtual_account;type:bool;not null;default:false" json:"is_virtual_account"`
}

func (PayOutSeller) TableName() string {
	return "payout_seller"
}
