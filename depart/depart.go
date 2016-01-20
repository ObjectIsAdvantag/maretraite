// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Age légal de départ à la retraite (as of January 2016)
//
// Ce service vous permet de connaitre votre âge légal de départ à la retraite et vous précise :
// - votre âge légal de départ à la retraite
// - le point de départ possible de votre retraite
// - la durée de cotisation pour obtenir une retraite à taux plein
// - l'âge et la date auxquels vous obtiendrez automatiquement une retraite à taux plein
package depart

import (
	"time"
)

// La structure InfosDepartLegal regroupe les données légales relatives au départ à la retraite.
// Cette structure est calculée à partir d'une date de naissance, voir la fonction CalculerDepartLegal
type InfosDepartEnRetraite struct {
	TrimestresTauxPlein int					`json:"nbTrimestres"`            // nombre de trimestres afin de disposer de sa retraite à taux plein
	AgeDépartMin        AnneesMoisJours        `json:"ageDepartMinimum"`        // âge à partir duquel il est possible de percevoir une retraite, mais elle ne sera pas à taux plein si le nombre de trimestres cotisés n'est pas celui requis
																				//DateDepartMin		string				`json:"dateDepartMinimum"`// date à partir de laquelle il est possible de percevoir une retraite
	AgeTauxPleinAuto    AnneesMoisJours        `json:"ageTauxPleinAutomatique"` // âge à partir duquel il est possible de percevoir une retraite à taux plein même si le nombre de trimestres cotisés n'est pas suffisant
																				//DateDepartAuto		string				`json:"dateDepartAutomatique"`// date à partir de laquelle il est possible de percevoir une retraite à taux plein même si le nombre de trimestres cotisés n'est pas suffisant
	AgeDépartExigible   AnneesMoisJours        `json:"ageDepartExigible"`       // âge à partir duquel l'employeur peut exiger un départ en retraite de l'employé
}

type TauxRetraite string
const (
	TAUX_PLEIN TauxRetraite = "Taux Plein"
	TAUX_PARTIEL TauxRetraite = "Taux Partiel"
)


type CalculDepart struct {
	Taux              TauxRetraite    // résultat du calcul du taux pour la date de départ
	Age               AnneesMoisJours // âge au moment du départ
	Date              time.Time       // date en JJ/MM/AAAA
	TrimestresCotisés int             // nombre de trimestres cotisés au final entre la date de début d'activité et le moment du départ
}

const ANNEE_MIN = 1900
const ANNEE_MAX = 2100

// Calcule les informations de départ légal à la retraite, à partir d'une date de naissance au format JJ/MM/AAAA
// En cas d'erreur, retourne l'erreur ainsi qu'une structure InfosDepartLegal vide
func CalculerDepartLegal(dateDeNaissance string) (InfosDepartEnRetraite, error) {
	date, err := parseDate(dateDeNaissance)
	if err != nil {
		return InfosDepartEnRetraite{}, err
	}

	return calculerDepartLegalInternal(date)
}


func calculerDepartLegalInternal(date time.Time) (InfosDepartEnRetraite, error) {
	trimestres, err := RechercherTrimestre(date)
	if err != nil {
		return InfosDepartEnRetraite{}, err
	}

	ageLegal, err := RechercherAgeLegal(date)
	if err != nil {
		return InfosDepartEnRetraite{}, err
	}

	return InfosDepartEnRetraite{
		TrimestresTauxPlein:trimestres,
		AgeDépartMin:ageLegal.AgeDepartMin,
		AgeTauxPleinAuto:ageLegal.AgeTauxPleinAuto,
		AgeDépartExigible:ageLegal.AgeDepartExigible,
	}, nil
}

// Calcule les conditions de départ en retraite à taux plein, à partir de l'année de naissance, et du nombre de trimestres acquis à la fin d'année précisée dans le relevé de situation individuelle.
// La date de naissance est au format DD/MM/AAAA.
// Le nombre de trimestres tels qui apparaissent à la date du relevé
// La dernière année   du relevé de situation individuelle
func CalculerDepartTauxPlein(dateSaisie string, trimestresAcquis int, annéeDuRelevé int) (CalculDepart, error) {
	dateNaissance, err := parseDate(dateSaisie)
	if err != nil {
		return CalculDepart{}, err
	}

	infosDepart, err := calculerDepartLegalInternal(dateNaissance)
	if err != nil {
		return CalculDepart{}, err
	}

	// La formule de calcul est du départ à taux plein
	// DateDuRelevé + (TrimestresTauxPlein - trimestresAcquis)*4
	trimestresAcotiser := infosDepart.TrimestresTauxPlein - trimestresAcquis
	anneesAcotiser := int(trimestresAcotiser / 4)
	moisAcotiser := trimestresAcotiser*3 - 12*anneesAcotiser
	dateRelevé := time.Date(annéeDuRelevé + 1, 1, 1, 0, 0, 0, 0, time.UTC)
	dateRetraite := dateRelevé.AddDate(anneesAcotiser,moisAcotiser,0)

	age, err := CalculerAge(dateNaissance, dateRetraite)
	if err != nil {
		return CalculDepart{}, err
	}

	return CalculDepart{
		Taux:TAUX_PLEIN,
		Age: age,
		Date: dateRetraite,
		TrimestresCotisés:infosDepart.TrimestresTauxPlein,
	}, nil
}
