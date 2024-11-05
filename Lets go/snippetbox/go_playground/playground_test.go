package go_playground

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	result, err := HashPassword("password")
	if err != nil {
		t.Error(err)
	}

	t.Errorf("\n%s\n", result)
}