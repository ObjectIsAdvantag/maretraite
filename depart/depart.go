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
	TrimestresMinimum	int				`json:"nbTrimestresMinimum"`      // nombre de trimestres cotisés au minimum afin de pouvoir partir en retraitei
	TrimestresTauxPlein int             `json:"nbTrimestresTauxPlein"`    // nombre de trimestres afin de disposer de sa retraite à taux plein
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

const NB_TRIMESTRES_MIN_SOUS_TAUX_PLEIN = 20

var ErrDateLimites = errors.New("la date n'est pas entre le 01/01/1900 et aujourd'hui")
var ErrDateRelevé = errors.New("Le date du relevé est incorrecte, ou date de plus de 20 ans")

// Calcule les informations de départ légal à la retraite, à partir d'une date de naissance au format JJ/MM/AAAA
// En cas d'erreur, retourne l'erreur ainsi qu'une structure InfosDepartLegal vide
func CalculerInfosLégales(dateJJMMAAAA string) (InfosDepartEnRetraite, error) {
	date, err := parseDateNaissance(dateJJMMAAAA)
	if err != nil {
		return InfosDepartEnRetraite{}, err
	}

	return calculerDépartLégalInterne(date)
}

func calculerDépartLégalInterne(dateNaissance time.Time) (InfosDepartEnRetraite, error) {
	trimestres, err := RechercherTauxPlein(dateNaissance)
	if err != nil {
		return InfosDepartEnRetraite{}, err
	}

	ageLegal, err := RechercherAgeLegal(dateNaissance)
	if err != nil {
		return InfosDepartEnRetraite{}, err
	}

	return InfosDepartEnRetraite{
		TrimestresMinimum:	 trimestres - NB_TRIMESTRES_MIN_SOUS_TAUX_PLEIN,
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
func CalculerDépartTauxPlein(dateJJMMAAAA string, trimestresAcquis int, annéeDuRelevé int) (CalculDépart, error) {
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
	dateRelevé, err := VérifierRelevé(annéeDuRelevé)
	if err != nil {
		return CalculDépart{}, err
	}
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

// Vérifie la date du relevé, et retourne la date à partir de laquelle les nouvelles cotisations débuteront
func VérifierRelevé(annéeDuRelevé int) (time.Time, error) {
	// Le relevé doit daté de moins de 20 ans
	if annéeDuRelevé < time.Now().Year()-20 || annéeDuRelevé > time.Now().Year() {
		return time.Time{}, ErrDateRelevé
	}

	return time.Date(annéeDuRelevé+1, 1, 1, 0, 0, 0, 0, time.UTC), nil

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


type DepartImpossible struct {
	Motif string
}

// Calcule les conditions d'un départ en retraite à la date spécifiée
// La date de naissance est au format DD/MM/AAAA.
// Le nombre de trimestres tels qui apparaissent sur le relevé de situtation individuelle.
// La date du relevé de situation individuelle (ligne "Salarié du régime général de sécurité sociale (CNAV) - ANNEE)
// Ce calcul est théorique dans la mesure où il suppose que vous restiez en pleine activité depuis votre dernier relevé et jusqu'à votre départ en retraite
//
// Les résultats sont :
//    - soit le départ n'est pas possible à la date indiquée, avec le motif (manque de trimestre ou age insuffisant)
//    - si le départ est possible, les conditions sont détaillées (le taux de la pension)
func CalculerConditionsDepart(naissanceJJMMAAAA string, trimestresAcquis int, annéeDuRelevé int, departJJMMAAA string) (CalculDépart, DepartImpossible, error) {
	dateNaissance, err := parseDateNaissance(naissanceJJMMAAAA)
	if err != nil {
		return CalculDépart{}, DepartImpossible{}, err
	}
	dateDepart, err := StringToTime(departJJMMAAA)
	if err != nil {
		return CalculDépart{}, DepartImpossible{}, err
	}

	return calculerConditionsDepartInternal(dateNaissance, trimestresAcquis, annéeDuRelevé, dateDepart)
}

func calculerConditionsDepartInternal(dateNaissance time.Time, trimestresAcquis int, annéeDuRelevé int, dateDepart time.Time) (CalculDépart, DepartImpossible, error) {

	infosDepart, err := calculerDépartLégalInterne(dateNaissance)
	if err != nil {
		return CalculDépart{}, DepartImpossible{}, err
	}

	// A-t-on atteint l'age de départ en retraite
	if dateDepart.Before(infosDepart.DateDépartMin) {
		motif := fmt.Sprintf("Vous n'avez pas atteint %s, l'âge légal de départ en retraite pour votre année de naissance", infosDepart.AgeDépartMin.AgeEnAnnees())
		return CalculDépart{}, DepartImpossible{Motif:motif}, nil
	}

	// A-t-on cotisé suffisamment de trimestre
	dateReleve, err := VérifierRelevé(annéeDuRelevé)
	if err != nil {
		return CalculDépart{}, DepartImpossible{}, err
	}
	trimestresComplementaires, err := NombreDeTrimestresEntre(dateReleve, dateDepart)
	if err != nil {
		return CalculDépart{}, DepartImpossible{}, err
	}
	totalTrimestres := trimestresAcquis + trimestresComplementaires

	if totalTrimestres < infosDepart.TrimestresMinimum {
		motif := fmt.Sprintf("Vous n'aurez pas cotisé suffisamment de trimestres le %s, %d trimestres requis contre %d trimestres cotisés si vous n'avez pas d'interruption d'activité", TimeToString(dateDepart),  infosDepart.TrimestresMinimum, totalTrimestres)
		return CalculDépart{}, DepartImpossible{Motif:motif}, nil
	}

	// Calcul des conditions de départ
	// TODO + Vérifier qu'il y a bien un nb mini de trimestres à cotiser

	return CalculDépart{}, DepartImpossible{}, nil




	/*
		// La formule de calcul du départ à taux plein :
		// DateDuRelevé (revue au 1er janvier de l'année suivante) + 4 * (TrimestresTauxPlein - TrimestresAcquisAlaDateDuRelevé)
		trimestresRestants := infosDepart.TrimestresTauxPlein - trimestresAcquis
		anneesRestantes := int(trimestresRestants / 4)
		moisRestants := 3*trimestresRestants - 12*anneesRestantes
		dateRelevé, err := calculerDateReleve(annéeDuRelevé)
		if err != nil {
			return CalculDépart{}, err
		}
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
		*/
}
