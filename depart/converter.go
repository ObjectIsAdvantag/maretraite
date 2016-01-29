// Copyright 2016, Stève Sfartz
// Licensed under the MIT License

// Fonctions de conversion d'une chaîne de caractères en date au format JJ/MM/AAAA
package depart

import (
	"errors"
	"time"
)

const JJMMAAADateFormat = "02/01/2006"

var ErrDateVide = errors.New("la date n'est pas renseignée")
var ErrDateFormatInvalide = errors.New("le format de date n'est pas valide")

// La fonction StringToTime convertit une chaîne de caractères au format JJ/MM/AAAA en date de type time.Time
func StringToTime(dateJJMMAAAA string) (time.Time, error) {
	if dateJJMMAAAA == "" {
		return time.Time{}, ErrDateVide
	}

	// Parser la date
	res, err := time.Parse(JJMMAAADateFormat, dateJJMMAAAA)
	if err != nil {
		return time.Time{}, ErrDateFormatInvalide
	}

	return res, nil
}

// La fonction TimeToString convertit un type Time en chaîne au format JJ/MM/AAAA
func TimeToString(date time.Time) string {
	return date.Format(JJMMAAADateFormat)
}
