package errors

import (
	"crypto/rand"
	"encoding/base64"
	"math"
	"testing"
)

func TestEqual(t *testing.T) {
	pair := []struct {
		E1     int
		E2     int
		expect bool
	}{
		{10010, 10010, true},
		{20200, 20200, true},
		{30300, 30300, true},
		{40142, 23234, false},
	}

	for _, tc := range pair {
		e1 := NewErr(200, tc.E1, randomBase64String(10))
		e2 := NewErr(200, tc.E2, randomBase64String(10))

		actual := Equal(e1, e2)

		if actual != tc.expect {
			t.Errorf("E1: %v  E2: %v  Expect: %v Actual: %v", *e1, *e2, tc.expect, actual)
		}
	}
}

func TestWrapAndCause(t *testing.T) {
	pair := []struct {
		E1      int
		message string
	}{
		{20010, "create user"},
	}

	for _, tc := range pair {
		E1 := NewErr(200, tc.E1, randomBase64String(10))

		err := Cause(Wrap(E1, tc.message))
		switch e := err.(type) {
		case *Err:
			if e.BusinessCode != tc.E1 {
				t.Errorf("business code error: expect: %d  actual: %d", tc.E1, e.BusinessCode)
			}
			if e.Message != E1.Message {
				t.Errorf("message changed: expect: %s  actual: %s", E1.Message, e.Message)
			}

		default:
			t.Errorf("Cause didn't get type *Err: E1: %v e: %v", *E1, e)
		}

	}
}

func randomBase64String(l int) string {
	buff := make([]byte, int(math.Ceil(float64(l)/float64(1.33333333333))))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:l] // strip 1 extra character we get from odd length results
}
