package compression

import (
	"fmt"
	"testing"
)

func TestSimpleCompression(t *testing.T) {
	fmt.Println("*** starting!!! ***")
	str := "aaabbbbbbcccccccdeeefghjiiiijjppjjpppiiioowee324143542344343"

	fmt.Println(SimpleCompression(str))
}
