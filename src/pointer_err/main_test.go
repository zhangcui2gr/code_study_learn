package pointer_err

import "testing"

func Test_main(t *testing.T) {
	w1 := Wallet{}
	w1.deposit(Bitcoin(10))

	got := w1.Balance()
	want := Bitcoin(5)

	if got != want {
		t.Errorf("got : %s, want : %s", got, want)
	}
}
