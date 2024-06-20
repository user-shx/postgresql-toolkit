package crypto

import (
	"testing"
)

func TestPasswd(t *testing.T) {
	for i := 0; i < 10; i++ {
		l := Intn(12)
		if l < 8 { // make sure it's greater than 8
			l += 8
		}
		t.Logf("generating random password of length %d", l)
		p, e := Password(l)
		if e != nil {
			t.Error(e)
		}
		t.Log(p)
		if len(p) != l {
			t.Fail()
		}
	}
}
