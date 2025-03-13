package functions

import (
	"strconv"
	"strings"
)

type (
	Text  string
	Color string
)

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

const (
	// format = "\033"
	format  = "\x1b"
	format2 = "\033"
	// Text
	Text_Bold          Text = format + "[1m"
	Text_Italic        Text = format + "[3m"
	Text_Underline     Text = format + "[4m"
	Text_Blink         Text = format + "[5m"
	Text_Reversed      Text = format + "[7m"
	Text_Hiddens       Text = format + "[8m"
	Text_Strikethrough Text = format + "[9m"
	// Color
	Color_Black   Text = format + "[38;2;0;0;0m" // ANSI CODE with RGB
	Color_Red     Text = format + "[38;2;255;0;0m"
	Color_Green   Text = format + "[38;2;0;255;0m"
	Color_Yellow  Text = format + "[38;2;255;255;0m"
	Color_Orange  Text = format + "[38;2;255;136;0m"
	Color_Blue    Text = format + "[38;2;0;0;255m"
	Color_Magenta Text = format + "[38;2;120;0;76m"
	Color_Cyan    Text = format + "[38;2;0;162;211m"
	Color_White   Text = format + "[38;2;255;255;255m"
	// Basic Colors
	Color_Black_b   = format2 + "[30;1m"
	Color_Red_b     = format2 + "[31;1m"
	Color_Green_b   = format2 + "[32;1m"
	Color_Yellow_b  = format2 + "[33;1m"
	Color_Orange_b  = format2 + "[34;1m"
	Color_Magenta_b = format2 + "[35;1m"
	Color_Cyan_b    = format2 + "[36;1m"
	Color_White_b   = format2 + "[37;1m"
	Color_Reset     = format2 + "[0m\n"
	// Reset
	ResetAll           Text = format + "[m"
	ResetBold          Text = format + "[22m"
	ResetItalic        Text = format + "[23m"
	ResetUnderline     Text = format + "[24m"
	ResetReversed      Text = format + "[27m"
	ResetHiddens       Text = format + "[28m"
	ResetStrikethrough Text = format + "[29m"
)

func ColorFontRGB(color RGB) string {
	res := (format + "[38;2;" +
		strconv.Itoa(int(color.Red)) + ";" +
		strconv.Itoa(int(color.Green)) + ";" +
		strconv.Itoa(int(color.Blue)) + "m")
	return res
}

func ColorBackGroundRGB(color RGB) string {
	res := (format + "[48;2;" +
		strconv.Itoa(int(color.Red)) + ";" +
		strconv.Itoa(int(color.Green)) + ";" +
		strconv.Itoa(int(color.Blue)) + "m")
	return res
}

func Hex2RGB(hex string) (RGB, error) {
	var rgb RGB
	if hex[0] == '#' {
		hex = strings.Replace(hex, "#", "", 1)
	}
	values, err := strconv.ParseUint(string(hex), 16, 32)
	if err != nil {
		return RGB{}, err
	}

	rgb = RGB{
		Red:   uint8(values >> 16),
		Green: uint8((values >> 8) & 0xFF),
		Blue:  uint8(values & 0xFF),
	}

	return rgb, nil
}
