package extensions

type BarcodeExtension interface {
	Ean13() string
	Ean8() string
	Isbn10() string
	Isbn13() string
}
