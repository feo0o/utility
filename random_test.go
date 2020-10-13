package utility

import "testing"

func TestRandomStr(t *testing.T) {
	t.Log(RandomString(6))
}

func TestRandomInt(t *testing.T) {
	for i := 0; i <= 20; i++ {
		r, err := RandomInt(10, 20)
		if err != nil {
			t.Error(err)
		}
		t.Log(r)
	}
}
