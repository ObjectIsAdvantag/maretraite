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
	"fmt"
	"errors"
)

// La structure InfosDepartLegal regroupe les données légales relatives au départ à la retraite.
// Cette structure est calculée à partir d'une date de naissance, voir la fonction CalculerDepartLegal
type InfosDepartEnRetraite struct {
	TrimestresTauxPlein int             `json:"nbTrimestres"`             // nombre de trimestres afin de disposer de sa retraite à taux plein
	AgeDépartMin        AnneesMoisJours `json:"ageDepartMinimum"`         // âge à partir duquel il est possible de percevoir une retraite, mais elle ne sera pas à taux plein si le nombre de trimestres cotisés n'est pas celui requis
	DateDépartMin       time.Time       `json:"dateDepartMinimum"`        // date à partir de laquelle il est possible de percevoir une retraite
	AgeTauxPleinAuto    AnneesMoisJours `json:"ageTauxPleinAutomatique"`  // âge à partir duquel il est possible de percevoir une retraite à taux plein même si le nombre de trimestres cotisés n'est pas suffisant
	DateTauxPleinAuto   time.Time       `json:"dateTauxPleinAutomatique"` // date à partir de laquelle il est possible de percevoir une retraite à taux plein même si le nombre de trimestres cotisés n'est pas suffisant
	AgeDépartExigible   AnneesMoisJours `json:"ageDepartExigible"`        // âge à partir duquel l'employeur peut exiger un départ en retraite de l'employé
	DateDépartExigible  time.Time       `json:"dateDepartExigible"`       // date à partir de laquelle l'employeur peut exiger un départ en retraite de l'employé
}

type TauxRetraite string
const (
	TAUX_PLEIN   TauxRetraite = "Taux Plein"
	TAUX_PARTIEL TauxRetraite = "Taux Partiel"
)

type CalculDépart struct {
	Taux               TauxRetraite    // résultat du calcul du taux pour la date de départ
	Age                AnneesMoisJours // âge au moment du départ
	Date               time.Time       // date en JJ/MM/AAAA
	TrimestresCotisés  int             // nombre de trimestres cotisés au final entre la date de début d'activité et le moment du départ
	TrimestresRestants int             // nombre de trimestres restants à cotiser
}

const ANNEE_NAISSANCE_MIN = 1900
var ANNEE_NAISSANCE_MAX = time.Now().Year()
const ANNEE_MIN = 0
const ANNEE_MAX = 2100  // arbitrary limit, because we need one

var ErrDateLimites = errors.New("la date n'est pas entre le 01/01/1900 et aujourd'hui")

// Calcule les informations de départ légal à la retraite, à partir d'une date de naissance au format JJ/MM/AAAA
// En cas d'erreur, retourne l'erreur ainsi qu'une structure InfosDepartLegal vide
func CalculerDépartLégal(dateJJMMAAAA string) (InfosDepartEnRetraite, error) {
	date, err := parseDateNaissance(dateJJMMAAAA)
	if err != nil {
		return InfosDepartEnRetraite{}, err
	}

	return calculerDépartLégalInterne(date)
}

func calculerDépartLégalInterne(dateNaissance time.Time) (InfosDepartEnRetraite, error) {
	trimestres, err := RechercherTrimestre(dateNaissance)
	if err != nil {
		return InfosDepartEnRetraite{}, err
	}

	ageLegal, err := RechercherAgeLegal(dateNaissance)
	if err != nil {
		return InfosDepartEnRetraite{}, err
	}

	return InfosDepartEnRetraite{
		TrimestresTauxPlein: trimestres,
		AgeDépartMin:        ageLegal.AgeDepartMin,
		DateDépartMin:		 DatePlusAge(dateNaissance, ageLegal.AgeDepartMin),
		AgeTauxPleinAuto:    ageLegal.AgeTauxPleinAuto,
		DateTauxPleinAuto:	 DatePlusAge(dateNaissance, ageLegal.AgeTauxPleinAuto),
		AgeDépartExigible:   ageLegal.AgeDepartExigible,
		DateDépartExigible:	 DatePlusAge(dateNaissance, ageLegal.AgeDepartExigible),
	}, nil
}


// Calcule les conditions de départ en retraite à taux plein, à partir de l'année de naissance, et du nombre de trimestres acquis à la fin d'année précisée dans le relevé de situation individuelle.
// La date de naissance est au format DD/MM/AAAA.
// Le nombre de trimestres tels qui apparaissent sur le relevé de situtation individuelle.
// La date du relevé de situation individuelle (ligne "Salarié du régime général de sécurité sociale (CNAV) - ANNEE)
// Ce calcul est théorique dans la mesure où il repose sur plusieurs hypothèses :
//    - que vous soyez en pleine activité depuis votre dernier relevé et jusqu'à votre départ en retraite
//    - que vous ayez atteint l'âge minimal de départ en retraite pour votre année de naissance (voir fonction
// Par ailleurs, vous pourrez partir à taux plein si l'âge retourné par cette fonction est supérieur à l'âge automatique à taux plein défini pour votre année de naissance
func CalculerDépartTauxPleinThéorique(dateJJMMAAAA string, trimestresAcquis int, annéeDuRelevé int) (CalculDépart, error) {
	dateNaissance, err := parseDateNaissance(dateJJMMAAAA)
	if err != nil {
		return CalculDépart{}, err
	}

	infosDepart, err := calculerDépartLégalInterne(dateNaissance)
	if err != nil {
		return CalculDépart{}, err
	}

	// La formule de calcul du départ à taux plein :
	// DateDuRelevé (revue au 1er janvier de l'année suivante) + 4 * (TrimestresTauxPlein - TrimestresAcquisAlaDateDuRelevé)
	trimestresRestants := infosDepart.TrimestresTauxPlein - trimestresAcquis
	anneesRestantes := int(trimestresRestants / 4)
	moisRestants := 3*trimestresRestants - 12*anneesRestantes
	dateRelevé := time.Date(annéeDuRelevé+1, 1, 1, 0, 0, 0, 0, time.UTC)
	dateRetraiteTauxPlein := dateRelevé.AddDate(anneesRestantes, moisRestants, 0)

	age, err := CalculerDurée(dateNaissance, dateRetraiteTauxPlein)
	if err != nil {
		return CalculDépart{}, err
	}

	return CalculDépart{
		Taux:               TAUX_PLEIN,
		Age:                age,
		Date:               dateRetraiteTauxPlein,
		TrimestresCotisés:  infosDepart.TrimestresTauxPlein,
		TrimestresRestants: trimestresRestants,
	}, nil
}


func parseDateNaissance(dateJJMMAAA string) (time.Time, error) {
	dateNaissance, err := StringToTime(dateJJMMAAA)
	if err != nil {
		return time.Time{}, err
	}

	// Vérifier les bornes de la date de naissance
	min, _ := time.ParseInLocation(JJMMAAADateFormat, fmt.Sprintf("01/01/%d", ANNEE_NAISSANCE_MIN), time.UTC)
	if dateNaissance.Before(min) || dateNaissance.After(time.Now()) {
		return time.Time{}, ErrDateLimites
	}

	return dateNaissance, nil
}

