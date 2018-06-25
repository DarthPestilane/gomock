package gomock

import (
	"github.com/DarthPestilane/gomock/random/basic"
	"github.com/DarthPestilane/gomock/random/internet"
	"github.com/DarthPestilane/gomock/random/name"
	"math"
)

type Mocker struct{}

func NewMock() *Mocker {
	return new(Mocker)
}

func (*Mocker) Name(middleName ...string) string {
	var n string
	if len(middleName) > 0 {
		n = middleName[0]
	}
	option := name.Option{
		MiddleName: n,
	}
	return name.Name(option)

}

func (m *Mocker) NameMale(middleName ...string) string {
	var n string
	if len(middleName) > 0 {
		n = middleName[0]
	}
	option := name.Option{
		Gender:     name.GenderMale,
		MiddleName: n,
	}
	return name.Name(option)
}

func (m *Mocker) NameFemale(middleName ...string) string {
	var n string
	if len(middleName) > 0 {
		n = middleName[0]
	}
	option := name.Option{
		Gender:     name.GenderFemale,
		MiddleName: n,
	}
	return name.Name(option)
}

func (m *Mocker) Bool() bool {
	return basic.Bool()
}

func (m *Mocker) Int() int {
	return basic.Int()
}

func (m *Mocker) IntBetween(begin, end int) int {
	return basic.Int(basic.OptionNumber{Min: begin, Max: end})
}

func (m *Mocker) Natural() int {
	return basic.Natural()
}

func (m *Mocker) NaturalBetween(begin, end int) int {
	return basic.Natural(basic.OptionNumber{Max: end, Min: begin})
}

func (m *Mocker) Float() float64 {
	return basic.Float()
}

func (m *Mocker) FloatBetween(begin, end int) float64 {
	return basic.Float(basic.OptionNumber{Max: end, Min: begin})
}

func (m *Mocker) FloatWithDecimal(minDecimalLen, maxDecimalLen int) float64 {
	return basic.Float(basic.OptionNumber{Max: math.MaxInt32, Min: math.MinInt32, MaxDecimal: maxDecimalLen, MinDecimal: minDecimalLen})
}

func (m *Mocker) FloatBetweenWithDecimal(begin, end, minDecimalLen, maxDecimalLen int) float64 {
	return basic.Float(basic.OptionNumber{Max: end, Min: begin, MaxDecimal: maxDecimalLen, MinDecimal: minDecimalLen})
}

func (m *Mocker) Char() string {
	return basic.Char()
}

func (m *Mocker) CharFrom(str string) string {
	return basic.Char(str)
}

func (m *Mocker) String() string {
	return basic.String()
}

func (m *Mocker) StringFrom(str string) string {
	return basic.String(basic.OptionString{FromChars: str})
}

func (m *Mocker) StringWithLength(min, max int) string {
	return basic.String(basic.OptionString{MinLen: min, MaxLen: max})
}

func (m *Mocker) StringWithLengthFrom(min, max int, str string) string {
	return basic.String(basic.OptionString{MinLen: min, MaxLen: max, FromChars: str})
}

func (m *Mocker) Email() string {
	return internet.Email()
}

func (m *Mocker) EmailWithDomain(domain string) string {
	return internet.Email(domain)
}

func (m *Mocker) DomainName() string {
	return internet.DomainName()
}

func (m *Mocker) DomainWord() string {
	return internet.DomainWord()
}

func (m *Mocker) Url() string {
	return internet.Url()
}
