package iteration

import "testing"

func Test_BinaryGap(t *testing.T) {
	if r := BinaryGap(-1); r == 0 {
		t.Log("Success the test for input value < 0.")
	} else {
		t.Error("Fail the test for input value < 0.")
	}

	if r := BinaryGap(32); r == 0 {
		t.Log("Success the test for input value = 32.")
	} else {
		t.Error("Fail the test for input value = 32.")
	}

	if r := BinaryGap(15); r == 0 {
		t.Log("Success the test for input value = 15.")
	} else {
		t.Error("Fail the test for input value = 15.")
	}

	if r := BinaryGap(1041); r == 5 {
		t.Log("Success the test for input value = 1041.")
	} else {
		t.Error("Fail the test for input value = 1041.")
	}

	if r := BinaryGap(1043); r == 5 {
		t.Log("Success the test for input value = 1043.")
	} else {
		t.Error("Fail the test for input value = 1043.")
	}
}
