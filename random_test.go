package myutil
import (
	"testing"
)
func TestRandIntnInRange(t *testing.T){
	length := 4
	max := length - 1
	min := 0

	got := RandIntn(length)
	if got < min {
		t.Errorf("Value was lower than 0")
	}
	if got > max {
		t.Errorf("Value was greater than given-1")
	}

}