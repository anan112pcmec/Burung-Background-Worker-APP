package media_ekstension

var PhotoValidExt = map[string]bool{"jpg": true, "jpeg": true, "png": true, "webp": true}

var VideoValistExt = map[string]bool{"mp4": true, "mov": true, "avi": true, "wmv": true, "mpeg": true, "mpg": true}

var DokumenValidExt = map[string]bool{
	// PDF
	"pdf": true,

	// Word Documents
	"doc":  true,
	"docx": true,

	// OpenDocument Format
	"odt": true, // text document
	"ods": true, // spreadsheet
	"odp": true, // presentation

	// RTF & Text
	"rtf": true,
	"txt": true,
	"md":  true,

	// Spreadsheet
	"xls":  true,
	"xlsx": true,
	"csv":  true, // tetap dokumen berbasis teks

	// Presentation
	"ppt":  true,
	"pptx": true,

	// Publishing / e-book documents
	"epub": true,
	"mobi": true,
	"azw":  true,

	// Kompresi dokumen (sering dipakai untuk lampiran file laporan)
	"zip": true,
	"rar": true,
	"7z":  true,
}
