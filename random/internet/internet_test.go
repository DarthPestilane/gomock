package internet

import (
	"testing"
)

func TestEmail(t *testing.T) {
	for i := 0; i < 100; i++ {
		em := Email()
		if em == "" {
			t.Fail()
		}
	}
	for i := 0; i < 100; i++ {
		em := Email("domain.com")
		if em == "" {
			t.Fail()
		}
	}
}

func TestUrl(t *testing.T) {
	for i := 0; i < 100; i++ {
		u := Url()
		if u == "" {
			t.Fail()
		}
	}
}

func TestMacAddr(t *testing.T) {
	for i := 0; i < 100; i++ {
		u := MacAddr()
		if u == "" || len(u) != 17 {
			t.Fail()
		}
	}
}
