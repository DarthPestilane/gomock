package gofake

import (
	"github.com/DarthPestilane/gofake/random/basic"
	"github.com/DarthPestilane/gofake/random/internet"
	"github.com/DarthPestilane/gofake/random/name"
	"math"
)

type Faker struct{}

func NewFaker() *Faker {
	return new(Faker)
}

func (*Faker) Name(middleName ...string) string {
	var n string
	if len(middleName) > 0 {
		n = middleName[0]
	}
	option := name.Option{
		MiddleName: n,
	}
	return name.Name(option)
}

func (m *Faker) NameMale(middleName ...string) string {
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

func (m *Faker) NameFemale(middleName ...string) string {
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

func (m *Faker) Bool() bool {
	return basic.Bool()
}

func (m *Faker) Int() int {
	return basic.Int()
}

func (m *Faker) IntBetween(begin, end int) int {
	return basic.Int(basic.OptionNumber{Min: begin, Max: end})
}

func (m *Faker) Natural() int {
	return basic.Natural()
}

func (m *Faker) NaturalBetween(begin, end int) int {
	return basic.Natural(basic.OptionNumber{Max: end, Min: begin})
}

func (m *Faker) Float() float64 {
	return basic.Float()
}

func (m *Faker) FloatBetween(begin, end int) float64 {
	return basic.Float(basic.OptionNumber{Max: end, Min: begin})
}

func (m *Faker) FloatWithDecimal(minDecimalLen, maxDecimalLen int) float64 {
	return basic.Float(basic.OptionNumber{Max: math.MaxInt32, Min: math.MinInt32, MaxDecimal: maxDecimalLen, MinDecimal: minDecimalLen})
}

func (m *Faker) FloatBetweenWithDecimal(begin, end, minDecimalLen, maxDecimalLen int) float64 {
	return basic.Float(basic.OptionNumber{Max: end, Min: begin, MaxDecimal: maxDecimalLen, MinDecimal: minDecimalLen})
}

func (m *Faker) Char() string {
	return basic.Char()
}

func (m *Faker) CharFrom(str string) string {
	return basic.Char(str)
}

func (m *Faker) String() string {
	return basic.String()
}

func (m *Faker) StringFrom(str string) string {
	return basic.String(basic.OptionString{FromChars: str})
}

func (m *Faker) StringWithLength(min, max int) string {
	return basic.String(basic.OptionString{MinLen: min, MaxLen: max})
}

func (m *Faker) StringWithLengthFrom(min, max int, str string) string {
	return basic.String(basic.OptionString{MinLen: min, MaxLen: max, FromChars: str})
}

func (m *Faker) Email() string {
	return internet.Email()
}

func (m *Faker) EmailWithDomain(domain string) string {
	return internet.Email(domain)
}

func (m *Faker) DomainName() string {
	return internet.DomainName()
}

func (m *Faker) DomainWord() string {
	return internet.DomainWord()
}

func (m *Faker) Url() string {
	return internet.Url()
}

func (m *Faker) MacAddr() string {
	return internet.MacAddr()
}
