// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

package depart

import (
	"testing"
	"log"
)


func TestHorsLimitesInf(t *testing.T) {
	naissance := "19/01/1890"
	if _, err := CalculerInfosLégales(naissance); err != ErrDateLimites {
		t.Errorf("parse error, under limit for: %s, err:", naissance, err)
	}
}

func TestHorsLimitesInfSup(t *testing.T) {
	naissance := "19/01/2017"
	if _, err := CalculerInfosLégales(naissance); err != ErrDateLimites {
		t.Errorf("parse error, over limit for: %s, err:", naissance, err)
	}
}

func TestCalculerInfosDeDepartRetaitePourMoi(t *testing.T) {
	naissance := "24/12/1971"
	res, err := CalculerInfosLégales(naissance)
	if err != nil {
		t.Errorf("impossible de calculer les informations de départ en retraite, error: %s", err)
		return
	}

	expTrimestres := 171
	if res.TrimestresTauxPlein != expTrimestres {
		t.Errorf("erreur calcul des trimestres: %d contre %d attendus", res.TrimestresTauxPlein, expTrimestres)
	}

	expAge := AnneesMoisJours{Annees: 62, Mois: 0}
	if res.AgeDépartMin != expAge {
		t.Errorf("erreur calcul de l'age de départ min: %v contre %v attendu", res.AgeDépartMin, expAge)
	}

	expDate, _ := StringToTime("24/12/2033")
	if res.DateDépartMin != expDate {
		t.Errorf("erreur calcul de la date de départ min: %v contre %v attendu", res.DateDépartMin, expDate)
	}

	expAge = AnneesMoisJours{Annees: 67, Mois: 0}
	if res.AgeTauxPleinAuto != expAge {
		t.Errorf("erreur calcul de l'age du taux plein automatique: %v contre %v attendu", res.AgeTauxPleinAuto, expAge)
	}

	expDate, _ = StringToTime("24/12/2038")
	if res.DateTauxPleinAuto != expDate {
		t.Errorf("erreur calcul de la date de départ min: %v contre %v attendu", res.DateTauxPleinAuto, expDate)
	}

	expAge = AnneesMoisJours{Annees: 70, Mois: 0}
	if res.AgeDépartExigible != expAge {
		t.Errorf("erreur calcul de l'age du taux plein automatique: %v contre %v attendu", res.AgeDépartExigible, expAge)
	}

	expDate, _ = StringToTime("24/12/2041")
	if res.DateDépartExigible != expDate {
		t.Errorf("erreur calcul de la date de départ min: %v contre %v attendu", res.DateDépartExigible, expDate)
	}
}

func TestCalculerDonnesRetraiteTauxPleinPourMoi(t *testing.T) {
	naissance := "24/12/1971"
	res, err := CalculerDépartTauxPlein(naissance, 87, 2014)
	if err != nil {
		t.Errorf("Impossible de calculer les conditions de départ à taux plein, err: %s", err)
		return
	}

	expAgeTauxPlein := AnneesMoisJours{Annees: 64, Mois: 0, Jours: 8}
	if res.Age != expAgeTauxPlein {
		t.Errorf("erreur calcul de l'age de départ effectif à taux plein min: %v contre %v attendu", res.Age, expAgeTauxPlein)
	}

	if res.Date.IsZero() {
		t.Errorf("erreur calcul de l'age de départ effectif à taux plein min: date vide")
	}

	expDateTauxPlein := "01/01/2036"
	comp := TimeToString(res.Date)
	if comp != expDateTauxPlein {
		t.Errorf("erreur calcul de la date de départ effectif à taux plein min: %s contre %s attendu", comp, expDateTauxPlein)
	}
}

