package styles

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

func TestStyles(t *testing.T) {
	Println(Styles.Bold.Print("bold"))
	Println(Styles.Dim.Print("dim"))
	Println(Styles.Italic.Print("italic"))
	Println(Styles.Underline.Print("underline"))
	Println(Styles.Inverse.Print("inverse"))
	Println(Styles.Hidden.Print("hidden"))
	Println(Styles.Strikethrough.Print("strikethrough"))

	Println(Styles.Black.Print("black"))
	Println(Styles.Red.Print("red"))
	Println(Styles.Green.Print("green"))
	Println(Styles.Yellow.Print("yellow"))
	Println(Styles.Blue.Print("blue"))
	Println(Styles.Magenta.Print("magenta"))
	Println(Styles.Cyan.Print("cyan"))
	Println(Styles.White.Print("white"))
	Println(Styles.Gray.Print("gray"))

	Println(Styles.BgBlack.Print("bgblack"))
	Println(Styles.BgRed.Print("bgred"))
	Println(Styles.BgGreen.Print("bggreen"))
	Println(Styles.BgYellow.Print("bgyellow"))
	Println(Styles.BgBlue.Print("bgblue"))
	Println(Styles.BgMagenta.Print("bgmagenta"))
	Println(Styles.BgCyan.Print("bgcyan"))
	Println(Styles.BgWhite.Print("bgwhite"))
	Println(Styles.BgGray.Print("bggray"))

	assert.Equal(t, Styles.Green.Open, "\x1b[32m")
	assert.Equal(t, Styles.BgGreen.Open, "\x1b[42m")
	assert.Equal(t, Styles.Green.Close, "\x1b[39m")
}

func Println(s string) {
	fmt.Println(s, Styles.Reset.Print(" "))
}
