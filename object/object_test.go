package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Za Wardo"}
	hello2 := &String{Value: "Za Wardo"}
	diff1 := &String{Value: "my name is what"}
	diff2 := &String{Value: "my name is what"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings same content diff hash keys")
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings same content diff hash keys")
	}
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with diff content have same keys")
	}
}
