package behaviourflag_test

import (
	"findip/config/behaviourflag"
	"testing"
)

func TestUsageMethodReturnsCorrectText(t *testing.T) {
	flag := behaviourflag.PrintVersionFlag{}
	result := flag.Usage()

	expected := "print the program version"
	if result != expected {
		t.Errorf("expected usage text to be '%s'. got '%s'", expected, result)
	}
}

func TestStringMethodReturnsEmptyString(t *testing.T) {
	flag := behaviourflag.PrintVersionFlag{}
	result := flag.String()

	if result != "" {
		t.Errorf("expected string text to be empty string. got '%s'", result)
	}
}

func TestIsBoolFlagMethodReturnsTrue(t *testing.T) {
	flag := behaviourflag.PrintVersionFlag{}
	result := flag.IsBoolFlag()

	if !result {
		t.Error("expected IsBoolFlag result to be true. got false")
	}
}

func TestIsSetMethodReflectsSetValue(t *testing.T) {
	flag := behaviourflag.PrintVersionFlag{}

	result := flag.IsSet()
	if result {
		t.Error("expected IsSet result to be false. got true")
	}

	flag.Set("")
	result = flag.IsSet()
	if !result {
		t.Error("expected IsSet result to be true. got false")
	}
}

func TestSetMethodDoesntReturnError(t *testing.T) {
	flag := behaviourflag.PrintVersionFlag{}

	err := flag.Set("")

	if err != nil {
		t.Errorf("expected Set methods error to be nil. got '%v'", err)
	}
}
