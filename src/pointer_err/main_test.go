package pointer_err

import "testing"

//func Test_main(t *testing.T) {
//	w1 := Wallet{}
//	w1.deposit(Bitcoin(10))
//
//	got := w1.Balance()
//	want := Bitcoin(5)
//
//	if got != want {
//		t.Errorf("got : %s, want : %s", got, want)
//	}
//}

func Test_main2(t *testing.T) {
	assertNoEqual := func(got, want Bitcoin, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("got : %s, want : %s", got, want)
		}
	}

	assertnoError := func(t *testing.T, err error) {
		if err == nil {
			t.Error("want a err but didnt get")
		}
	}
	assertInsufficentError := func(t *testing.T, err error, err2 error) {
		if err != err2 {
			t.Errorf("want '%s' , got '%s'", err, err2)
		}
	}

	t.Run("test Deposit", func(t *testing.T) {
		w1 := Wallet{}
		w1.deposit(Bitcoin(10))

		got := w1.Balance()
		want := Bitcoin(10)

		assertNoEqual(got, want, t)
	})

	t.Run("test Withdraw", func(t *testing.T) {
		w2 := Wallet{}
		w2.deposit(Bitcoin(10))
		err := w2.Withdraw(Bitcoin(11))

		got := w2.Balance()
		want := Bitcoin(5)

		assertnoError(t, err)
		assertNoEqual(got, want, t)
		assertInsufficentError(t, err, insufficetMoneyErr)
	})
}
