package cmd

import "testing"

func testConvertToCamelCase(t *testing.T, arg, expected string) {
	actual := convertToCamelCase(arg)
	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func TestConvertToCamelCaseSuite(t *testing.T) {
	testConvertToCamelCase(t, "simple", "simple")
	testConvertToCamelCase(t, "less simple", "lessSimple")
	testConvertToCamelCase(t, "CasE mIXture namE", "caseMixtureName")
	testConvertToCamelCase(t, "MultiPle    spaceS", "multipleSpaces")
}

func TestConvertLowerCamelCaseToUpperSuite(t *testing.T) {
	actual := convertLowerCamelCaseToUpper("camelCase")
	expected := "CamelCase"
	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
