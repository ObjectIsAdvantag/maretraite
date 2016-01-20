// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Fonctions de conversion d'une chaîne de caractères en date au format JJ/MM/AAAA
package depart


import (
	"errors"
	"time"
	"fmt"
)

const JJMMAAADateFormat = "2/1/2006"

var ErrDateVide = errors.New("la date n'est pas renseignée")
var ErrDateFormatInvalide = errors.New("le format de date n'est pas valide")
var ErrDateLimites = errors.New("la date n'est pas entre le 01/01/1900 et aujourd'hui")

// La fonction StringToTime convertit une chaîne de caractères au format JJ/MM/AAAA en date de type time.Time
func StringToTime(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, ErrDateVide
	}

	// Parser la date
	res, err := time.Parse(JJMMAAADateFormat, date)
	if err != nil {
		return time.Time{}, ErrDateFormatInvalide
	}

	// Vérifier que la date est bien entre le 1/1/1900 et aujourd'hui
	min, _ := time.ParseInLocation(JJMMAAADateFormat, fmt.Sprintf("01/01/%d", ANNEE_MIN), time.UTC)
	if res.Before(min) || res.After(time.Now()) {
		return time.Time{}, ErrDateLimites
	}

	return res, nil
}

// La fonction TimeToString convertit une date de type time.Time en une chaîne de caractères au format JJ/MM/AAAA
func TimeToString(date time.Time) (string, error) {
	if date.IsZero() {
		return "", ErrDateVide
	}

	return date.Format(JJMMAAADateFormat), nil
}