package styles

import "strconv"

const PREFIX = "\u001b["
const SUFFIX = "m"

type Code [2]int

type Style struct {
	Open,
	Close string
}

func (s *Style) Print(text string) string {
	return s.Open + text + s.Close
}

type styles struct {
	Reset,
	Bold,
	Dim,
	Italic,
	Underline,
	Inverse,
	Hidden,
	Strikethrough,

	Black,
	Red,
	Green,
	Yellow,
	Blue,
	Magenta,
	Cyan,
	White,
	Gray,

	BgBlack,
	BgRed,
	BgGreen,
	BgYellow,
	BgBlue,
	BgMagenta,
	BgCyan,
	BgWhite,
	BgGray Style
}

var Codes = map[string]Code{
	"Reset":         {0, 0},
	"Bold":          {1, 22},
	"Dim":           {2, 22},
	"Italic":        {3, 23},
	"Underline":     {4, 24},
	"Inverse":       {7, 27},
	"Hidden":        {8, 28},
	"Strikethrough": {9, 29},

	"Black":   {30, 39},
	"Red":     {31, 39},
	"Green":   {32, 39},
	"Yellow":  {33, 39},
	"Blue":    {34, 39},
	"Magenta": {35, 39},
	"Cyan":    {36, 39},
	"White":   {37, 39},
	"Gray":    {90, 39},

	"BgBlack":   {40, 49},
	"BgRed":     {41, 49},
	"BgGreen":   {42, 49},
	"BgYellow":  {43, 49},
	"BgBlue":    {44, 49},
	"BgMagenta": {45, 49},
	"BgCyan":    {46, 49},
	"BgWhite":   {47, 49},
	"BgGray":    {90, 49},
}

var Styles = styles{
	Reset:         CreateStyle("Reset"),
	Bold:          CreateStyle("Bold"),
	Dim:           CreateStyle("Dim"),
	Italic:        CreateStyle("Italic"),
	Underline:     CreateStyle("Underline"),
	Inverse:       CreateStyle("Inverse"),
	Hidden:        CreateStyle("Hidden"),
	Strikethrough: CreateStyle("Strikethrough"),

	Black:   CreateStyle("Black"),
	Red:     CreateStyle("Red"),
	Green:   CreateStyle("Green"),
	Yellow:  CreateStyle("Yellow"),
	Blue:    CreateStyle("Blue"),
	Magenta: CreateStyle("Magenta"),
	Cyan:    CreateStyle("Cyan"),
	White:   CreateStyle("White"),
	Gray:    CreateStyle("Gray"),

	BgBlack:   CreateStyle("BgBlack"),
	BgRed:     CreateStyle("BgRed"),
	BgGreen:   CreateStyle("BgGreen"),
	BgYellow:  CreateStyle("BgYellow"),
	BgBlue:    CreateStyle("BgBlue"),
	BgMagenta: CreateStyle("BgMagenta"),
	BgCyan:    CreateStyle("BgCyan"),
	BgWhite:   CreateStyle("BgWhite"),
	BgGray:    CreateStyle("BgGray"),
}

func CreateStyle(s string) Style {
	c := Codes[s]
	return Style{
		Open:  PREFIX + strconv.Itoa(c[0]) + SUFFIX,
		Close: PREFIX + strconv.Itoa(c[1]) + SUFFIX,
	}
}
