package depart

import (
	"testing"
)


func TestParsePasDeDateDeNaissance(t *testing.T) {
	if _, err := parseDate(""); err != ErrDateVide {
		t.Errorf("empty input check error")
	}
}

func TestParseBadFormat(t *testing.T) {
	if _, err := parseDate("2/1/71"); err != ErrDateFormatInvalide {
		t.Errorf("format check error")
	}
}

func TestParseUnderLimit(t *testing.T) {
	if _, err := parseDate("19/01/1890"); err != ErrDateLimites {
		t.Errorf("format check error")
	}
}

func TestParseOverLimit(t *testing.T) {
	if _, err := parseDate("19/01/2020"); err != ErrDateLimites {
		t.Errorf("format check error")
	}
}

func TestParseCorrectFormat1900(t *testing.T) {
	if _, err := parseDate("01/01/1971"); err != nil {
		t.Errorf("format check error")
	}
}

func TestParseCorrectFormat2000(t *testing.T) {
	if _, err := parseDate("02/03/2009"); err != nil {
		t.Errorf("format check error")
	}
}

/*
func calculPourMoi(t *testing) {
	result := depart.calculer("24/12/1971")
	expected := &DepartResult{
		ageLegalMin:62,
		dateDepartMin:"01/01/2034",
		trimestres:171,
		ageLegalAutomatique:67,
		dateDepartAutomatique:"01/01/2039",
	}
	if result.ageLegalMin != 62 {
		t.Errors
	}
	if result.dateDepartMin !=  {
		t.errors
	}
	if result.

}*/