// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Ensemble de fonctions utilitaires pour manipuler des dates et âges
//
// Le type AnneesMoisJours permet d'exprimer des dates et âges en Années, Mois et Jour,
// Le type AnneesMoisJours peut être converti en type time.Time (UTC) et vice versa.
// Enfin, il est possible d'obtenir une durée en AnneesMoisJours à partir de la différence entre 2 dates.
package depart

import (
	"fmt"
	"time"
	"bytes"
	"strconv"
)

// Le type AnneesMois représente une date sous la forme d'un nombre d'années, de mois et de jour
// Exemple : 62 ans et 7 mois
type AnneesMoisJours struct {
	Annees int `json:"annees,required"` // nombre d'années
	Mois   int `json:"mois,omitempty"`  // nombre de mois
	Jours  int `json:"jours,omitempty"` // nombre de jours
}

// Retourne la valeur en années
func (amj AnneesMoisJours) AgeEnAnneesFloat() float32 {
	return float32(amj.Annees) + float32(amj.Mois)/12 + float32(amj.Jours/365)
}

// Retourne la valeur en mois
func (amj AnneesMoisJours) AgeEnMoisFloat() float32 {
	return float32(amj.Annees*12) + float32(amj.Mois) + float32(amj.Jours/365*12)
}

// Retourne la valeur en années
// exemple : 44,3 ans
func (amj AnneesMoisJours) AgeEnAnneesVirguleMois() string {
	return fmt.Sprintf("%.1f", amj.AgeEnAnneesFloat()) + " ans"
}


// Retourne la valeur en années
// exemple : 44 ans
func (amj AnneesMoisJours) AgeEnAnnees() string {
	return string(int(amj.AgeEnAnneesFloat())) + " ans"
}

// Retourne la valeur en années et en mois
// exmple : 44 ans et 2 mois
func (amj AnneesMoisJours) AgeEnAnneesMois() string {
	var buffer bytes.Buffer

	initié := false
	switch amj.Annees {
		case 0:

		case 1 :
			buffer.WriteString("1 an")
			initié = true

		default:
			buffer.WriteString(strconv.Itoa(amj.Annees))
			buffer.WriteString(" ans")
			initié = true
	}

	mois := float32(amj.Mois) + float32(amj.Jours/365*12)
	if mois > 0 {
		if initié {
			buffer.WriteString(" et ")
		}

		buffer.WriteString(strconv.Itoa(int(mois)))
		buffer.WriteString(" mois")
	}

	return buffer.String()
}

// Crée un nouvel objet de type time.Time pour l' AnneesMoisJours spécifié
func AnneesMoisJourToTime(amj AnneesMoisJours) (time.Time, error) {
	if amj.Annees < ANNEE_MIN || amj.Annees > ANNEE_MAX {
		return time.Time{}, ErrDateFormatInvalide
	}
	if amj.Mois < 1 || amj.Mois > 12 {
		return time.Time{}, ErrDateFormatInvalide
	}
	if amj.Jours < 1 || amj.Jours > 32 {
		return time.Time{}, ErrDateFormatInvalide
	}

	return time.Date(amj.Annees, time.Month(amj.Mois), amj.Jours, 0, 0, 0, 0, time.UTC), nil
}

// Crée un nouvel objet de type AnneesMoisJours pour la date spécifiée
func TimeToAnneesMoisJour(t time.Time) (AnneesMoisJours, error) {
	if t.IsZero() {
		return AnneesMoisJours{}, ErrDateFormatInvalide
	}

	return AnneesMoisJours{t.Year(), int(t.Month()), t.Day()}, nil
}

// Implementation alternative basée sur time.AddDate() avec arguments négatifs
// Les premiers tests semblent retourner des résultats différents
// [TODO] A creuser
func CalculerDuréeAlt(depuis time.Time, jusque time.Time) (AnneesMoisJours, error) {
	comp := jusque.AddDate(depuis.Year() * -1, int(depuis.Month()) * -1, depuis.Day() * -1)
	return TimeToAnneesMoisJour(comp)
}

// Cette fonction calcule une durée au format AnneesMoisJour, en tenant compte du calendrier réel pour les dates comparées
func CalculerDurée(depuis time.Time, jusque time.Time) (AnneesMoisJours, error) {

	if jusque.Before(depuis) {
		return AnneesMoisJours{}, fmt.Errorf("la date de fin:%s se situe avant la date de début: %s", jusque, depuis)
	}

	// UPDATE 2016/1/21 La fonction Time.AddDate accepte des arguments négatifs !
	// Etudier s'il est possible de remplacer cet algo maison par :
	// comp, _ := jusque.AddDate(depuis.Year() * -1, int(depuis.Month()) * -1, depuis.Day() * -1)
	// return TimeToAnneesMoisJour(comp), nil

	// Formule de calcul de l'age en année / mois
	// soient AAAA2/MM2/DD2 - AAAA1/MM1/DD1
	//
	// 1. Se rendre à l'année cible et voir si on a dépassé
	// - si ce n'est pas le cas, ok
	// - si c'est le cas, revenir 1 an en arrière,
	// - mémoriser l'année cible, le bond réalisé en année, et le fait qu'on a dû ou non s'arrêter un an avant
	//
	// 2. Se rendre sur le mois cible candidat et calculer la différence de jours
	// - si elle est positive, ne rien faire, on peut réaliser l'opération
	// - si elle est négative, il faut opérer un changement de mois (avec le cas particulier du mois de janvier qui se transforme en décembre)
	// - mémoriser le mois cible et calculer le bond réalisé en mois
	//
	// 3. Se rendre sur l'année et mois cible, et calculer la différence en seconde entre les 2 dates
	// - convertir cette différence en jours

	// 1.
	tempDateCible, _ := AnneesMoisJourToTime(AnneesMoisJours{
		Annees: jusque.Year(),
		Mois:   int(depuis.Month()),
		Jours:  depuis.Day(),
	})

	anneeCible := jusque.Year()
	changementAnnee := 0
	if tempDateCible.After(jusque) {
		anneeCible--
		changementAnnee = 1
	}

	// 2.
	moisCible := int(jusque.Month())
	changementMois2 := 0
	changementAnnee2 := 0
	if depuis.Day() > jusque.Day() {
		moisCible--
		changementMois2 = 1
		if moisCible == 0 {
			moisCible = 12
			changementAnnee2 = 1
		}
	}
	nbMois := int(jusque.Month()) + 12*changementAnnee - int(depuis.Month()) - changementMois2

	// 3.
	tempDateCible, _ = AnneesMoisJourToTime(AnneesMoisJours{
		Annees: jusque.Year() - changementAnnee2,
		Mois:   moisCible,
		Jours:  depuis.Day(),
	})
	deltaJours := jusque.Sub(tempDateCible).Minutes() / 60 / 24

	return AnneesMoisJours{
		Annees: anneeCible - depuis.Year(),
		Mois:   nbMois,
		Jours:  int(deltaJours),
	}, nil
}

// Calcule une nouvelle date en ajoutant un type Time et un type AnneesMoisHomme
func DatePlusAge(date time.Time, age AnneesMoisJours) time.Time {
	return date.AddDate(age.Annees, age.Mois, age.Jours)
}
