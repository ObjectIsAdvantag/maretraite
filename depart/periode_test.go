// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

package depart

import (
	"testing"
)


func TestPeriodePourMoi(t *testing.T) {
	input := "24/12/1971"
	birthDate, err := parseDate(input)
	if err != nil {
		t.Errorf("parse error in periode check")
	}

	periode, err := RechercherAgeLegal(birthDate)
	if err != nil {
		t.Errorf("get error in periode check")
	}

	if periode.AgeDepartLegal.EnAnnees() != 62 {
		t.Errorf("mauvais age de départ en retraite")
	}

	if (periode.AgeTauxPlein.EnAnnees() != 67) {
		t.Errorf("mauvais age de départ à taux plein")
	}

	if (periode.AgeRetraiteForcee.EnAnnees() != 70) {
		t.Errorf("mauvais age de départ forcé")
	}
}
