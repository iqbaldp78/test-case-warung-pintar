package message

import (
	"errors"
	"testing"
)

func TestNewMessage(t *testing.T) {
	err := errors.New("")
	input := []struct {
		code   int
		err    interface{}
		fields []string
	}{
		{4, nil, []string{}},
		{6, err, []string{"test"}},
	}
	output := []string{
		"Wrong parameter",
		"test failed",
	}

	for index := range input {
		result := New(input[index].code, input[index].err, input[index].fields...)
		if result.Message != output[index] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
		}
	}
}

func TestMessageInitial(t *testing.T) {
	input := []Error{
		{"", "001", "", 0},
		{},
	}
	output := []bool{false, true}

	for index := range input {
		result := input[index].IsInitial()
		if result != output[index] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
		}
	}
}
