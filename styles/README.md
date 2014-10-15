# ansi-styles

> [ANSI escape codes](http://en.wikipedia.org/wiki/ANSI_escape_code#Colors_and_Styles) for styling strings in the terminal

## Usage

```go
import "fmt"
import "github.com/go-libs/ansi/styles"

fmt.Println(styles.Red.Open, "Hello world!", styles.Red.Close)
fmt.Println(styles.Green.Print("Go Go Go!"))
fmt.Println(styles.BgRed.Print(styles.White.Print("2014, 10 17")))
```
