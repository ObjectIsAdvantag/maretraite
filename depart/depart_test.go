// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

package depart

import (
	"testing"
)


func TestCalculerDepartLegalPourMoi(t *testing.T) {
	naissance := "24/12/1971"
	res, err := CalculerDepartLegal(naissance, 0)
	if err != nil {
		t.Errorf("impossible de calculer les informations de départ en retraite, error: %s", err)
	}

	expTrimestres := 171
	if res.TrimestresRequis != expTrimestres {
		t.Errorf("erreur calcul des trimestres: %d contre %d attendus", res.TrimestresRequis, expTrimestres)
	}

	expAgeDepartMin := AgeEnAnneesMois{Annees:62, Mois:0}
	if res.AgeDepartMin != expAgeDepartMin {
		t.Errorf("erreur calcul de l'age de départ min: %v contre %v attendu", res.AgeDepartMin, expAgeDepartMin)
	}

	expAgeTauxPlein := AgeEnAnneesMois{Annees:67, Mois:0}
	if res.AgeTauxPleinAuto != expAgeTauxPlein {
		t.Errorf("erreur calcul de l'age du taux plein min: %v contre %v attendu", res.AgeTauxPleinAuto, expAgeDepartMin)
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