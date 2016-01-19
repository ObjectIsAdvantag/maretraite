// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Référentiel des dates légales de départ à la retraite en fonction des années de naissance
package depart

import (
	"time"
	"errors"
)

// Le type AgeEnAnneesMois représente un age de départ en retraite sous la forme d'un age en années et mois,
// Exemple : 62 ans et 7 mois
type AgeEnAnneesMois struct {
	Annees      int		`json:"annees,required"`		// nombre d'années
	Mois 		int		`json:"mois,omitempty"`			// nombre de mois
}

// Le type PeriodeDepartLegal représente l'age minimal de départ à la retraite.
// Ce type est utilisé pour associer une période légale à une durée de naissance.
// Exemple :  pour les personnes nées à partir du  01/01/1954, 61 ans et 7 mois pour l'âge minimum et 66 ans et 7 mois pour une retraite à taux plein

type PeriodeDepartLegal struct {
	Depuis					time.Time			// date de naissance à partir de laquelle les âges de départs en retraite s'appliquent, vide s'il n'y a pas de borne inférieure
	Jusque					time.Time			// date de naissance jusque laquelle les âges de départs en retraite s'appliquent, vide s'il n'y a pas de borne supérieure
	AgeDepartLegal			AgeEnAnneesMois		// age minimal à partir duquel on peut percevoir sa retraite
	AgeTauxPlein			AgeEnAnneesMois		// age à partir duquel on percevra sa retraite à taux plein même si le nombre de trimestre cotisés n'est pas suffisant
	AgeRetraiteForcee		AgeEnAnneesMois		// age à partir duquel un employeur peut exiger qu'un salarié prenne sa retraite même sans son consentement
}

// Le type EntreeReferentiel est utilisé pour exposer l'ensemble du référentiel des périodes de départ à la retraite sous forme d'API
type EntreeReferentiel struct {
	NaissanceApres			string         		 // date de naissance à partir de laquelle s'applique la législation
	Legislation				PeriodeDepartLegal
}

//
type ReferentielPeriodeLegales struct {
	nom 					string							// nom commun du référentiel
	depuis					time.Time						// date depuis laquelle le référentiel est en vigueur
	url 					string							// URL vers le référentiel ministériel
	periodes 				map[time.Time]PeriodeDepartLegal
}

var ErrPeriodeNonTrouvee = errors.New("pas de période pour la date spécifiée")

// Le référentiel par défaut, en vigueur en Janvier 2016
// voir http://www.la-retraite-en-clair.fr/cid3190611/a-quel-age-peut-partir-la-retraite.html
var ref ReferentielPeriodeLegales

func init() {
	// Chargement du référentiel par défaut
	ref.nom = "Réforme de 2010"
	ref.depuis, _ = time.Parse(JJMMAAADateFormat, "01/01/2010") 	// TODO : A vérifier
	ref.url = "http://www.la-retraite-en-clair.fr/cid3190611/a-quel-age-peut-partir-la-retraite.html"
	ref.periodes = make(map[time.Time]PeriodeDepartLegal)

	// TODO ajouter les périodes avant
	ref.ajouterPeriode("01/01/1955", "", AgeEnAnneesMois{62,0}, AgeEnAnneesMois{67,0},  AgeEnAnneesMois{70,0})
}

// Ajoute une période au référentiel.
// La borne inférieure est incluse, la borne supérieure est excluse.
// Si la période s'étend jusqu'à aujourd'hui, la borne supérieure peut être vide.
// Si la borne inférieure est vide, la période de départ n'est pas bornée.
// Si la borne supérieure est vide, la période de fin n'est pas bornée.
func (r *ReferentielPeriodeLegales) ajouterPeriode(depuis string, jusque string, ageMinimal AgeEnAnneesMois, ageAuTauxPlein AgeEnAnneesMois, ageRetraiteForcee AgeEnAnneesMois) error {
	// L'une ou l'autre des bornes peut être vide mais pas les 2
	if depuis == "" && jusque == "" {
		return ErrDateFormatInvalide
	}
	from := time.Time{}
	var err error
	if depuis != "" {
		if from, err = parseDate(depuis); err != nil {
			return err
		}
	}
	to := time.Time{}
	if jusque != "" {
		if to, err = parseDate(jusque); err != nil {
			return err
		}
	}

	ref.periodes[from] = PeriodeDepartLegal{from, to, ageMinimal, ageAuTauxPlein, ageRetraiteForcee}
	return nil
}

// Retourne les périodes de départ en retraite pour une personne née à la date spécifiée
func GetPeriodeDepartEnRetraite(néLe time.Time) (PeriodeDepartLegal, error) {
	for _, periode := range ref.periodes {
		if dateDansLaPeriode(néLe, periode.Depuis, periode.Jusque) {
			return periode, nil
		}
	}
	return PeriodeDepartLegal{}, ErrPeriodeNonTrouvee
}

// Vérifie si la date spécifiée est dans l'intervalle demandé,
// sachant que la borne min ou max peuvent être égale à zéro (cas où il n'y a de limite inf ou sup)
func dateDansLaPeriode(date time.Time, depuis time.Time, jusque time.Time) bool {
	if depuis.IsZero() {
		return date.Before(jusque)
	}

	if jusque.IsZero() {
		return date.After(depuis)
	}

	return date.Before(jusque) && date.After(depuis)
}