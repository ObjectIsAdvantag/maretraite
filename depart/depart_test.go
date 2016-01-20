// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

package depart

import (
	"testing"
)


func TestCalculerInfosDeDepartRetaitePourMoi(t *testing.T) {
	naissance := "24/12/1971"
	res, err := CalculerDepartLegal(naissance)
	if err != nil {
		t.Errorf("impossible de calculer les informations de départ en retraite, error: %s", err)
		return
	}

	expTrimestres := 171
	if res.TrimestresTauxPlein != expTrimestres {
		t.Errorf("erreur calcul des trimestres: %d contre %d attendus", res.TrimestresTauxPlein, expTrimestres)
	}

	expAgeDepartMin := AnneesMoisJours{Annees:62, Mois:0}
	if res.AgeDépartMin != expAgeDepartMin {
		t.Errorf("erreur calcul de l'age de départ min: %v contre %v attendu", res.AgeDépartMin, expAgeDepartMin)
	}

	expAgeTauxPlein := AnneesMoisJours{Annees:67, Mois:0}
	if res.AgeTauxPleinAuto != expAgeTauxPlein {
		t.Errorf("erreur calcul de l'age du taux plein automatique: %v contre %v attendu", res.AgeTauxPleinAuto, expAgeDepartMin)
	}
}

func TestCalculerDonnesRetraiteTauxPleinPourMoi(t *testing.T) {
	naissance := "24/12/1971"
	res, err := CalculerDepartTauxPlein (naissance, 87, 2014)
	if err != nil {
		t.Errorf("Impossible de calculer les conditions de départ à taux plein, err: %s", err)
		return
	}

	expAgeTauxPlein := AnneesMoisJours{Annees:64, Mois:0}
	if res.Age != expAgeTauxPlein {
		t.Errorf("erreur calcul de l'age de départ effectif à taux plein min: %v contre %v attendu", res.Age, expAgeTauxPlein)
	}

	if res.Date.IsZero() {
		t.Errorf("erreur calcul de l'age de départ effectif à taux plein min: date vide")
	}

	expDateTauxPlein := "1/1/2036"
	comp, err := unparseDate(res.Date)
	if err != nil {
		t.Errorf("erreur calcul de la date de départ effectif à taux plein min, cannot parse date: %s", res.Date)
	}
	if comp != expDateTauxPlein {
		t.Errorf("erreur calcul de la date de départ effectif à taux plein min: %s contre %s attendu", comp, expDateTauxPlein)
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