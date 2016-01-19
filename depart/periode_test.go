// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

package depart

import (
	"testing"
	"log"
)


func TestPeriodePourMoi(t *testing.T) {
	input := "24/12/1971"
	néLe, err := parseDate(input)
	if err != nil {
		t.Errorf("parse error in periode check")
	}

	periode, err := GetPeriodeDepartEnRetraite(néLe)
	if err != nil {
		t.Errorf("get error in periode check")
	}

	log.Printf("Periode %s trouvée par né le: %v", periode, input)
}
