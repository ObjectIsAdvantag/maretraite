// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

package depart

import (
	"testing"
)


func TestTrimestresPourMoi(t *testing.T) {
	input := "24/12/1971"
	birthDate, err := StringToTime(input)
	if err != nil {
		t.Errorf("parse error in periode check")
	}

	expected := 171
	nombre, err := RechercherTrimestre(birthDate)
	if err != nil {
		t.Errorf("get error in trimestre check")
	}
	if nombre != expected {
		t.Errorf("mauvais nombre de trimestres: %d au lieu de %d", nombre, expected)
	}
}

