# mbalign

mbalign aligns the string including multibyte character.

## Example

```golang
import (
    "fmt"
    "github.com/yumaito/mbalign"
)

func main() {
    str := "あいうえお"

    alignStrLeft := mbalign.Align(str, 12, '.', mbalign.LeftAlign)
    fmt.Println(alignStrLeft) // あいうえお..

    alignStrRight := mbalign.Align(str, 12, '.', mbalign.RightAlign)
    fmt.Println(alignStrRight) // ..あいうえお

    alignStr := mbalign.Align(str, 9, '.', mbalign.LeftAlign)
    fmt.Println(alignStr) // あいうえ.

    width := mbalign.GetWidth("あいうえお") // 10
}
```
