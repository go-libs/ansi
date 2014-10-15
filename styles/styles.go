package styles

// http://en.wikipedia.org/wiki/ANSI_escape_code#Colors

import "strconv"

const PREFIX = "\u001b["
const SUFFIX = "m"

type Code [2]int

type Style struct {
	Type, Open, Close string
}

func (s *Style) Print(text string) string {
	return s.Open + text + s.Close
}

var Codes = map[string]Code{
	// Reset
	"Reset": {0, 0},

	// Attributes
	"Bold":          {1, 22},
	"Dim":           {2, 22},
	"Italic":        {3, 23},
	"Underline":     {4, 24},
	"BlinkSlow":     {5, 25},
	"BlinkRapid":    {6, 26},
	"Inverse":       {7, 27},
	"Hidden":        {8, 28},
	"Strikethrough": {9, 29},

	// Foreground colors
	"Black":        {30, 39},
	"Red":          {31, 39},
	"Green":        {32, 39},
	"Yellow":       {33, 39},
	"Blue":         {34, 39},
	"Magenta":      {35, 39},
	"Cyan":         {36, 39},
	"LightGray":    {37, 39},
	"DarkGray":     {90, 39},
	"LightRed":     {91, 39},
	"LightGreen":   {92, 39},
	"LightYellow":  {93, 39},
	"LightBlue":    {94, 39},
	"LightMagenta": {95, 39},
	"LightCyan":    {96, 39},
	"White":        {97, 39},

	// Background colors
	"BgBlack":        {40, 49},
	"BgRed":          {41, 49},
	"BgGreen":        {42, 49},
	"BgYellow":       {43, 49},
	"BgBlue":         {44, 49},
	"BgMagenta":      {45, 49},
	"BgCyan":         {46, 49},
	"BgLightGray":    {47, 49},
	"BgGray":         {90, 49},
	"BgDarkGray":     {100, 49},
	"BgLightRed":     {101, 49},
	"BgLightGreen":   {102, 49},
	"BgLightYellow":  {103, 49},
	"BgLightBlue":    {104, 49},
	"BgLightMagenta": {105, 49},
	"BgLightCyan":    {106, 49},
	"BgWhite":        {107, 49},
}

var (
	Reset         = CreateStyle("Reset")
	Bold          = CreateStyle("Bold")
	Dim           = CreateStyle("Dim")
	Italic        = CreateStyle("Italic")
	Underline     = CreateStyle("Underline")
	BlinkSlow     = CreateStyle("BlinkSlow")
	BlinkRapid    = CreateStyle("BlinkRapid")
	Inverse       = CreateStyle("Inverse")
	Hidden        = CreateStyle("Hidden")
	Strikethrough = CreateStyle("Strikethrough")

	Black        = CreateStyle("Black")
	Red          = CreateStyle("Red")
	Green        = CreateStyle("Green")
	Yellow       = CreateStyle("Yellow")
	Blue         = CreateStyle("Blue")
	Magenta      = CreateStyle("Magenta")
	Cyan         = CreateStyle("Cyan")
	White        = CreateStyle("White")
	LightGray    = CreateStyle("LightGray")
	DarkGray     = CreateStyle("DarkGray")
	LightRed     = CreateStyle("LightRed")
	LightGreen   = CreateStyle("LightGreen")
	LightYellow  = CreateStyle("LightYellow")
	LightBlue    = CreateStyle("LightBlue")
	LightMagenta = CreateStyle("LightMagenta")
	LightCyan    = CreateStyle("LightCyan")
	//Gray    = CreateStyle("Gray")

	BgBlack        = CreateStyle("BgBlack")
	BgRed          = CreateStyle("BgRed")
	BgGreen        = CreateStyle("BgGreen")
	BgYellow       = CreateStyle("BgYellow")
	BgBlue         = CreateStyle("BgBlue")
	BgMagenta      = CreateStyle("BgMagenta")
	BgCyan         = CreateStyle("BgCyan")
	BgWhite        = CreateStyle("BgWhite")
	BgLightGray    = CreateStyle("BgLightGray")
	BgGray         = CreateStyle("BgGray")
	BgDarkGray     = CreateStyle("BgDarkGray")
	BgLightRed     = CreateStyle("BgLightRed")
	BgLightGreen   = CreateStyle("BgLightGreen")
	BgLightYellow  = CreateStyle("BgLightYellow")
	BgLightBlue    = CreateStyle("BgLightBlue")
	BgLightMagenta = CreateStyle("BgLightMagenta")
	BgLightCyan    = CreateStyle("BgLightCyan")
)

func CreateStyle(s string) Style {
	c := Codes[s]
	return Style{
		Type:  s,
		Open:  PREFIX + strconv.Itoa(c[0]) + SUFFIX,
		Close: PREFIX + strconv.Itoa(c[1]) + SUFFIX,
	}
}
