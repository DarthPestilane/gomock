package basic

import (
	"fmt"
	"github.com/DarthPestilane/gofake/helper"
	"math"
	"strconv"
	"strings"
)

var (
	Nums    = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	NumsStr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	Lower   = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	Upper   = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	Symbol  = []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "[", "]"}
)

func Bool() bool {
	array := []bool{true, false}
	var res bool
	helper.RandomSlice(array, &res)
	return res
}

type OptionNumber struct {
	Min        int
	Max        int
	MaxDecimal int
	MinDecimal int
}

func Int(option ...OptionNumber) int {
	opt := OptionNumber{
		Min: math.MinInt32,
		Max: math.MaxInt32,
	}
	if len(option) > 0 {
		opt = option[0]
	}
	return helper.Random(1+int(math.Abs(float64(opt.Max)-float64(opt.Min)))) + int(math.Min(float64(opt.Max), float64(opt.Min)))
}

func Natural(option ...OptionNumber) int {
	opt := OptionNumber{
		Min: 0,
		Max: math.MaxInt32,
	}
	if len(option) > 0 {
		opt = option[0]
		if opt.Min < 0 {
			opt.Min = 0
		}
	}
	return Int(opt)
}

func Float(option ...OptionNumber) float64 {
	var opt OptionNumber
	if len(option) > 0 {
		opt = option[0]
		if opt.MinDecimal <= 0 {
			opt.MinDecimal = 1
		}
		if opt.MaxDecimal <= 0 {
			opt.MaxDecimal = 10
		}
	} else {
		opt.Min = math.MinInt32
		opt.Max = math.MaxInt32
		opt.MinDecimal = 1
		opt.MaxDecimal = 10
	}
	randInt := Int(opt)
	decimalLen := Natural(OptionNumber{Min: opt.MinDecimal, Max: opt.MaxDecimal})
	var randDecimal string
	numsGt0 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < decimalLen; i++ {
		var char int
		if i+1 == decimalLen { // last
			helper.RandomSlice(numsGt0, &char)
		} else {
			helper.RandomSlice(Nums, &char)
		}
		randDecimal += fmt.Sprintf("%d", char)
	}
	flt, err := strconv.ParseFloat(fmt.Sprintf("%d.%s", randInt, randDecimal), 10)
	if err != nil {
		panic(fmt.Errorf("parse string to float failed: %v", err))
	}
	return flt
}

// Char gives only one random charactor
func Char(fromChars ...string) string {
	var all []string
	if len(fromChars) == 0 {
		all = append(all, NumsStr...)
		all = append(all, Lower...)
		all = append(all, Upper...)
		all = append(all, Symbol...)
	} else {
		all = fromChars
	}
	var char string
	helper.RandomSlice(strings.Split(strings.Join(all, ""), ""), &char)
	return char
}

type OptionString struct {
	MaxLen    int
	MinLen    int
	FromChars string // pick up strings form here
}

// String gives a random string
func String(option ...OptionString) string {
	var opt OptionString
	if len(option) > 0 {
		opt = option[0]
	}
	if opt.MinLen <= 0 {
		opt.MinLen = 10
	}
	if opt.MaxLen <= 0 {
		opt.MaxLen = 20
	}
	if opt.FromChars == "" {
		opt.FromChars = strings.Join(NumsStr, "") + strings.Join(Lower, "") + strings.Join(Upper, "") + strings.Join(Symbol, "")
	}
	length := Natural(OptionNumber{Min: opt.MinLen, Max: opt.MaxLen})
	var str string
	for i := 0; i < length; i++ {
		str += Char(opt.FromChars)
	}
	return str
}
