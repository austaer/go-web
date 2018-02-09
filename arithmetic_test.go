package gotest

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("failed")
	} else {
		t.Log("pss")
	}
}

func Test_Multiplication(t *testing.T) {
	if i := Multiplication(12, 12); i != 144 {
		t.Error("failed")
	} else {
		t.Log("pass")
	}
}
