package testunitario

import "testing"

func TestSuma(t *testing.T) {
	total := Suma(5, 5)

	if total != 10 {
		t.Errorf("Suma Incorrecta, tiene %d y se eséraba %d", total, 10)
	}
}
