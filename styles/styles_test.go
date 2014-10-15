package styles

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

func TestStyles(t *testing.T) {
	Println(Bold.Print("bold"))
	Println(Dim.Print("dim"))
	Println(Italic.Print("italic"))
	Println(Underline.Print("underline"))
	Println(BlinkSlow.Print("blinkslow"))
	Println(BlinkRapid.Print("blinkrapid"))
	Println(Inverse.Print("inverse"))
	Println(Hidden.Print("hidden"))
	Println(Strikethrough.Print("strikethrough"))

	Println(Black.Print("black"))
	Println(Red.Print("red"))
	Println(Green.Print("green"))
	Println(Yellow.Print("yellow"))
	Println(Blue.Print("blue"))
	Println(Magenta.Print("magenta"))
	Println(Cyan.Print("cyan"))
	Println(White.Print("white"))
	Println(Gray.Print("gray"))

	Println(BgBlack.Print("bgblack"))
	Println(BgRed.Print("bgred"))
	Println(BgGreen.Print("bggreen"))
	Println(BgYellow.Print("bgyellow"))
	Println(BgBlue.Print("bgblue"))
	Println(BgMagenta.Print("bgmagenta"))
	Println(BgCyan.Print("bgcyan"))
	Println(BgWhite.Print("bgwhite"))
	Println(BgGray.Print("bggray"))

	assert.Equal(t, Green.Open, "\x1b[32m")
	assert.Equal(t, BgGreen.Open, "\x1b[42m")
	assert.Equal(t, Green.Close, "\x1b[39m")
}

func Println(s string) {
	fmt.Println(s, Reset.Print(" "))
}
