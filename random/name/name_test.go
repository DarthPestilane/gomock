package name

import (
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	randName := Name()
	if randName == "" {
		t.Fail()
	}
	option := Option{MiddleName: "midname"}
	withMidName := Name(option)
	if !strings.Contains(withMidName, "midname") {
		t.Fail()
		t.Log(withMidName)
	}
	option.Gender = GenderFemale
	femaleName := Name(option)
	first := strings.Split(femaleName, " ")[0]
	var ok bool
	for _, n := range FirstNamesFemale {
		if first == n {
			ok = true
			break
		}
	}
	if !ok {
		t.Fail()
		t.Log(femaleName)
	}
}
