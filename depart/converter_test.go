// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

package depart

import (
	"testing"
	"time"
)


func TestParsePasDeDateDeNaissance(t *testing.T) {
	input := ""
	if _, err := parseDate(input); err != ErrDateVide {
		t.Errorf("parse error, empty input for: %s, err:", input, err)
	}
}

func TestParseBadFormat(t *testing.T) {
	input := "2/1/71"
	if _, err := parseDate(input); err != ErrDateFormatInvalide {
		t.Errorf("parse error, bad format for: %s, err:", input, err)
	}
}

func TestParseUnderLimit(t *testing.T) {
	input := "19/01/1890"
	if _, err := parseDate(input); err != ErrDateLimites {
		t.Errorf("parse error, under limit for: %s, err:", input, err)
	}
}

func TestParseOverLimit(t *testing.T) {
	input := "19/01/2020"
	if _, err := parseDate(input); err != ErrDateLimites {
		t.Errorf("parse error, over limit for: %s, err:", input, err)

	}
}

func TestParseCorrectFormat1900(t *testing.T) {
	input := "01/01/1971"
	if _, err := parseDate(input); err != nil {
		t.Errorf("parse error, correct format for: %s, err:", input, err)
	}
}

func TestParseCorrectFormat2000(t *testing.T) {
	input := "02/03/2009"
	if _, err := parseDate(input); err != nil {
		t.Errorf("parse error, correct format for: %s, err:", input, err)
	}
}

func TestParseUnparse(t *testing.T) {
	input := "24/12/1971"
	var tmpDate time.Time
	var err error
	if tmpDate, err = parseDate(input); err != nil {
		t.Errorf("parse error, parse then unparse for: %s, err: %s", input, err)
	}
	var output string
	if output, err = unparseDate(tmpDate); err != nil {
		t.Errorf("parse error, parse then unparse for: %s, err: %s", input, err)
	}
	if output != input {
		t.Errorf("parse error, string: %s parsed then unparse as: %s", input, output)
	}

}
