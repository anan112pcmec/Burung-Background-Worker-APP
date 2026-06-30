package notification_error

import "errors"

var (
	ErrorGagalKirim     = errors.New("gagal kirim ke notificationing app")
	ErrorGagalParsing   = errors.New("gagal parse data notificationing")
	ErrorDataTidakCocok = errors.New("gagal data tidak cocok")
)
