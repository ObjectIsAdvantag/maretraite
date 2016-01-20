// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

package depart

import (
	"testing"
)


func TestPeriodePourMoi(t *testing.T) {
	input := "24/12/1971"
	birthDate, err := StringToTime(input)
	if err != nil {
		t.Errorf("parse error in periode check")
	}

	periode, err := RechercherAgeLegal(birthDate)
	if err != nil {
		t.Errorf("get error in periode check")
	}

	expected := AnneesMoisJours{Annees:62, Mois:0}
	if periode.AgeDepartMin != expected {
		t.Errorf("mauvais age de départ en retraite")
	}

	expected = AnneesMoisJours{Annees:67, Mois:0}
	if (periode.AgeTauxPleinAuto != expected) {
		t.Errorf("mauvais age de départ à taux plein")
	}

	expected = AnneesMoisJours{Annees:70, Mois:0}
	if (periode.AgeDepartExigible != expected) {
		t.Errorf("mauvais age de départ forcé")
	}
}