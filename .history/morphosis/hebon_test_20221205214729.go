package morphosis

import (
	"fmt"
	"testing"
)

func TestHebonConvert(t *testing.T) {
	hebon :=ToHebon("こんにちは")
	fmt.Println(hebon)
}