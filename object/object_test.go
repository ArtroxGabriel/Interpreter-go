package object_test

import (
	"testing"

	"github.com/ArtroxGabriel/interpreter/object"
)

func TestStringHashKey(t *testing.T) {
	hello1 := &object.String{Value: "Hello World"}
	hello2 := &object.String{Value: "Hello World"}
	diff1 := &object.String{Value: "My name is Johnny"}
	diff2 := &object.String{Value: "My name is Johnny"}
	num1 := &object.String{Value: "1"}
	num2 := &object.Integer{Value: 1}

	if hello1.HashKey() != hello2.HashKey() {
		t.Error("string with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Error("string with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Error("string with different content have same hash keys")
	}

	if num1.HashKey() == num2.HashKey() {
		t.Error("Integer and String have same hash keys")
	}
}
