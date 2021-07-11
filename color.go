package duitdraw

import "image/color"

// A Color represents an RGBA value, 8 bits per element. Red is the high 8
// bits, green the next 8 and so on.
type Color uint32

const (
	Opaque        Color = 0x272932FF
	Transparent   Color = 0x00000000 /* only useful for allocimage memfillcolor */
	Black         Color = 0xDFDFDFFF
	White         Color = 0x272932FF
	Red           Color = 0xFF6C6BFF
	Green         Color = 0x98BE65FF
	Blue          Color = 0x51AFEFFF
	Cyan          Color = 0x46D9FFFF
	Magenta       Color = 0xC678DDFF
	Yellow        Color = 0xECBE7BFF
	Paleyellow    Color = 0xECBE7BFF
	Darkyellow    Color = 0xECBE7BFF
	Darkgreen     Color = 0x98BE65FF
	Palegreen     Color = 0x98BE65FF
	Medgreen      Color = 0x98BE65FF
	Darkblue      Color = 0x51AFEFFF
	Palebluegreen Color = 0x51AFEFFF
	Paleblue      Color = 0x46D9FFFF
	Bluegreen     Color = 0x46D9FFFF
	Greygreen     Color = 0x98BE65FF
	Palegreygreen Color = 0x98BE65FF
	Yellowgreen   Color = 0xECBE7BFF
	Medblue       Color = 0x51AFEFFF
	Greyblue      Color = 0x5C6370FF
	Palegreyblue  Color = 0xDCDFE4FF
	Purpleblue    Color = 0xC678DDFF

	Notacolor Color = 0xFFFFFF00
	Nofill    Color = Notacolor
)

func (c Color) rgba() color.RGBA {
	return color.RGBA{
		R: uint8(c >> 24),
		G: uint8(c >> 16),
		B: uint8(c >> 8),
		A: uint8(c),
	}
}
