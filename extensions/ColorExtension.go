package extensions

type ColorExtension interface {
	HexColor() string
	SafeHexColor() string
	RgbColorAsArray() string
	RgbColor() string
	RgbCssColor() string
	RgbaCssColor() string
	SafeColorName() string
	ColorName() string
	HslColor() string
	HslColorAsArray() []int
}
