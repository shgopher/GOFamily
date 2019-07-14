package TN

import (
	"fmt"
	"testing"
)

var (
	x = []int{1, 55, 33, 21323, 55334, 12232, 4333, 233232, 433, 23, 2332, 32, 23, 3232, 23, 2332323244, 45, 5, 34}
)

func TestTOPKQuickSelect(t *testing.T) {
	fmt.Println(TOPKQuickSelect(x, 2))
}
