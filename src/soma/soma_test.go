package soma_test

import (
	"soma"
	"testing"
)

func TestSoma(t *testing.T) {
	result := soma.Soma(3, 8)
	if result != 11 {
		t.Errorf("soma(3,8) = %d; want 11", result)
	}
}
