package notification_models

type NotificationPengguna struct {
	IDPengguna int64   `json:"id_pengguna"`
	Pengirim   string  `json:"pengirim"`
	Judul      string  `json:"judul"`
	Pesan      string  `json:"pesan"`
	Pop        float32 `json:"pop"`
	Archive    bool    `json:"archive"`
	Inbox      bool    `json:"inbox"`
	Activity   bool    `json:"activity"`
	CreatedAt  string  `json:"created_at"` // Format ISO 8601 (e.g., "2026-06-26T13:00:00Z")
	ExpiredAt  string  `json:"expired_at"` // Opsional: batas waktu event realtime ini valid
	Data       struct {
		Metadata map[string]interface{} `json:"metadata"`
		Special  interface{}            `json:"special"`
	} `json:"data"`
}

// NotificationSeller adalah payload realtime untuk penjual/merchant
type NotificationSeller struct {
	IDSeller  int64   `json:"id_seller"`
	Pengirim  string  `json:"pengirim"`
	Judul     string  `json:"judul"`
	Pesan     string  `json:"pesan"`
	Pop       float32 `json:"pop"`
	Archive   bool    `json:"archive"`
	Inbox     bool    `json:"inbox"`
	Activity  bool    `json:"activity"`
	CreatedAt string  `json:"created_at"`
	ExpiredAt string  `json:"expired_at"`
	Data      struct {
		Metadata map[string]interface{} `json:"metadata"`
		Special  interface{}            `json:"special"`
	} `json:"data"`
}

// NotificationKurir adalah payload realtime untuk driver/kurir
type NotificationKurir struct {
	IDKurir   int64   `json:"id_kurir"`
	Pengirim  string  `json:"pengirim"`
	Judul     string  `json:"judul"`
	Pesan     string  `json:"pesan"`
	Pop       float32 `json:"pop"`
	Archive   bool    `json:"archive"`
	Inbox     bool    `json:"inbox"`
	Activity  bool    `json:"activity"`
	CreatedAt string  `json:"created_at"`
	ExpiredAt string  `json:"expired_at"`
	Data      struct {
		Metadata map[string]interface{} `json:"metadata"`
		Special  interface{}            `json:"special"`
	} `json:"data"`
}
