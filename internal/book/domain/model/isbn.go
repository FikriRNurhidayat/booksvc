package book_model

import "regexp"

const ISBNPattern = `^(?=(?:\D*\d){10}(?:(?:\D*\d){3})?$)[\d-]+$`

var isbnCheck, _ = regexp.Compile(ISBNPattern)

type ISBN string

func (i ISBN) IsValid() bool {
	return isbnCheck.MatchString(i.String())
}

func (i ISBN) String() string {
	return string(i)
}
