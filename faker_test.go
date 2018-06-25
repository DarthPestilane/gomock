package gofake

import (
	"strconv"
	"testing"
)

func TestFaker_Name(t *testing.T) {
	t.Log(NewFaker().Name())
	t.Log(NewFaker().Name("midname"))
}

func TestFaker_IntBetween(t *testing.T) {
	t.Log(NewFaker().IntBetween(-10, 100))
}

func TestFaker_Int(t *testing.T) {
	t.Log(NewFaker().Int())
}

func TestFaker_NameMale(t *testing.T) {
	t.Log(NewFaker().NameMale())
	t.Log(NewFaker().NameMale("male"))
}

func TestFaker_NameFemale(t *testing.T) {
	t.Log(NewFaker().NameFemale())
	t.Log(NewFaker().NameFemale("female"))
}

func TestFaker_Natural(t *testing.T) {
	t.Log(NewFaker().Natural())
}

func TestFaker_NaturalBetween(t *testing.T) {
	t.Log(NewFaker().NaturalBetween(100, 200))
}

func TestFaker_Float(t *testing.T) {
	t.Logf("%f", NewFaker().Float())
}

func TestFaker_FloatBetween(t *testing.T) {
	t.Logf("%f", NewFaker().FloatBetween(-100, 100))
}

func TestFaker_FloatWithDecimal(t *testing.T) {
	t.Logf("%s", strconv.FormatFloat(NewFaker().FloatWithDecimal(2, 8), 'f', -1, 64))
}

func TestFaker_FloatBetweenWithDecimal(t *testing.T) {
	t.Logf("%f", NewFaker().FloatBetweenWithDecimal(-100, 100, 2, 6))
}

func TestFaker_Char(t *testing.T) {
	t.Log(NewFaker().Char())
}

func TestFaker_CharFrom(t *testing.T) {
	t.Log(NewFaker().CharFrom("abcd"))
}

func TestFaker_String(t *testing.T) {
	t.Log(NewFaker().String())
}

func TestFaker_StringFrom(t *testing.T) {
	t.Log(NewFaker().StringFrom("HelloWorld"))
}

func TestFaker_StringWithLength(t *testing.T) {
	t.Log(NewFaker().StringWithLength(10, 15))
}

func TestFaker_StringWithLengthFrom(t *testing.T) {
	t.Log(NewFaker().StringWithLengthFrom(10, 15, "HelloWorld"))
}
