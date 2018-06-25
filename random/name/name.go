package name

import (
	"fmt"
	"github.com/DarthPestilane/gofake/helper"
)

var (
	LastNames = []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones",
		"Miller", "Davis", "Garcia", "Rodriguez", "Wilson",
		"Martinez", "Anderson", "Taylor", "Thomas", "Hernandez",
		"Moore", "Martin", "Jackson", "Thompson", "White",
		"Lopez", "Lee", "Gonzalez", "Harris", "Clark",
		"Lewis", "Robinson", "Walker", "Perez", "Hall",
		"Young", "Allen",
	}

	FirstNames = append(FirstNamesMale, FirstNamesFemale...)

	FirstNamesMale = []string{
		"James", "John", "Robert", "Michael", "William",
		"David", "Richard", "Charles", "Joseph", "Thomas",
		"Christopher", "Daniel", "Paul", "Mark", "Donald",
		"George", "Kenneth", "Steven", "Edward", "Brian",
		"Ronald", "Anthony", "Kevin", "Jason", "Matthew",
		"Gary", "Timothy", "Jose", "Larry", "Jeffrey",
		"Frank", "Scott", "Eric",
	}

	FirstNamesFemale = []string{
		"Mary", "Patricia", "Linda", "Barbara", "Elizabeth",
		"Jennifer", "Maria", "Susan", "Margaret", "Dorothy",
		"Lisa", "Nancy", "Karen", "Betty", "Helen",
		"Sandra", "Donna", "Carol", "Ruth", "Sharon",
		"Michelle", "Laura", "Sarah", "Kimberly", "Deborah",
		"Jessica", "Shirley", "Cynthia", "Angela", "Melissa",
		"Brenda", "Amy", "Anna",
	}
)

const (
	_ = iota
	GenderMale
	GenderFemale
)

type Option struct {
	Gender     int
	MiddleName string
}

func Name(option ...Option) string {
	var opt Option
	if len(option) > 0 {
		opt = option[0]
	}
	var firstName, lastName string
	if opt.Gender == GenderMale {
		helper.RandomSlice(FirstNamesMale, &firstName)
	} else if opt.Gender == GenderFemale {
		helper.RandomSlice(FirstNamesFemale, &firstName)
	} else {
		// random gender
		helper.RandomSlice(FirstNames, &firstName)
	}
	helper.RandomSlice(LastNames, &lastName)
	if opt.MiddleName != "" {
		return fmt.Sprintf("%s %s %s", firstName, opt.MiddleName, lastName)
	}
	// random middle name
	return fmt.Sprintf("%s %s", firstName, lastName)
}
