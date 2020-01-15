package KATANAME

import "testing"

func TestKATANAME(t *testing.T) {

	actual := KATANAME()
	expected := "example"

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

}
