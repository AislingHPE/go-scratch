package regex

import (
	"fmt"
	"regexp"
)

func Reg() {
	re := regexp.MustCompile(".*\\.fifo$") //nolint:gosimple
	a := re.Find([]byte("aslefjas.fifo"))
	fmt.Print(a)
}
