package basic

import (
	"strings"
	"math"
	"testing"
)

func TestBool(t *testing.T) {
	Bool()
}

func TestNatural(t *testing.T) {
	option := OptionNumber{
		Min: 10,
		Max: 20,
	}
	for i := 0; i < 100; i++ {
		val := Natural(option)
		// t.Log(val)
		if val > option.Max || val < option.Min {
			t.Fail()
		}
	}

	for i := 0; i < 100; i++ {
		val := Natural()
		// t.Log(val)
		if val > math.MaxInt32 || val < 0 {
			t.Fail()
		}
	}
}

func TestInt(t *testing.T) {
	option := OptionNumber{
		Min: -10,
		Max: 20,
	}
	for i := 0; i < 100; i++ {
		val := Int(option)
		// t.Log(val)
		if val > option.Max || val < option.Min {
			t.Fail()
		}
	}

	for i := 0; i < 100; i++ {
		val := Int()
		// t.Log(val)
		if val > math.MaxInt32 || val < math.MinInt32 {
			t.Fail()
		}
	}
}

func TestFloat(t *testing.T) {
	for i := 0; i < 100; i++ {
		val := Float()
		// t.Logf("%f", val)
		if val > math.MaxInt32+1 || val < math.MinInt32-1 {
			t.Fail()
		}
	}

	option := OptionNumber{
		Min:        -10,
		Max:        20,
		MinDecimal: 0,
		MaxDecimal: 2,
	}
	for i := 0; i < 100; i++ {
		val := Float(option)
		// t.Log(val)
		if val > float64(option.Max+1) || val < float64(option.Min-1) {
			t.Fail()
		}
	}
}

func TestChar(t *testing.T) {
	for i := 0; i < 100; i++ {
		Char()
	}

	for i := 0; i < 100; i++ {
		val := Char("test", "me")
		if val != "t" && val != "e" && val != "s" && val != "m" {
			t.Fail()
		}
		// t.Logf("%s", val)
	}
}

func TestString(t *testing.T) {
	for i := 0; i < 100; i++ {
		val := String()
		// t.Log(val)
		if val == "" {
			t.Fail()
		}
	}

	option := OptionString{
		MaxLen:    8,
		MinLen:    8,
		FromChars: strings.Join(NumsStr, "") + strings.Join(Lower, ""),
	}
	for i := 0; i < 100; i++ {
		val := String(option)
		// t.Logf("%s", val)
		if len(val) != 8 {
			t.Fail()
		}
	}
}
