package internet

import (
	"strings"
	"testing"
)

func TestEmail(t *testing.T) {
	for i := 0; i < 100; i++ {
		em := Email()
		// t.Log(em)
		if em == "" {
			t.Fail()
		}
	}
	for i := 0; i < 100; i++ {
		em := Email("domain.com")
		// t.Log(em)
		if em == "" || !strings.HasSuffix(em, "domain.com") {
			t.Fail()
		}
	}
}

func TestUrl(t *testing.T) {
	for i := 0; i < 100; i++ {
		u := Url()
		// t.Log(u)
		if u == "" {
			t.Fail()
		}
	}
}

func TestMacAddr(t *testing.T) {
	for i := 0; i < 100; i++ {
		u := MacAddr()
		// t.Log(u)
		if u == "" || len(u) != 17 {
			t.Fail()
		}
	}
}

func TestIPv4(t *testing.T) {
	for i := 0; i < 100; i++ {
		u := IPv4()
		// t.Log(u)
		if u == "" || len(u) > 15 || len(u) < 7 {
			t.Fail()
		}
	}
}

func TestIPv6(t *testing.T) {
	for i := 0; i < 100; i++ {
		u := IPv6()
		// t.Log(u)
		if u == "" || len(u) != 39 {
			t.Fail()
		}
	}
}

func TestLocalIPv4(t *testing.T) {
	for i := 0; i < 100; i++ {
		u := LocalIPv4()
		// t.Log(u)
		if u == "" || len(u) < 8 || len(u) > 15 {
			t.Fail()
		}
	}
}

func Test_long2ip(t *testing.T) {
	if long2ip(16777216) != "1.0.0.0" {
		t.Fail()
	}
}

func Test_ip2long(t *testing.T) {
	if ip2long("1.0.0.0") != 16777216 {
		t.Fail()
	}
}

func Test_split(t *testing.T) {
	from := "12345678"
	if split(from, 2, ":") != "12:34:56:78" {
		t.Fail()
	}
}
