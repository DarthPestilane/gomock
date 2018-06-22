package gomock

import (
	"strconv"
	"testing"
)

func TestMocker_Name(t *testing.T) {
	t.Log(NewMock().Name())
	t.Log(NewMock().Name("midname"))
}

func TestMocker_IntBetween(t *testing.T) {
	t.Log(NewMock().IntBetween(-10, 100))
}

func TestMocker_Int(t *testing.T) {
	t.Log(NewMock().Int())
}

func TestMocker_NameMale(t *testing.T) {
	t.Log(NewMock().NameMale())
	t.Log(NewMock().NameMale("male"))
}

func TestMocker_NameFemale(t *testing.T) {
	t.Log(NewMock().NameFemale())
	t.Log(NewMock().NameFemale("female"))
}

func TestMocker_Natural(t *testing.T) {
	t.Log(NewMock().Natural())
}

func TestMocker_NaturalBetween(t *testing.T) {
	t.Log(NewMock().NaturalBetween(100, 200))
}

func TestMocker_Float(t *testing.T) {
	t.Logf("%f", NewMock().Float())
}

func TestMocker_FloatBetween(t *testing.T) {
	t.Logf("%f", NewMock().FloatBetween(-100, 100))
}

func TestMocker_FloatWithDecimal(t *testing.T) {
	t.Logf("%s", strconv.FormatFloat(NewMock().FloatWithDecimal(2, 8), 'f', -1, 64))
}

func TestMocker_FloatBetweenWithDecimal(t *testing.T) {
	t.Logf("%f", NewMock().FloatBetweenWithDecimal(-100, 100, 2, 6))
}

func TestMocker_Char(t *testing.T) {
	t.Log(NewMock().Char())
}

func TestMocker_CharFrom(t *testing.T) {
	t.Log(NewMock().CharFrom("abcd"))
}

func TestMocker_String(t *testing.T) {
	t.Log(NewMock().String())
}

func TestMocker_StringFrom(t *testing.T) {
	t.Log(NewMock().StringFrom("HelloWorld"))
}

func TestMocker_StringWithLength(t *testing.T) {
	t.Log(NewMock().StringWithLength(10, 15))
}

func TestMocker_StringWithLengthFrom(t *testing.T) {
	t.Log(NewMock().StringWithLengthFrom(10, 15, "HelloWorld"))
}