func TestCalculerInfosDeDepartRetaitePourValérie(t *testing.T) {
	naissance := "04/07/1974"
	res, err := CalculerInfosLégales(naissance)
	if err != nil {
		t.Errorf("impossible de calculer les informations de départ en retraite, error: %s", err)
		return
	}

	expTrimestres := 172
	if res.TrimestresTauxPlein != expTrimestres {
		t.Errorf("erreur calcul des trimestres: %d contre %d attendus", res.TrimestresTauxPlein, expTrimestres)
	}

	expAgeDepartMin := AnneesMoisJours{Annees: 62, Mois: 0}
	if res.AgeDépartMin != expAgeDepartMin {
		t.Errorf("erreur calcul de l'age de départ min: %v contre %v attendu", res.AgeDépartMin, expAgeDepartMin)
	}

	expDateDepartMin, _ := StringToTime("04/07/2036")
	if res.DateDépartMin != expDateDepartMin {
		t.Errorf("erreur calcul de la date de départ min: %v contre %v attendu", res.DateDépartMin, expDateDepartMin)
	}

	expAgeTauxPlein := AnneesMoisJours{Annees: 67, Mois: 0}
	if res.AgeTauxPleinAuto != expAgeTauxPlein {
		t.Errorf("erreur calcul de l'age du taux plein automatique: %v contre %v attendu", res.AgeTauxPleinAuto, expAgeDepartMin)
	}
}

func TestCalculerDonnesRetraiteTauxPleinPourValérie(t *testing.T) {
	naissance := "04/07/1974"
	res, err := CalculerDépartTauxPlein(naissance, 69, 2014)
	if err != nil {
		t.Errorf("Impossible de calculer les conditions de départ à taux plein, err: %s", err)
		return
	}

	expAgeTauxPlein := AnneesMoisJours{Annees: 66, Mois: 2, Jours: 27}
	if res.Age != expAgeTauxPlein {
		t.Errorf("erreur calcul de l'age de départ effectif à taux plein min: %v contre %v attendu", res.Age, expAgeTauxPlein)
	}

	if res.Date.IsZero() {
		t.Errorf("erreur calcul de l'age de départ effectif à taux plein min: date vide")
	}

	expDateTauxPlein := "01/10/2040"
	comp := TimeToString(res.Date)
	if comp != expDateTauxPlein {
		t.Errorf("erreur calcul de la date de départ effectif à taux plein min: %s contre %s attendu", comp, expDateTauxPlein)
	}
}

func TestCalculerConditionsDepartTropTot(t *testing.T) {
	dateDepart := "24/12/2032"
	dateNaissance := "24/12/1971"
	trimestresRelevé := 87
	dateRelevé := 2014
	_, impossible, err := CalculerConditionsDepart(dateNaissance, trimestresRelevé, dateRelevé, dateDepart)
	if err != nil {
		t.Errorf("Impossible de calculer les conditions de départ, should not happen")
		return
	}

	if (impossible == DepartImpossible{}) {
		t.Errorf("Erreur, le départ ne devrait pas être possible")
		return
	}

	if impossible.Code != TROP_TOT {
		t.Errorf("Raison de départ impossible incorrecte, code attendu %s, vs obtenu %s", TROP_TOT, impossible.Code)
		return
	}

	log.Printf("Si vous êtes né le %s, et que vous avez cotisé %d trimestres à fin %d", dateNaissance, trimestresRelevé, dateRelevé)
	log.Printf("vous ne pourrez pas partir en retraite le %s, motif %s", dateDepart, impossible.Motif)
}

func TestCalculerConditionsDepartPasAssezDeTrimestres(t *testing.T) {
	dateDepart := "24/12/2033"
	dateNaissance := "24/12/1971"
	trimestresRelevé := 87
	dateRelevé := 2014
	log.Printf("Si vous êtes né le %s, et que vous avez cotisé %d trimestres à fin %d", dateNaissance, trimestresRelevé, dateRelevé)
	_, impossible, err := CalculerConditionsDepart(dateNaissance, trimestresRelevé, dateRelevé, dateDepart)
	if err != nil {
		t.Errorf("Impossible de calculer les conditions de départ, should not happen")
		return
	}

	if (impossible == DepartImpossible{}) {
		t.Errorf("Erreur, le départ ne devrait pas être possible")
		return
	}

	if impossible.Code != TRIMESTRES_INSUFFISANTS {
		t.Errorf("Raison de départ impossible incorrecte, code attendu %s, vs obtenu %s", TROP_TOT, impossible.Code)
		return
	}

	log.Printf("vous ne pourrez pas partir en retraite le %s, motif %s", dateDepart, impossible.Motif)
}
