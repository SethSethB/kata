package kataname

import "testing"

func TestKataName(t *testing.T) {

	actual := kataName()
	expected := "example"

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

}
