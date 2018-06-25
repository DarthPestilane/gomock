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

func (f *Faker) NameMale(middleName ...string) string {
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

func (f *Faker) NameFemale(middleName ...string) string {
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

func (f *Faker) Bool() bool {
	return basic.Bool()
}

func (f *Faker) Int() int {
	return basic.Int()
}

func (f *Faker) IntBetween(begin, end int) int {
	return basic.Int(basic.OptionNumber{Min: begin, Max: end})
}

func (f *Faker) Natural() int {
	return basic.Natural()
}

func (f *Faker) NaturalBetween(begin, end int) int {
	return basic.Natural(basic.OptionNumber{Max: end, Min: begin})
}

func (f *Faker) Float() float64 {
	return basic.Float()
}

func (f *Faker) FloatBetween(begin, end int) float64 {
	return basic.Float(basic.OptionNumber{Max: end, Min: begin})
}

func (f *Faker) FloatWithDecimal(minDecimalLen, maxDecimalLen int) float64 {
	return basic.Float(basic.OptionNumber{Max: math.MaxInt32, Min: math.MinInt32, MaxDecimal: maxDecimalLen, MinDecimal: minDecimalLen})
}

func (f *Faker) FloatBetweenWithDecimal(begin, end, minDecimalLen, maxDecimalLen int) float64 {
	return basic.Float(basic.OptionNumber{Max: end, Min: begin, MaxDecimal: maxDecimalLen, MinDecimal: minDecimalLen})
}

func (f *Faker) Char() string {
	return basic.Char()
}

func (f *Faker) CharFrom(str string) string {
	return basic.Char(str)
}

func (f *Faker) String() string {
	return basic.String()
}

func (f *Faker) StringFrom(str string) string {
	return basic.String(basic.OptionString{FromChars: str})
}

func (f *Faker) StringWithLength(min, max int) string {
	return basic.String(basic.OptionString{MinLen: min, MaxLen: max})
}

func (f *Faker) StringWithLengthFrom(min, max int, str string) string {
	return basic.String(basic.OptionString{MinLen: min, MaxLen: max, FromChars: str})
}

func (f *Faker) Email() string {
	return internet.Email()
}

func (f *Faker) EmailWithDomain(domain string) string {
	return internet.Email(domain)
}

func (f *Faker) DomainName() string {
	return internet.DomainName()
}

func (f *Faker) DomainWord() string {
	return internet.DomainWord()
}

func (f *Faker) Url() string {
	return internet.Url()
}

func (f *Faker) MacAddr() string {
	return internet.MacAddr()
}

func (f *Faker) IPv4() string {
	return internet.IPv4()
}

func (f *Faker) IPv6() string {
	return internet.IPv6()
}

func (f *Faker) LocalIPv4() string {
	return internet.LocalIPv4()
}
