package calc

import "testing"

func TestAdd(t *testing.T) {
	var result int 
	result = Add(15,10)
	if result!= 25 {
		t.Error("Expected 25, got ", result)
	}
}
func TestSubtract(t *testing.T) {
	var result int
	result = Subtract(15,10)
	if result!= 5 {
		t.Error("Expected 5, got ", result)
	}
}
