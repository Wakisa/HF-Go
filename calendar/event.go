package calendar

import (
	"errors"       // for creatting error values
	"unicode/utf8" // Used to count the number of runes in a string
)

type Event struct {
	title string
	Date
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 { // if the title as more than 30 characters, return an error
		return errors.New("invalid title or too long")
	}

	e.title = title
	return nil
}
