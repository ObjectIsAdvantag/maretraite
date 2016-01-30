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
	nombre, err := RechercherTauxPlein(birthDate)
	if err != nil {
		t.Errorf("get error in trimestre check")
	}
	if nombre != expected {
		t.Errorf("mauvais nombre de trimestres: %d au lieu de %d", nombre, expected)
	}
}

func TestTrimestresPourValérie(t *testing.T) {
	input := "04/07/1974"
	birthDate, err := StringToTime(input)
	if err != nil {
		t.Errorf("parse error in periode check")
	}

	expected := 172
	nombre, err := RechercherTauxPlein(birthDate)
	if err != nil {
		t.Errorf("get error in trimestre check")
	}
	if nombre != expected {
		t.Errorf("mauvais nombre de trimestres: %d au lieu de %d", nombre, expected)
	}
}

func TestNombreDeTrimestres(t *testing.T) {
	depuis, _ := StringToTime("01/01/1994")
	jusque, _ := StringToTime("01/01/2034")
	expected := 120
	nombre, err := NombreDeTrimestresEntre(depuis, jusque)

	if err != nil {
		t.Errorf("impossible de calculer le nombre de trimestres entre le %s et le %s\n", depuis, jusque)
	}
	if nombre != expected {
		t.Errorf("mauvais nombre de trimestres: %d au lieu de %d", nombre, expected)
	}
}
